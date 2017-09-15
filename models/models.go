package models

import (
	"errors"
	"os"
	"sort"
	"strconv"
	"sync"
)

var (
	NotFound         error = errors.New("NotFound")
	timeNow          int
	userVisitMap     = make(map[int][]int)
	locationVisitMap = make(map[int][]int)
	mutexUserVisit   = &sync.RWMutex{}
)

type FloatPrecision5 float32

type UserVisit struct {
	Mark      int    `json:"mark"`
	VisitedAt int    `json:"visited_at"`
	Place     string `json:"place"`
}

type UserVisitsSl struct {
	Visits []UserVisit `json:"visits"`
}

func init() {
	file, err := os.Open("/tmp/data/options.txt")
	if err != nil {
		file, err = os.Open("/home/andrey/go/src/hlcupdoc/data/FULL/data/options.txt")
		if err != nil {
			PanicOnErr(err)
		}
	}
	defer file.Close()

	timestampBytes := make([]byte, 10)
	_, err = file.Read(timestampBytes)
	PanicOnErr(err)

	timeNow, err = strconv.Atoi(string(timestampBytes))
	PanicOnErr(err)
}

type UserVisitSt []UserVisit

func (slice UserVisitSt) Len() int {
	return len(slice)
}

func (slice UserVisitSt) Less(i, j int) bool {
	return slice[i].VisitedAt < slice[j].VisitedAt
}

func (slice UserVisitSt) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func SetVisits(visit Visit) {
	//user, err1 := GetUser(visit.User)
	//location, err2 := GetLocation(visit.Location)

	mutexUserVisit.Lock()
	//if err1 == nil && err2 == nil {
	if uv, ok := userVisitMap[visit.User]; ok {
		userVisitMap[visit.User] = append(uv, visit.Id)
	} else {
		userVisitMap[visit.User] = []int{visit.Id}
	}

	if lv, ok := locationVisitMap[visit.Location]; ok {
		locationVisitMap[visit.Location] = append(lv, visit.Id)
	} else {
		locationVisitMap[visit.Location] = []int{visit.Id}
	}

	//fmt.Println(userVisitMap[visit.User])
	//fmt.Println(locationVisitMap[visit.Location])

	//}
	mutexUserVisit.Unlock()
}

func updateUsetLocationVisit(visit Visit) {
	mutexUserVisit.Lock()

	oldUserId := visitMap[visit.Id].User
	if oldUserId != visit.User {
		for i, uv := range userVisitMap[oldUserId] {
			if uv == visit.Id {
				userVisitMap[oldUserId] = remove(userVisitMap[oldUserId], i)
				userVisitMap[visit.User] = append(userVisitMap[visit.User], visit.Id)
				break
			}
		}
	}
	oldLocationId := visitMap[visit.Id].Location
	if oldLocationId != visit.Location {
		for i, lv := range locationVisitMap[oldLocationId] {
			if lv == visit.Id {
				locationVisitMap[oldLocationId] = remove(locationVisitMap[oldLocationId], i)
				locationVisitMap[visit.Location] = append(locationVisitMap[visit.Location], visit.Id)
				break
			}
		}
	}
	mutexUserVisit.Unlock()
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func GetAverage(locationId, fromDate, toDate, fromAge, toAge int, gender string) (float64, error) {
	var avg float64 = 0

	var marksSum = 0
	var markCount = 0

	for _, visitId := range locationVisitMap[locationId] {

		visit, ok := visitMap[visitId]
		if !ok {
			continue
		}

		if fromDate != 0 && visit.VisitedAt <= fromDate {
			continue
		}

		if toDate != 0 && visit.VisitedAt >= toDate {
			continue
		}

		user, ok := userMap[visit.User]
		if !ok {
			continue
		}

		if len(gender) != 0 && user.Gender != gender {
			continue
		}

		if fromAge != 0 && timeNow-user.BirthDate < fromAge*31557600 {
			continue
		}

		if toAge != 0 && timeNow-user.BirthDate > toAge*31557600 {
			continue
		}

		marksSum += visit.Mark
		markCount++
	}

	if markCount > 0 {
		avg = float64(marksSum)/float64(markCount) + 0.00000001
	}

	return avg, nil
}

func SelectVisits(userId, fromDate, toDate, toDistance int, country string) (UserVisitsSl, error) {
	var userVisitsSl UserVisitsSl
	var userVisits = UserVisitSt{}

	for _, visitId := range userVisitMap[userId] {

		visit, ok := visitMap[visitId]
		if !ok {
			continue
		}

		if fromDate != 0 && visit.VisitedAt <= fromDate {
			continue
		}

		if toDate != 0 && visit.VisitedAt >= toDate {
			continue
		}

		location, ok := locationMap[visit.Location]
		if !ok {
			continue
		}

		if len(country) != 0 && location.Country != country {
			continue
		}

		if toDistance != 0 && location.Distance >= toDistance {
			continue
		}

		userVisits = append(userVisits, UserVisit{visit.Mark, visit.VisitedAt, location.Place})
	}

	sort.Sort(userVisits)

	userVisitsSl = UserVisitsSl{Visits: userVisits}

	return userVisitsSl, nil
}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
