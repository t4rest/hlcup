package main

import (
	"database/sql"
	"github.com/valyala/fasthttp"
	"highload/models"
	"strconv"
	"time"
	"github.com/mailru/easyjson"
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

	var avg models.VisitAvg
	var conditions []models.Condition

	id, err = strconv.Atoi(ctx.UserValue("id").(string))

	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	if gender != "" && !(gender == "m" || gender == "f") {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	idCondition := models.Condition{
		Param:         "location",
		Value:         strconv.Itoa(id),
		Operator:      "=",
		JoinCondition: "and",
	}
	conditions = append(conditions, idCondition)

	_, err = models.GetLocation(int32(id))

	if err == models.NotFound {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	if fromDate > 0 {
		conditions = append(conditions, models.Condition{
			Param:         "visited_at ",
			Value:         strconv.Itoa(fromDate),
			Operator:      ">",
			JoinCondition: "and",
		})
	}

	if toDate > 0 {
		conditions = append(conditions, models.Condition{
			Param:         "visited_at ",
			Value:         strconv.Itoa(toDate),
			Operator:      "<",
			JoinCondition: "and",
		})
	}

	if fromAge > 0 {
		conditions = append(conditions, models.Condition{
			Param:         "birth_date ",
			Value:         strconv.Itoa(int(time.Now().AddDate(-fromAge, 0, 0).Unix())),
			Operator:      "<",
			JoinCondition: "and",
		})
	}

	if toAge > 0 {
		conditions = append(conditions, models.Condition{
			Param:         "birth_date ",
			Value:         strconv.Itoa(int(time.Now().AddDate(-toAge, 0, 0).Unix())),
			Operator:      ">",
			JoinCondition: "and",
		})
	}

	if len(gender) > 0 {
		conditions = append(conditions, models.Condition{
			Param:         "gender ",
			Value:         "'" + gender + "'",
			Operator:      "=",
			JoinCondition: "and",
		})
	}

	avg, err = models.GetAverage(conditions)

	if err == sql.ErrNoRows {
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		return
	}

	response, err := easyjson.Marshal(avg)
	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	ctx.SetBody(response)
}

func Visits(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	var id int
	var fromDate, toDate, toDistance int
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

 	var conditions []models.Condition

	id, err = strconv.Atoi(ctx.UserValue("id").(string))

	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	idCondition := models.Condition{
		Param:         "user",
		Value:         strconv.Itoa(id),
		Operator:      "=",
		JoinCondition: "and",
	}
	conditions = append(conditions, idCondition)

	_, err = models.GetUser(int32(id))

	if err == models.NotFound {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	if fromDate > 0 {
		conditions = append(conditions, models.Condition{
			Param:         "visited_at ",
			Value:         strconv.Itoa(fromDate),
			Operator:      ">",
			JoinCondition: "and",
		})
	}

	if toDate > 0 {
		conditions = append(conditions, models.Condition{
			Param:         "visited_at ",
			Value:         strconv.Itoa(toDate),
			Operator:      "<",
			JoinCondition: "and",
		})
	}

	if len(country) > 0 {
		conditions = append(conditions, models.Condition{
			Param:         "country",
			Value:         "'" + country + "'",
			Operator:      "=",
			JoinCondition: "and",
		})
	}

	if toDistance > 0 {
		conditions = append(conditions, models.Condition{
			Param:         "distance",
			Value:         strconv.Itoa(toDistance),
			Operator:      "<",
			JoinCondition: "and",
		})
	}

	visits, err = models.SelectVisits(conditions, models.Sort{Fields: []string{"visited_at"}, Direction: "asc"})

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
