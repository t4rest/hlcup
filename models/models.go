package models

import (
	"errors"
	"os"
	"strconv"
	"sync"
	"sort"
)

var (
	NotFound error = errors.New("NotFound")
	timeNow int
	userVisitMap map[int][]int
	locationVisitMap map[int][]int
	mutexUserVisit = &sync.RWMutex{}
)


type FloatPrecision5 float32

type UserVisit struct {
	Mark      uint8  `json:"mark"`
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

func GetAverage(id, fromDate, toDate, fromAge, toAge int, gender string) (float64, error) {
	var avg float64 = 0

	var marksSum uint8 = 0
	var markCount int = 0

	for _, sl := range locationVisitMap[id] {

		if fromDate != 0 && sl.Visit.VisitedAt <= fromDate {
			continue
		}

		if toDate != 0 && sl.Visit.VisitedAt >= toDate {
			continue
		}

		if len(gender) != 0 && gender != sl.User.Gender {
			continue
		}

		if fromAge != 0 && fromAge*31557600 >= timeNow-sl.User.BirthDate {
			continue
		}

		if toAge != 0 && toAge <= int((timeNow-sl.User.BirthDate)/31557600) {
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

	for _, sl := range userVisitMap[id] {

		if fromDate != 0 && sl.Visit.VisitedAt <= fromDate {
			continue
		}

		if toDate != 0 && sl.Visit.VisitedAt >= toDate {
			continue
		}

		if len(country) != 0 && sl.Location.Country != country {
			continue
		}

		if toDistance != 0 && toDistance <= sl.Location.Distance {
			continue
		}

		userVisits = append(userVisits, UserVisit{sl.Visit.Mark, sl.Visit.VisitedAt, sl.Location.Place})
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
