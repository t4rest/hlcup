package models

import (
	"errors"
	"os"
	"strconv"
)

var (
	NotFound error = errors.New("NotFound")
)

var timeNow int

func init() {
	file, err := os.Open("/tmp/data/options.txt")
	if err != nil {
		file, err = os.Open("/Users/Adnrey/go/src/hlcupdocs/data/FULL/data/options.txt")
		if err != nil {
			PanicOnErr(err)
		}
	}

	timestampBytes := make([]byte, 10)
	_, err = file.Read(timestampBytes)
	PanicOnErr(err)

	timeNow, err = strconv.Atoi(string(timestampBytes))
	PanicOnErr(err)

	println("on start")
	println(timeNow)
}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
