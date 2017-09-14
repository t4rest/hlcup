package models

import (
	"sync"
)

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate int    `json:"birth_date"`
	//Visits    []Visit `json:"-"`
}

type Users struct {
	Users []*User `json:"users"`
}

var userMap map[int]*User
var mutexUser *sync.RWMutex

func init() {
	userMap = make(map[int]*User)
	mutexUser = &sync.RWMutex{}
}

func SetUser(user *User) {
	mutexUser.Lock()
	userMap[user.ID] = user
	mutexUser.Unlock()

}

func GetUser(id int) (*User, error) {
	mutexUser.RLock()
	user, ok := userMap[id]
	mutexUser.RUnlock()

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

func InsertUser(user *User) {
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
	}

	return true
}

func UpdateUser(user *User, userNew *User) int {

	if userNew.BirthDate != 0 {
		user.BirthDate = userNew.BirthDate
	}
	if userNew.Gender != "" {
		user.Gender = userNew.Gender
	}
	if userNew.FirstName != "" {
		user.FirstName = userNew.FirstName
	}
	if userNew.LastName != "" {
		user.LastName = userNew.LastName
	}
	if userNew.Email != "" {
		user.Email = userNew.Email
	}

	return 1
}
