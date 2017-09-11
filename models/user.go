package models

import (
	"sync"
	"github.com/pkg/errors"
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
	Users []*User `json:"users"`
}

var userMap map[int32]*User
var mutexUser *sync.Mutex

func init() {
	userMap = make(map[int32]*User)
	mutexUser = &sync.Mutex{}
}

func SetUser(user *User) {
	mutexUser.Lock()
	defer mutexUser.Unlock()

	userMap[user.ID] = user
}

func GetUser(id int32) (*User, error) {
	mutexUser.Lock()
	defer mutexUser.Unlock()

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

func UpdateUser(user *User, params map[string]interface{}, userNew *User) (int64, error) {
	if len(params) < 1 {
		return 0, errors.New("error")
	}

	userNew.ID = user.ID
	if userNew.BirthDate == 0 { userNew.BirthDate = user.BirthDate }
	if userNew.Gender == "" { userNew.Gender = user.Gender }
	if userNew.FirstName == "" { userNew.FirstName = user.FirstName }
	if userNew.LastName == "" { userNew.LastName = user.LastName }
	if userNew.Email == "" { userNew.Email = user.Email }

	SetUser(userNew)

	return 1, nil
}
