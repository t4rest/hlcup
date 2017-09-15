package main

import (
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"hl/models"
	"strconv"
	"strings"
)

//get user by id
func GetUser(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	param := ctx.UserValue("id").(string)
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	user, err := models.GetUser(id)
	if err == models.NotFound {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	response, err := easyjson.Marshal(user)
	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	ctx.SetBody(response)
}

//create new user
func CreateUser(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	user := models.User{}

	body := ctx.PostBody()
	if len(body) == 0 || strings.Contains(string(body), "null") {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err := easyjson.Unmarshal(body, &user)
	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	models.SetUser(user)

	ctx.SetBody(resp)
}

//update user
func UpdateUser(ctx *fasthttp.RequestCtx) {
	param := ctx.UserValue("id")
	if param == nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	if param == "new" {
		CreateUser(ctx)
		return
	}

	ctx.SetContentType("application/json;charset=utf-8")

	id, err := strconv.Atoi(param.(string))
	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	user, err := models.GetUser(id)
	if err == models.NotFound {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	userNew := models.User{}
	body := ctx.PostBody()

	if len(body) == 0 || strings.Contains(string(body), "null") {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err = easyjson.Unmarshal(body, &userNew)
	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	if userNew.Id != 0 {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	models.UpdateUser(user, userNew, strings.Contains(string(body), "birth_date"))

	ctx.SetBody(resp)
}
