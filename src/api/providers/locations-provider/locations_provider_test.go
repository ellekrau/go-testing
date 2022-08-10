package locationsprovider

import (
	"encoding/json"
	"go-testing/src/api/utils/errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Timeout
// Error from API
// Invalid error interface
// Valid json response
// Invalid json response

func TestGetCountryApiValidJsonResponse(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(`{"id":"BR","name":"Brasil","time_zone":"GMT-03:00","geo_information":{"location":{"latitude":-23.6821604,"longitude":-46.875494}},"states":[{"id":"BR-AC","name":"Acre"},{"id":"BR-AL","name":"Alagoas"},{"id":"BR-AP","name":"Amapá"},{"id":"BR-AM","name":"Amazonas"},{"id":"BR-BA","name":"Bahia"},{"id":"BR-CE","name":"Ceará"},{"id":"BR-DF","name":"Distrito Federal"},{"id":"BR-ES","name":"Espírito Santo"},{"id":"BR-GO","name":"Goiás"},{"id":"BR-MA","name":"Maranhão"},{"id":"BR-MT","name":"Mato Grosso"},{"id":"BR-MS","name":"Mato Grosso do Sul"},{"id":"BR-MG","name":"Minas Gerais"},{"id":"BR-PR","name":"Paraná"},{"id":"BR-PB","name":"Paraíba"},{"id":"BR-PA","name":"Pará"},{"id":"BR-PE","name":"Pernambuco"},{"id":"BR-PI","name":"Piauí"},{"id":"BR-RN","name":"Rio Grande do Norte"},{"id":"BR-RS","name":"Rio Grande do Sul"},{"id":"BR-RJ","name":"Rio de Janeiro"},{"id":"BR-RO","name":"Rondônia"},{"id":"BR-RR","name":"Roraima"},{"id":"BR-SC","name":"Santa Catarina"},{"id":"BR-SE","name":"Sergipe"},{"id":"BR-SP","name":"São Paulo"},{"id":"BR-TO","name":"Tocantins"}]}`))
	}))
	defer mockServer.Close()
	GetCountryUri = mockServer.URL

	country, err := LocationsProvider.GetCountry("BR")

	if country == nil {
		t.Fatal("country shouldn't be nil")
	}

	if err != nil {
		t.Fatal("error should be nil")
	}

	if country.ID != "BR" {
		t.Errorf("country ID should be 'BR'")
	}

	if country.Name != "Brasil" {
		t.Errorf("country name should be 'Brasil'")
	}

	if country.TimeZone != "GMT-03:00" {
		t.Errorf("country timezone should be 'GMT-03:00'")
	}

	if len(country.States) != 27 {
		t.Errorf("country states should have len 27, bur received len %d", len(country.States))
	}
}

func TestGetCountryApiCountryNotFoundError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNotFound)
		response, _ := json.Marshal(errors.ApiError{
			Status:  http.StatusNotFound,
			Message: "Country not found",
		})
		writer.Write(response)
	}))
	defer mockServer.Close()
	GetCountryUri = mockServer.URL

	country, err := LocationsProvider.GetCountry("BRA")

	if country != nil {
		t.Errorf("country should be nil")
	}

	if err == nil {
		t.Fatal("error shouldn't be nil")
	}

	if err.Status != http.StatusNotFound {
		t.Errorf("should return 'not found' error status, but returned '%d'", err.Status)
	}

	if err.Message != "Country not found" {
		t.Errorf("should return 'Country not found 'BR'' error message, but returned '%s'", err.Message)
	}
}

func TestGetCountryApiInvalidJsonResponse(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(`{"id":123"}`))
	}))
	defer mockServer.Close()
	GetCountryUri = mockServer.URL

	country, err := LocationsProvider.GetCountry("")

	if country != nil {
		t.Errorf("country should be nil")
	}

	if err == nil {
		t.Fatal("error shouldn't be nil")
	}

	if err.Status != http.StatusInternalServerError {
		t.Errorf("should return 'internal server error' error status, but returned %d", err.Status)
	}

	if err.Message != "unmarshal country data error when trying get '' country" {
		t.Errorf("should return 'unmarshal country data error when trying get 'BR' country, but returned %s", err.Message)
	}
}

func TestGetCountryRestClientError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Duration(timeoutLimit) * time.Millisecond)
		writer.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()
	GetCountryUri = mockServer.URL

	country, err := LocationsProvider.GetCountry("BR")

	if country != nil {
		t.Errorf("country should be nil")
	}

	if err == nil {
		t.Fatal("error shouldn't be nil")
	}

	if err.Status != http.StatusInternalServerError {
		t.Errorf("should return 'internal server error' error status code")
	}

	if err.Message != "invalid rest client error when getting country 'BR'" {
		t.Errorf("should return 'invalid rest client error when getting country 'BR'' error message, but returned '%s'", err.Message)
	}
}

func TestGetCountryApiInvalidErrorInterface(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(`{"status": "404","message": "Country not found"}`))
	}))
	defer mockServer.Close()
	GetCountryUri = mockServer.URL

	country, err := LocationsProvider.GetCountry("BR")

	if country != nil {
		t.Fatal("country should be nil")
	}

	if err == nil {
		t.Fatal("error shouldn't be nil")
	}

	if err.Status != http.StatusInternalServerError {
		t.Errorf("should return 'internal server error' error status")
	}

	if err.Message != "invalid error response interface when trying get 'BR' country" {
		t.Errorf("should return 'invalid error response interface when trying get 'BR' country' error message, but returned '%s'", err.Message)
	}
}
