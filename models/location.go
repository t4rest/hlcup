package models

import (
	"sync"
)

type Location struct {
	Id       int    `json:"id"`
	Place    string `json:"place"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Distance int    `json:"distance"`
}

type Locations struct {
	Locations []Location `json:"locations"`
}

var locationMap = make(map[int]Location)
var mutexLocation = &sync.RWMutex{}

func SetLocation(location Location) {
	mutexLocation.Lock()
	locationMap[location.Id] = location
	mutexLocation.Unlock()
}

func GetLocation(id int) (Location, error) {
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
		SetLocation(location)
	}
}

func UpdateLocation(location Location, locationNew Location) int {

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

	SetLocation(location)

	return 1
}
