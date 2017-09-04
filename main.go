package main

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {

	err := ImportDataFromZip()

	if err != nil {
		println(err.Error())
	}

	router := fasthttprouter.New()

	router.GET("/users/:id", GetUser)
	router.POST("/users/:id", UpdateUser)

	router.GET("/locations/:id", GetLocation)
	router.POST("/locations/:id", UpdateLocation)

	router.GET("/visits/:id", GetVisit)
	router.POST("/visits/:id", UpdateVisit)

	router.GET("/users/:id/visits", Visits)
	router.GET("/locations/:id/avg", AvgVisits)

	log.Fatal(fasthttp.ListenAndServe(":80", router.Handler))
}
