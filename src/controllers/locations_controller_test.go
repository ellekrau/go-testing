package controllers

import (
	"encoding/json"
	"go-testing/domain/locations"
	"go-testing/services"
	"go-testing/utils/errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	getCountryFunction func(countryID string) (*locations.Country, *errors.ApiError)
)

type locationsServiceMock struct{}

func (lsm *locationsServiceMock) GetCountry(countryID string) (*locations.Country, *errors.ApiError) {
	return getCountryFunction(countryID)
}

func TestGetCountryNoError(t *testing.T) {
	getCountryFunction = func(countryID string) (*locations.Country, *errors.ApiError) {
		return &locations.Country{
			ID:   "BR",
			Name: "Brasil",
		}, nil
	}
	services.LocationService = &locationsServiceMock{}

	request := httptest.NewRequest(http.MethodGet, "/locations/country/BR", nil)
	responseRecorder := httptest.NewRecorder()

	GetCountry(responseRecorder, request)

	var countryResponse locations.Country
	json.Unmarshal(responseRecorder.Body.Bytes(), &countryResponse)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("shoud return status code 'OK' in, but returned '%d'", responseRecorder.Code)
	}

	if countryResponse.ID != "BR" {
		t.Errorf("shoud return 'BR' in country ID, but returned '%s'", countryResponse.ID)
	}

	if countryResponse.Name != "Brasil" {
		t.Errorf("shoud return 'Brasil' in country name, but returned '%s'", countryResponse.Name)
	}
}

func TestGetCountryNotFound(t *testing.T) {
	getCountryFunction = func(countryID string) (*locations.Country, *errors.ApiError) {
		return nil, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: "Country not found",
		}
	}
	services.LocationService = &locationsServiceMock{}

	request := httptest.NewRequest(http.MethodGet, "/locations/country/BR", nil)
	responseRecorder := httptest.NewRecorder()

	GetCountry(responseRecorder, request)

	var apiErrResponse errors.ApiError
	json.Unmarshal(responseRecorder.Body.Bytes(), &apiErrResponse)

	if apiErrResponse.Status != http.StatusNotFound {
		t.Errorf("should return 'not found' error status, but returned '%d'", apiErrResponse.Status)
	}

	if apiErrResponse.Message != "Country not found" {
		t.Errorf("should return 'Country not found' error message, but returned '%s'", apiErrResponse.Message)
	}
}
