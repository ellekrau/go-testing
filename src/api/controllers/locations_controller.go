package controllers

import (
	"encoding/json"
	locationsprovider "go-testing/src/api/providers/locations-provider"
	"net/http"
	"strings"
)

func GetCountry(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	countryID := strings.Replace(request.URL.Path, "/country/", "", 1)
	country, apiError := locationsprovider.GetCountry(countryID)
	if apiError != nil {
		writer.WriteHeader(apiError.Status)
		responseBody, _ := json.Marshal(*apiError)
		writer.Write(responseBody)
		return
	}

	writer.WriteHeader(http.StatusOK)
	responseBody, _ := json.Marshal(*country)
	writer.Write(responseBody)
	return
}
