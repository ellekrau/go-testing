package locations

type Country struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	TimeZone       string         `json:"time_zone"`
	GeoInformation GeoInformation `json:"geo_information"`
	States         []State        `json:"states"`
}

type GeoInformation struct {
	Location GeoLocation `json:"location"`
}

type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type State struct {
	ID   string
	Name string
}
