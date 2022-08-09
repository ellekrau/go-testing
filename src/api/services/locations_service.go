package services

import (
	"go-testing/src/api/domain/locations"
	locationsprovider "go-testing/src/api/providers/locations-provider"
	"go-testing/src/api/utils/errors"
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
