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

func SetVisit(visit Visit, update bool) {
	mutexVisit.Lock()
	visitMap[visit.Id] = visit
	mutexVisit.Unlock()

	if !update {
		SetVisits(visit)
	}
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
		SetVisit(visit, false)
	}
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

	updateUsetLocationVisit(visit)

	SetVisit(visit, true)

	return 1
}