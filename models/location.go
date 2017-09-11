package models

import (
	"github.com/pkg/errors"
	"sync"
)

type Location struct {
	ID       int32  `json:"id"`
	Place    string `json:"place"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Distance int32  `json:"distance"`
}

type Locations struct {
	Locations []*Location `json:"locations"`
}

var locationMap map[int32]*Location
var mutexLocation *sync.Mutex

func init() {
	locationMap = make(map[int32]*Location)
	mutexLocation = &sync.Mutex{}
}

func SetLocation(location *Location) {
	mutexLocation.Lock()
	defer mutexLocation.Unlock()

	locationMap[location.ID] = location
}

func GetLocation(id int32) (*Location, error) {
	mutexLocation.Lock()
	defer mutexLocation.Unlock()

	location, ok := locationMap[id]

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

func UpdateLocation(location *Location, params map[string]interface{}, locationNew *Location) (int64, error) {
	if len(params) < 1 {
		return 0, errors.New("error")
	}

	locationNew.ID = location.ID
	if locationNew.Place == "" {
		locationNew.Place = location.Place
	}
	if locationNew.Country == "" {
		locationNew.Country = location.Country
	}
	if locationNew.City == "" {
		locationNew.City = location.City
	}
	if locationNew.Distance == 0 {
		locationNew.Distance = location.Distance
	}

	SetLocation(locationNew)

	return 1, nil
}
