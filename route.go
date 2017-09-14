package main

import (
	"database/sql"
	"fmt"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"hl/models"
	"strconv"
)

var resp []byte = []byte("{}")

func AvgVisits(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	var id int
	var avg float64

	fromDate, err := strconv.Atoi(string(ctx.FormValue("fromDate")))
	if len(ctx.FormValue("fromDate")) != 0 && err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	toDate, err := strconv.Atoi(string(ctx.FormValue("toDate")))
	if len(ctx.FormValue("toDate")) != 0 && err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	fromAge, err := strconv.Atoi(string(ctx.FormValue("fromAge")))
	if len(ctx.FormValue("fromAge")) != 0 && err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	toAge, err := strconv.Atoi(string(ctx.FormValue("toAge")))
	if len(ctx.FormValue("toAge")) != 0 && err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	gender := string(ctx.FormValue("gender"))
	if gender != "" && gender != "f" && gender != "m" {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	id, err = strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	_, err = models.GetLocation(id)
	if err == models.NotFound {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	avg, err = models.GetAverage(id, fromDate, toDate, fromAge, toAge, gender)
	if err != nil {
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		return
	}

	ctx.WriteString(fmt.Sprintf("{\"avg\" : %.5f}", avg))
}

func Visits(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	var visits models.UserVisitsSl

	fromDate, err := strconv.Atoi(string(ctx.FormValue("fromDate")))
	if len(ctx.FormValue("fromDate")) != 0 && err != nil {
		ctx.Error("Unsupported path", fasthttp.StatusBadRequest)
		return
	}

	toDate, err := strconv.Atoi(string(ctx.FormValue("toDate")))
	if len(ctx.FormValue("toDate")) != 0 && err != nil {
		ctx.Error("Unsupported path", fasthttp.StatusBadRequest)
		return
	}

	toDistance, err := strconv.Atoi(string(ctx.FormValue("toDistance")))
	if len(ctx.FormValue("toDistance")) != 0 && err != nil {
		ctx.Error("Unsupported path", fasthttp.StatusBadRequest)
		return
	}

	country := string(ctx.FormValue("country"))

	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	_, err = models.GetUser(id)
	if err == models.NotFound {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	visits, err = models.SelectVisits(id, fromDate, toDate, toDistance, country)
	if err == sql.ErrNoRows {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	response, err := easyjson.Marshal(visits)
	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	ctx.SetBody(response)
}
