package test

import (
	"encoding/json"
	locationsprovider "go-testing/src/api/providers/locations-provider"
	"go-testing/src/api/utils/errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetCountryNotFound(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(`{"status": 404,"message": "Country not found"}`))
	}))
	defer mockServer.Close()
	locationsprovider.GetCountryUri = mockServer.URL

	response, err := http.Get(baseURL + "/locations/country/BR")
	if err != nil {
		t.Error("error should be nil")
	}

	if response == nil {
		t.Error("response shouldn't be nil")
		return
	}

	if response.StatusCode != http.StatusNotFound {
		t.Errorf("should return '404 - not found' response status code, but returned '%d'", response.StatusCode)
	}

	var responseApiError errors.ApiError
	decoder := json.NewDecoder(response.Body)
	decoder.Decode(&responseApiError)

	if responseApiError.Status != http.StatusNotFound {
		t.Errorf("should return '404 - not found' api error response status code, but returned '%d'", responseApiError.Status)
	}

	if responseApiError.Message != "Country not found" {
		t.Errorf("should return 'Country not found' api error message, but returned '%s'", responseApiError.Message)
	}
}

func TestGetCountryInternalServerError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Second * 2)
	}))
	defer mockServer.Close()
	locationsprovider.GetCountryUri = mockServer.URL

	response, err := http.Get(baseURL + "/locations/country/BR")
	if err != nil {
		t.Error("error should be nil")
		return
	}

	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("should return '500 - internal server error' status code, but returned '%d'", response.StatusCode)
	}

	var responseApiError errors.ApiError
	decoder := json.NewDecoder(response.Body)
	decoder.Decode(&responseApiError)

	if responseApiError.Status != http.StatusInternalServerError {
		t.Errorf("should return '500 - Internal Server Error' api error response status code, but returned '%d", responseApiError.Status)
	}

	if responseApiError.Message != "invalid rest client error when getting country 'BR'" {
		t.Errorf("should return 'invalid rest client error when getting country 'BR'' api error message, but returned '%s'", responseApiError.Message)
	}
}
