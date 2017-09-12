package models

import (
	"sort"
	"sync"
	"time"
)

type Visit struct {
	ID         int32 `json:"id,omitempty"`
	LocationID int32 `json:"location"`
	UserID     int32 `json:"user"`
	VisitedAt  int   `json:"visited_at"`
	Mark       uint8 `json:"mark"`
}

type UserVisit struct {
	Mark      uint8  `json:"mark"`
	VisitedAt int    `json:"visited_at"`
	Place     string `json:"place"`
}

type VisitAvg struct {
	Avg float32 `json:"avg"`
}

type Visits struct {
	Visits []*Visit `json:"visits"`
}

type UserVisitsSl struct {
	Visits []UserVisit `json:"visits"`
}

type UserVisits struct {
	Visit    *Visit
	Location *Location
	User     *User
}

type LocationVisits struct {
	Visit    *Visit
	Location *Location
	User     *User
}

var visitMap map[int32]*Visit
var userVisitMap map[int32][]UserVisits
var locationVisitMap map[int32][]LocationVisits
var mutexVisit *sync.RWMutex
var mutexUserVisit *sync.RWMutex

var timeNow int64 = time.Now().UnixNano() / int64(time.Millisecond)

const oneYear = 31557600

func init() {
	visitMap = make(map[int32]*Visit)
	userVisitMap = make(map[int32][]UserVisits)
	locationVisitMap = make(map[int32][]LocationVisits)
	mutexVisit = &sync.RWMutex{}
	mutexUserVisit = &sync.RWMutex{}
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

func SetVisit(visit *Visit) {
	mutexVisit.Lock()
	visitMap[visit.ID] = visit
	mutexVisit.Unlock()
}

func GetVisit(id int32) (*Visit, error) {
	mutexVisit.RLock()
	visit, ok := visitMap[id]
	mutexVisit.RUnlock()

	if !ok {
		return visit, NotFound
	}

	return visit, nil
}

func InsertVisits(visits Visits) {
	for _, visit := range visits.Visits {
		InsertVisit(visit)
	}
}

func InsertVisit(visit *Visit) {
	SetVisit(visit)

	mutexUserVisit.Lock()

	user, err1 := GetUser(visit.UserID)
	location, err2 := GetLocation(visit.LocationID)

	if err1 == nil && err2 == nil {
		userVisitMap[visit.UserID] = append(userVisitMap[visit.UserID], UserVisits{visit, location, user})
		locationVisitMap[visit.LocationID] = append(locationVisitMap[visit.LocationID], LocationVisits{visit, location, user})
	}

	mutexUserVisit.Unlock()
}

func GetVisitFields() []string {
	return []string{"id", "location", "user", "visited_at", "mark"}
}

func ValidateVsitParams(params map[string]interface{}, scenario string) (result bool) {
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
	}

	return true
}

func GetAverage(id, fromDate, toDate, fromAge, toAge int, gender string) (float64, error) {
	var avg float64 = 0

	var marksSum uint8 = 0
	var markCount int64 = 0

	for _, sl := range locationVisitMap[int32(id)] {

		if fromDate > 0 && fromDate > sl.Visit.VisitedAt {
			continue
		}

		if toDate > 0 && toDate < sl.Visit.VisitedAt {
			continue
		}

		if fromAge > 0 {
			if !((float64(time.Now().Unix())-float64(sl.User.BirthDate))/31557600 > float64(fromAge)) {
				continue
			}
		}

		if toAge > 0 {
			if !((float64(time.Now().Unix())-float64(sl.User.BirthDate))/31557600 < float64(toAge)) {
				continue
			}
		}

		if len(gender) > 0 && gender != sl.User.Gender {
			continue
		}

		marksSum += sl.Visit.Mark
		markCount++
	}

	if markCount > 0 {
		avg = float64(marksSum)/float64(markCount) + 0.00000001
	}

	return avg, nil
}

func SelectVisits(id, fromDate, toDate, toDistance int, country string) (UserVisitsSl, error) {
	var userVisitsSl UserVisitsSl
	var userVisits = UserVisitSt{}

	for _, sl := range userVisitMap[int32(id)] {

		if fromDate > 0 && fromDate > sl.Visit.VisitedAt {
			continue
		}

		if toDate > 0 && toDate < sl.Visit.VisitedAt {
			continue
		}

		if len(country) > 0 && country != sl.Location.Country {
			continue
		}

		if toDistance > 0 && int32(toDistance) <= sl.Location.Distance {
			continue
		}

		userVisits = append(userVisits, UserVisit{sl.Visit.Mark, sl.Visit.VisitedAt, sl.Location.Place})
	}

	sort.Sort(userVisits)

	userVisitsSl = UserVisitsSl{Visits: userVisits}

	return userVisitsSl, nil
}

func UpdateVisit(visit *Visit, visitNew *Visit) int64 {

	visitNew.ID = visit.ID
	if visitNew.LocationID == 0 {
		visitNew.LocationID = visit.LocationID
	}
	if visitNew.UserID == 0 {
		visitNew.UserID = visit.UserID
	}
	if visitNew.VisitedAt == 0 {
		visitNew.VisitedAt = visit.VisitedAt
	}
	if visitNew.Mark == 0 {
		visitNew.Mark = visit.Mark
	}

	SetVisit(visitNew)

	return 1
}
