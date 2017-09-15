package models

import (
	"sync"
)

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate int    `json:"birth_date"`
}

type Users struct {
	Users []User `json:"users"`
}

var userMap = make(map[int]User)
var mutexUser = &sync.RWMutex{}

func SetUser(user User) {
	mutexUser.Lock()
	userMap[user.Id] = user
	mutexUser.Unlock()
}

func GetUser(id int) (User, error) {
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
		SetUser(user)
	}
}

func UpdateUser(user User, userNew User, birthDateUpdate bool) int {

	if userNew.BirthDate != 0 || birthDateUpdate {
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

	SetUser(user)

	return 1
}
