package controllers

import (
	"encoding/json"
	"go-testing/src/api/services"
	"net/http"
	"strings"
)

func GetCountry(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	countryID := strings.Replace(request.URL.Path, "/locations/country/", "", 1)
	country, apiError := services.GetCountry(countryID)
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
