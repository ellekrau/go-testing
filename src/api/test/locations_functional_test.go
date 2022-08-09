package test

import (
	"net/http"
	"testing"
)

func TestGetCountryNotFound(t *testing.T) {
	response, err := http.Get(baseURL + "/locations/country/BR")

	if err != nil {
		t.Error("error should be nil")
	}

	if response == nil {
		t.Error("response shouldn't be nil")
	}
}
