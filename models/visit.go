package models

import (
	"fmt"
	"strings"
	"sync"
	"github.com/pkg/errors"
)

type Visit struct {
	ID         int32 `json:"id,omitempty"`
	LocationID int32 `json:"location"`
	UserID     int32 `json:"user"`
	VisitedAt  int   `json:"visited_at"`
	Mark       uint8 `json:"mark"`
}

type UserVisit struct {
	Mark      int    `json:"mark"`
	VisitedAt int    `json:"visited_at"`
	Place     string `json:"place"`
}

type VisitAvg struct {
	Avg float32 `json:"avg"`
}

type Visits struct {
	Visits []Visit `json:"visits"`
}

type UserVisitsSl struct {
	Visits []UserVisit `json:"visits"`
}

var visitMap map[int32]Visit
var mutexVisit *sync.Mutex

func init() {
	visitMap = make(map[int32]Visit)
	mutexVisit = &sync.Mutex{}
}

func SetVisit(visit Visit) {
	mutexVisit.Lock()
	defer mutexVisit.Unlock()

	visitMap[visit.ID] = visit
}

func GetVisit(id int32) (Visit, error) {
	mutexVisit.Lock()
	defer mutexVisit.Unlock()

	visit, ok := visitMap[id]

	if !ok {
		return visit, NotFound
	}

	return visit, nil
}

func InsertVisits(visits Visits) {

	var user User
	var location Location
	valueStrings := make([]string, 0, len(visits.Visits)+6)
	valueArgs := make([]interface{}, 0, (len(visits.Visits)+6)*5)
	for _, visit := range visits.Visits {
		SetVisit(visit)

		user, _ = GetUser(visit.UserID)
		location, _ = GetLocation(visit.LocationID)

		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, visit.ID)
		valueArgs = append(valueArgs, visit.LocationID)
		valueArgs = append(valueArgs, visit.UserID)
		valueArgs = append(valueArgs, visit.VisitedAt)
		valueArgs = append(valueArgs, visit.Mark)

		valueArgs = append(valueArgs, user.Gender)
		valueArgs = append(valueArgs, user.BirthDate)

		valueArgs = append(valueArgs, location.Country)
		valueArgs = append(valueArgs, location.Distance)
	}
	stmt := fmt.Sprintf(
		"INSERT IGNORE INTO visits (id, location, user, visited_at, mark, gender, birth_date, country, distance) "+
			"VALUES %s",
		strings.Join(valueStrings, ","),
	)

	_, err := db.Exec(stmt, valueArgs...)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func InsertVisit(visit Visit) {
	SetVisit(visit)

	var user User
	var location Location
	valueStrings := make([]string, 0, 5+6)
	valueArgs := make([]interface{}, 0, (5+6)*5)

	user, _ = GetUser(visit.UserID)
	location, _ = GetLocation(visit.LocationID)

	valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?, ?)")

	valueArgs = append(valueArgs, visit.ID)
	valueArgs = append(valueArgs, visit.LocationID)
	valueArgs = append(valueArgs, visit.UserID)
	valueArgs = append(valueArgs, visit.VisitedAt)
	valueArgs = append(valueArgs, visit.Mark)

	valueArgs = append(valueArgs, user.Gender)
	valueArgs = append(valueArgs, user.BirthDate)

	valueArgs = append(valueArgs, location.Country)
	valueArgs = append(valueArgs, location.Distance)

	stmt := fmt.Sprintf(
		"INSERT IGNORE INTO visits (id, location, user, visited_at, mark, gender, birth_date, country, distance) "+
			"VALUES %s",
		strings.Join(valueStrings, ","),
	)
	_, err := db.Exec(stmt, valueArgs...)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetVisitFields() []string {
	return []string{"id", "location", "user", "visited_at", "mark"}
}

func ValidatVsitParams(params map[string]interface{}, scenario string) (result bool) {
	if scenario == "insert" && len(params) != len(GetVisitFields()) {
		return false
	}

	for param, value := range params {
		if value == nil {
			return false
		}

		if scenario == "update" && param == "id" {
			return false
		}

		//if !StringInSlice(param, GetUserFields()) {
		//	return false
		//}
	}

	return true
}

func GetAverage(conditions []Condition) (VisitAvg, error) {
	var query string
	var conditionString string
	var visitAvg VisitAvg
	var average float32

	// where
	if len(conditions) > 0 {
		conditionString += "where "
	}
	for i := 0; i < len(conditions); i++ {
		condition := conditions[i]

		if i > 0 {
			conditionString += condition.JoinCondition + " "
		}

		conditionString += fmt.Sprintf("%s %s %s ", condition.Param, condition.Operator, condition.Value)
	}

	query = fmt.Sprintf("select round(avg(mark), 5) as average from visits %s", conditionString)

	err = db.QueryRow(query).Scan(&average)

	if err != nil {
		return visitAvg, err
	}

	visitAvg = VisitAvg{Avg: average}

	return visitAvg, nil
}

func SelectVisits(conditions []Condition, sort Sort) (UserVisitsSl, error) {
	var query string
	var conditionString string
	var sortString string
	var userVisitsSl UserVisitsSl
	var userVisits []UserVisit = []UserVisit{}

	// where
	if len(conditions) > 0 {
		conditionString += "where "
	}
	for i := 0; i < len(conditions); i++ {
		condition := conditions[i]

		if i > 0 {
			conditionString += condition.JoinCondition + " "
		}

		conditionString += fmt.Sprintf("%s %s %s ", condition.Param, condition.Operator, condition.Value)
	}

	// sort
	if len(sort.Fields) > 0 {
		sortString += " order by "

		for _, sortField := range sort.Fields {
			sortString += " " + sortField
		}

		sortString += " " + sort.Direction + " "
	}

	query = fmt.Sprintf("select mark, visited_at, location from visits %s %s", conditionString, sortString)

	rows, err := db.Query(query)

	if err != nil {
		return userVisitsSl, err
	}

	for rows.Next() {

		var mark int
		var visitedAt int
		var location int

		err = rows.Scan(&mark, &visitedAt, &location)
		if err != nil {
			return userVisitsSl, err
		}

		locationSt, _ := GetLocation(int32(location))

		r := UserVisit{mark, visitedAt, locationSt.Place}
		userVisits = append(userVisits, r)
	}

	userVisitsSl = UserVisitsSl{Visits: userVisits}

	return userVisitsSl, nil
}

func UpdateVisit(visit *Visit, params map[string]interface{}, conditions []Condition) (int64, error) {
	if len(params) < 1 {
		return 0, errors.New("error")
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

	location, ok := params["location"].(int32)
	if ok {
		visit.LocationID = location

		setString += fmt.Sprintf("%s = ?", "location")
		values = append(values, location)
	}

	user, ok := params["user"].(int32)
	if ok {
		visit.UserID = user

		if len(setString) != 0 {
			setString += ","
		}

		setString += fmt.Sprintf("%s = ?", "user")
		values = append(values, user)

	}

	visitedAt, ok := params["visited_at"].(int)
	if ok {
		visit.VisitedAt = visitedAt

		if len(setString) != 0 {
			setString += ","
		}

		setString += fmt.Sprintf("%s = ?", "visited_at")
		values = append(values, visitedAt)
	}

	mark, ok := params["mark"].(uint8)
	if ok {

		if len(setString) != 0 {
			setString += ","
		}

		visit.Mark = mark

		setString += fmt.Sprintf("%s = ?", "mark")
		values = append(values, params["mark"])
	}

	if len(setString) != 0 {
		query = fmt.Sprintf("update visits set %s %s", setString, conditionString)

		stmtIns, err := db.Prepare(query)

		if err != nil {
			return 0, err
		}
		defer stmtIns.Close()

		result, err := stmtIns.Exec(values...)

		return result.RowsAffected()
	}

	return 0, nil
}
