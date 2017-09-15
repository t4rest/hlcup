package models

import (
	"sync"
)

type Visit struct {
	Id        int `json:"id"`
	Location  int `json:"location"`
	User      int `json:"user"`
	VisitedAt int `json:"visited_at"`
	Mark      int `json:"mark"`
}

type Visits struct {
	Visits []Visit `json:"visits"`
}

var visitMap = make(map[int]Visit)
var mutexVisit = &sync.RWMutex{}

func SetVisit(visit Visit) {
	mutexVisit.Lock()
	visitMap[visit.Id] = visit
	mutexVisit.Unlock()
}

func GetVisit(id int) (Visit, error) {
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

func InsertVisit(visit Visit) {
	SetVisit(visit)
}

func UpdateVisit(visit Visit, visitNew Visit) int {

	if visitNew.Location != 0 {
		visit.Location = visitNew.Location
	}

	if visitNew.User != 0 {
		visit.User = visitNew.User
	}

	if visitNew.VisitedAt != 0 {
		visit.VisitedAt = visitNew.VisitedAt
	}

	if visitNew.Mark != 0 {
		visit.Mark = visitNew.Mark
	}

	SetVisit(visit)

	return 1
}
