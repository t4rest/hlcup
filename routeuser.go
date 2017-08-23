package main

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"highload/models"
	"strconv"
)

//get user by id
func GetUser(ctx *fasthttp.RequestCtx) {

	param := ctx.UserValue("id")
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

	user, err := models.GetUser(id)
	if err != nil {

		if err == models.NotFound {
			ctx.Error("", fasthttp.StatusNotFound)
			return
		}

		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	ctx.SetContentType("application/json;charset=utf-8")
	response, err := easyjson.Marshal(user)
	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	ctx.SetBody(response)
}

//create new user
func CreateUser(ctx *fasthttp.RequestCtx) {
	user := models.User{}
	var err error
	// check params
	err = easyjson.Unmarshal(ctx.PostBody(), &user)

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

	if !models.ValidateUserParams(params, "insert") {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	models.SetUser(user)

	ctx.SetContentType("application/json;charset=utf-8")
	ctx.SetBody([]byte("{}"))
	ctx.SetConnectionClose()
}

//update user
func UpdateUser(ctx *fasthttp.RequestCtx) {
	param := ctx.UserValue("id")
	var conditions []models.Condition
	var user models.User

	if param == nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	if param == "new" {
		CreateUser(ctx)
		return
	}

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

	user, err = models.GetUser(id)
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

	models.UpdateUser(user, params, conditions)

	ctx.SetContentType("application/json;charset=utf-8")
	ctx.SetBody([]byte("{}"))
	ctx.SetConnectionClose()
}
