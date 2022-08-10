package test

import (
	"encoding/json"
	"go-testing/src/api/domain/locations"
	locationsprovider "go-testing/src/api/providers/locations-provider"
	"go-testing/src/api/utils/errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetCountryNoError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(`{"id":"BR","name":"Brasil","time_zone":"GMT-03:00","geo_information":{"location":{"latitude":-23.6821604,"longitude":-46.875494}},"states":[{"id":"BR-AC","name":"Acre"},{"id":"BR-AL","name":"Alagoas"},{"id":"BR-AP","name":"Amapá"},{"id":"BR-AM","name":"Amazonas"},{"id":"BR-BA","name":"Bahia"},{"id":"BR-CE","name":"Ceará"},{"id":"BR-DF","name":"Distrito Federal"},{"id":"BR-ES","name":"Espírito Santo"},{"id":"BR-GO","name":"Goiás"},{"id":"BR-MA","name":"Maranhão"},{"id":"BR-MT","name":"Mato Grosso"},{"id":"BR-MS","name":"Mato Grosso do Sul"},{"id":"BR-MG","name":"Minas Gerais"},{"id":"BR-PR","name":"Paraná"},{"id":"BR-PB","name":"Paraíba"},{"id":"BR-PA","name":"Pará"},{"id":"BR-PE","name":"Pernambuco"},{"id":"BR-PI","name":"Piauí"},{"id":"BR-RN","name":"Rio Grande do Norte"},{"id":"BR-RS","name":"Rio Grande do Sul"},{"id":"BR-RJ","name":"Rio de Janeiro"},{"id":"BR-RO","name":"Rondônia"},{"id":"BR-RR","name":"Roraima"},{"id":"BR-SC","name":"Santa Catarina"},{"id":"BR-SE","name":"Sergipe"},{"id":"BR-SP","name":"São Paulo"},{"id":"BR-TO","name":"Tocantins"}]}`))
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

	if response.StatusCode != http.StatusOK {
		t.Errorf("should return '200 - OK' response status code, but returned '%d'", response.StatusCode)
	}

	var country locations.Country
	decoder := json.NewDecoder(response.Body)
	decoder.Decode(&country)

	if country.ID != "BR" {
		t.Errorf("should return 'BR' country ID, but returned '%s'", country.ID)
	}

	if country.Name != "Brasil" {
		t.Errorf("should return 'Brasil' country name, but returned '%s'", country.Name)
	}
}

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
