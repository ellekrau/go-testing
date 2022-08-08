package locationsprovider

import (
	"encoding/json"
	"fmt"
	"go-testing/src/api/domain/locations"
	"go-testing/src/api/utils/errors"
	"net/http"
	"time"
)

var (
	timeoutLimit  = 2000
	getCountryUri = "https://api.mercadolibre.com/countries"
)

func GetCountry(countryID string) (*locations.Country, *errors.ApiError) {
	httpClient := http.Client{
		Timeout: time.Millisecond * time.Duration(timeoutLimit),
	}

	response, err := httpClient.Get(fmt.Sprintf("%s/%s", getCountryUri, countryID))
	if err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid rest client error when getting country '%s'", countryID),
		}
	}
	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)

	if response.StatusCode > 299 {
		apiError := &errors.ApiError{}
		if err = decoder.Decode(apiError); err != nil {
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error response interface when trying get '%s' country", countryID),
			}
		}
		return nil, apiError
	}

	if response.StatusCode > 299 {
		return nil, nonSuccessStatusCode(decoder, countryID)
	}

	return successStatusCode(decoder, countryID)
}

func nonSuccessStatusCode(decoder *json.Decoder, countryID string) *errors.ApiError {
	apiError := &errors.ApiError{}
	if err := decoder.Decode(apiError); err != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid error response interface when trying get '%s' country", countryID),
		}
	}
	return apiError
}

func successStatusCode(decoder *json.Decoder, countryID string) (*locations.Country, *errors.ApiError) {
	country := &locations.Country{}
	if err := decoder.Decode(country); err != nil || country.ID == "" {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("unmarshal country data error when trying get '%s' country", countryID),
		}
	}
	return country, nil
}
