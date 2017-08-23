package main

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"highload/models"
	"strconv"
)

func GetLocation(ctx *fasthttp.RequestCtx) {
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

	location, err := models.GetLocation(id)
	if err != nil {

		if err == models.NotFound {
			ctx.Error("", fasthttp.StatusNotFound)
			return
		}

		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	ctx.SetContentType("application/json;charset=utf-8")
	response, err := easyjson.Marshal(location)
	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	ctx.SetBody(response)
}

func CreateLocation(ctx *fasthttp.RequestCtx) {
	location := models.Location{}
	var err error
	// check params
	err = easyjson.Unmarshal(ctx.PostBody(), &location)

	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}
	var params map[string]interface{}

	err = json.Unmarshal(ctx.PostBody(), &params)
	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	if !models.ValidateLocationParams(params, "insert") {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	models.SetLocation(location)

	ctx.SetContentType("application/json;charset=utf-8")
	ctx.SetBody([]byte("{}"))
	ctx.SetConnectionClose()
}

func UpdateLocation(ctx *fasthttp.RequestCtx) {
	param := ctx.UserValue("id")
	var conditions []models.Condition
	var location models.Location

	if param == nil {
		ctx.Error("", fasthttp.StatusBadRequest)
	}

	if param == "new" {
		CreateLocation(ctx)
		return
	}

	strId, ok := param.(string)
	if !ok {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseInt(strId, 10, 32)
	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	id := int32(id64)

	location, err = models.GetLocation(id)
	if err != nil {

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

	if !models.ValidateUserParams(params, "insert") {
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

	models.UpdateLocation(location, params, conditions)

	ctx.SetContentType("application/json;charset=utf-8")
	ctx.SetBody([]byte("{}"))
	ctx.SetConnectionClose()
}
