package main

import (
	"encoding/json"
	"fmt"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"highload/models"
	"strconv"
)

func GetVisit(ctx *fasthttp.RequestCtx) {
	param := ctx.UserValue("id")
	strId, ok := param.(string)

	if !ok {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseInt(strId, 10, 32)
	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	id := int32(id64)

	visit, err := models.GetVisit(id)
	if err != nil {
		fmt.Println(err.Error())

		if err == models.NotFound {
			ctx.Error("", fasthttp.StatusNotFound)
			return
		}

		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	ctx.SetContentType("application/json;charset=utf-8")
	response, err := easyjson.Marshal(visit)
	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	ctx.SetBody(response)
}

func CreateVisit(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	visit := models.Visit{}

	// check params
	err := easyjson.Unmarshal(ctx.PostBody(), &visit)

	if err != nil {
		fmt.Println(err.Error())
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	var params map[string]interface{}

	err = json.Unmarshal(ctx.PostBody(), &params)
	if err != nil {
		fmt.Println(err.Error())
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	if !models.ValidatVsitParams(params, "insert") {
		fmt.Println("visit create validation failed")
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	models.InsertVisit(visit)

	ctx.SetBody([]byte("{}"))
}

func UpdateVisit(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	param := ctx.UserValue("id")
	var conditions []models.Condition
	var visit models.Visit

	if param == nil {
		fmt.Println("visit update param nil")
		ctx.Error("", fasthttp.StatusBadRequest)
	}

	if param == "new" {
		CreateVisit(ctx)
		return
	}

	strId, ok := param.(string)
	if !ok {
		fmt.Println("visit update id != ok")
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseInt(strId, 10, 32)
	if err != nil {
		fmt.Println(err.Error())

		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	id := int32(id64)

	visit, err = models.GetVisit(id)
	if err != nil {
		fmt.Println(err.Error())

		if err == models.NotFound {
			ctx.Error("", fasthttp.StatusNotFound)
			return
		}

		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	var params map[string]interface{}

	err = json.Unmarshal(ctx.PostBody(), &params)
	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	if !models.ValidatVsitParams(params, "update") {
		fmt.Println("visit update validation failed")
		fmt.Println(params)

		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	userIdCondition := models.Condition{
		Param:         "id",
		Value:         strId,
		Operator:      "=",
		JoinCondition: "and",
	}
	conditions = append(conditions, userIdCondition)

	models.UpdateVisit(visit, params, conditions)

	ctx.SetBody([]byte("{}"))
}
