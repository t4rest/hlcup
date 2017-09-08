package main

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"highload/models"
	"strconv"
	"fmt"
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

		fmt.Println(err.Error())
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
		fmt.Println(err.Error())
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	ctx.SetBody(response)
}

func CreateLocation(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	location := models.Location{}
	var err error
	// check params
	err = easyjson.Unmarshal(ctx.PostBody(), &location)

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

	if !models.ValidateLocationParams(params, "insert") {
		fmt.Println("location ValidateLocationParams - insert")

		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	models.SetLocation(location)

	ctx.SetBody([]byte("{}"))
}

func UpdateLocation(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	param := ctx.UserValue("id")
	var conditions []models.Condition
	var location models.Location

	if param == nil {
		fmt.Println("location update param nil")
		ctx.Error("", fasthttp.StatusBadRequest)
	}

	if param == "new" {
		CreateLocation(ctx)
		return
	}

	strId, ok := param.(string)
	if !ok {
		fmt.Println("location update strId != ok")
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

	location, err = models.GetLocation(id)
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
		fmt.Println(err.Error())

		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	if !models.ValidateUserParams(params, "update") {
		fmt.Println("ValidateUserParams update")
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

	models.UpdateLocation(&location, params, conditions)

	ctx.SetBody([]byte("{}"))
}
