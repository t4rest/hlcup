package models

import (
	"sync"
)

type Location struct {
	ID       int    `json:"id"`
	Place    string `json:"place"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Distance int    `json:"distance"`
}

type Locations struct {
	Locations []*Location `json:"locations"`
}

var locationMap map[int]*Location
var mutexLocation *sync.RWMutex

func init() {
	locationMap = make(map[int]*Location)
	mutexLocation = &sync.RWMutex{}
}

func SetLocation(location *Location) {
	mutexLocation.Lock()
	locationMap[location.ID] = location
	mutexLocation.Unlock()
}

func GetLocation(id int) (*Location, error) {
	mutexLocation.RLock()
	location, ok := locationMap[id]
	mutexLocation.RUnlock()

	if !ok {
		return location, NotFound
	}

	return location, nil
}

func InsertLocations(locations Locations) {
	for _, location := range locations.Locations {
		InsertLocation(location)
	}
}

func InsertLocation(location *Location) {
	SetLocation(location)
}

func GetLocationFields() []string {
	return []string{"id", "place", "country", "city", "distance"}
}

func ValidateLocationParams(params map[string]interface{}, scenario string) (result bool) {
	if scenario == "insert" && len(params) != len(GetLocationFields()) {
		return false
	}

	for param, value := range params {
		if value == nil {
			return false
		}

		if scenario == "update" && param == "id" {
			return false
		}
	}

	return true
}

func UpdateLocation(location *Location, locationNew *Location) int {

	if locationNew.Place != "" {
		location.Place = locationNew.Place
	}
	if locationNew.Country != "" {
		location.Country = locationNew.Country
	}
	if locationNew.City != "" {
		location.City = locationNew.City
	}
	if locationNew.Distance != 0 {
		location.Distance = locationNew.Distance
	}

	return 1
}
