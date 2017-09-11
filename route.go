package main

import (
	"database/sql"
	"fmt"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"highload/models"
	"strconv"
)

func AvgVisits(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	var id, fromDate, toDate, fromAge, toAge int
	var err error

	if ctx.QueryArgs().Has("fromDate") {
		fromDate, err = ctx.QueryArgs().GetUint("fromDate")

		if err != nil {
			ctx.Error("", fasthttp.StatusBadRequest)
			return
		}
	}

	if ctx.QueryArgs().Has("toDate") {
		toDate, err = ctx.QueryArgs().GetUint("toDate")

		if err != nil {
			ctx.Error("", fasthttp.StatusBadRequest)
			return
		}
	}
	if ctx.QueryArgs().Has("fromAge") {
		fromAge, err = ctx.QueryArgs().GetUint("fromAge")

		if err != nil {
			ctx.Error("", fasthttp.StatusBadRequest)
			return
		}
	}
	if ctx.QueryArgs().Has("toAge") {
		toAge, err = ctx.QueryArgs().GetUint("toAge")

		if err != nil {
			ctx.Error("", fasthttp.StatusBadRequest)
			return
		}
	}

	var gender = (string)(ctx.QueryArgs().Peek("gender"))

	var avg float64
	id, err = strconv.Atoi(ctx.UserValue("id").(string))

	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	if gender != "" && !(gender == "m" || gender == "f") {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	_, err = models.GetLocation(int32(id))

	if err == models.NotFound {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	avg, err = models.GetAverage(id, fromDate, toDate, fromAge, toAge, gender)

	if err == sql.ErrNoRows {
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		return
	}

	ctx.WriteString(fmt.Sprintf("{\"avg\" : %.5f}", avg))
}

func Visits(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	var id, fromDate, toDate, toDistance int
	var err error
	var visits models.UserVisitsSl

	if ctx.QueryArgs().Has("fromDate") {
		fromDate, err = ctx.QueryArgs().GetUint("fromDate")

		if err != nil {
			ctx.Error("", fasthttp.StatusBadRequest)
			return
		}
	}

	if ctx.QueryArgs().Has("toDate") {
		toDate, err = ctx.QueryArgs().GetUint("toDate")

		if err != nil {
			ctx.Error("", fasthttp.StatusBadRequest)
			return
		}
	}

	if ctx.QueryArgs().Has("toDistance") {
		toDistance, err = ctx.QueryArgs().GetUint("toDistance")

		if err != nil {
			ctx.Error("", fasthttp.StatusBadRequest)
			return
		}
	}

	var country = (string)(ctx.QueryArgs().Peek("country"))

	id, err = strconv.Atoi(ctx.UserValue("id").(string))

	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	_, err = models.GetUser(int32(id))

	if err == models.NotFound {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	visits, err = models.SelectVisits(id, fromDate, toDate, toDistance, country)

	if err == sql.ErrNoRows {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	ctx.SetContentType("application/json;charset=utf-8")
	response, err := easyjson.Marshal(visits)
	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	ctx.SetBody(response)
}
