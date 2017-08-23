package models

import (
	"fmt"
	"sync"
)

type User struct {
	ID        int32  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate int    `json:"birth_date"`
	//Visits    []Visit `json:"-"`
}

type Users struct {
	Users []User `json:"users"`
}

var userMap map[int32]User
var mutexUser *sync.RWMutex

func init() {
	userMap = make(map[int32]User)
	mutexUser = &sync.RWMutex{}
}

func SetUser(user User) {
	mutexUser.RLock()
	defer mutexUser.RUnlock()

	userMap[user.ID] = user
}

func GetUser(id int32) (User, error) {
	//mutexUser.RLock()
	//defer mutexUser.RUnlock()

	user, ok := userMap[id]

	if !ok {
		return user, NotFound
	}

	return user, nil
}

func InsertUsers(users Users) {
	for _, user := range users.Users {
		InsertUser(user)
	}
}

func InsertUser(user User) {
	SetUser(user)
}

func GetUserFields() []string {
	return []string{"id", "email", "first_name", "last_name", "gender", "birth_date"}
}

func ValidateUserParams(params map[string]interface{}, scenario string) (result bool) {
	if scenario == "insert" && len(params) != len(GetUserFields()) {
		return false
	}

	for param, value := range params {
		if value == nil {
			return false
		}

		if scenario == "update" && param == "id" {
			return false
		}

		if !StringInSlice(param, GetUserFields()) {
			return false
		}
	}

	return true
}

func UpdateUser(user User, params map[string]interface{}, conditions []Condition) (rowsAffected int64, err error) {
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

	setString += fmt.Sprintf("%s = ?", "gender")
	values = append(values, params["gender"])

	setString += ","

	setString += fmt.Sprintf("%s = ?", "birth_date")
	values = append(values, params["birth_date"])

	email, ok := params["email"].(string)
	if ok {
		user.Email = email
	}

	firstName, ok := params["first_name"].(string)
	if ok {
		user.FirstName = firstName
	}

	lastName, ok := params["last_name"].(string)
	if ok {
		user.LastName = lastName
	}

	gender, ok := params["gender"].(string)
	if ok {
		user.Gender = gender
	}

	birthDate, ok := params["birth_date"].(int)
	if ok {
		user.BirthDate = birthDate
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
