package controllers

import (
	"encoding/json"
	"go-testing/services"
	"net/http"
	"strings"
)

func GetCountry(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	countryID := strings.Replace(request.URL.Path, "/locations/country/", "", 1)
	country, apiError := services.LocationService.GetCountry(countryID)
	if apiError != nil {
		writer.WriteHeader(apiError.Status)
		responseBody, _ := json.Marshal(*apiError)
		_, _ = writer.Write(responseBody)
		return
	}

	writer.WriteHeader(http.StatusOK)
	responseBody, _ := json.Marshal(*country)
	_, _ = writer.Write(responseBody)
	return
}
