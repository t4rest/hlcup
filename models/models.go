package models

import (
	"errors"
	"fmt"
	"database/sql"
)

var (
	db       *sql.DB
	err      error
	NotFound error = errors.New("NotFound")
)

type Condition struct {
	Param         string
	Value         string
	Operator      string
	JoinCondition string
}

type Join struct {
	Name      string
	Type      string
	Condition Condition
}

type Sort struct {
	Fields []string
	Direction string
}

type FloatPrecision5 float32

func (f FloatPrecision5) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.5f", f)), nil
}


func init() {
	println("on start")
}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}