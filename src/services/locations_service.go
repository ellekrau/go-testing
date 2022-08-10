package services

import (
	"go-testing/domain/locations"
	locationsprovider "go-testing/providers/locations-provider"
	"go-testing/utils/errors"
	"strings"
)

var (
	LocationService locationServiceInterface
)

type locationServiceInterface interface {
	GetCountry(countryID string) (*locations.Country, *errors.ApiError)
}

type locationService struct{}

func init() {
	LocationService = &locationService{}
}

func (ls *locationService) GetCountry(countryID string) (*locations.Country, *errors.ApiError) {
	countryID = strings.ToUpper(countryID)
	return locationsprovider.LocationsProvider.GetCountry(countryID)
}
