package models

import (
	"fmt"
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
	Locations []Location `json:"locations"`
}

var locationMap map[int32]Location
var mutexLocation *sync.RWMutex

func init() {
	locationMap = make(map[int32]Location)
	mutexLocation = &sync.RWMutex{}
}

func SetLocation(location Location) {
	mutexLocation.RLock()
	defer mutexLocation.RUnlock()

	locationMap[location.ID] = location
}

func GetLocation(id int32) (Location, error) {
	//mutexLocation.RLock()
	//defer mutexLocation.RUnlock()

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

func InsertLocation(location Location) {
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

		if !StringInSlice(param, GetLocationFields()) {
			return false
		}
	}

	return true
}

func UpdateLocation(visit Location, params map[string]interface{}, conditions []Condition) (rowsAffected int64, err error) {
	if len(params) < 1 {
		return
	}

	var query string
	var conditionString string
	var setString string
	var values []interface{}

	if len(conditions) > 0 {
		conditionString += "where "
	}
	for i := 0; i < len(conditions); i++ {
		condition := conditions[i]

		if i > 0 {
			conditionString += condition.JoinCondition + " "
		}

		conditionString += fmt.Sprintf("%s %s %s", condition.Param, condition.Operator, condition.Value)
	}

	setString += fmt.Sprintf("%s = ?", "country")
	values = append(values, params["country"])

	setString += ","

	setString += fmt.Sprintf("%s = ?", "distance")
	values = append(values, params["distance"])

	place, ok := params["place"].(string)
	if ok {
		visit.Place = place
	}

	country, ok := params["country"].(string)
	if ok {
		visit.Country = country
	}

	city, ok := params["city"].(string)
	if ok {
		visit.City = city
	}

	distance, ok := params["distance"].(int32)
	if ok {
		visit.Distance = distance
	}

	query = fmt.Sprintf("update visits set %s %s", setString, conditionString)

	stmtIns, err := db.Prepare(query)

	if err != nil {
		return
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(values...)

	return result.RowsAffected()
}
