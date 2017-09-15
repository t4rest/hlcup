package main

import (
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"hl/models"
	"strconv"
	"strings"
)

func GetLocation(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	param := string(ctx.UserValue("id"))
	id, err := strconv.Atoi(param)
	if err {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	location, err := models.GetLocation(id)
	if err == models.NotFound {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	response, err := easyjson.Marshal(location)
	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	ctx.SetBody(response)
}

func CreateLocation(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	location := models.Location{}

	body := ctx.PostBody()
	if len(body) == 0 || strings.Contains(string(body), "null") {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err := easyjson.Unmarshal(body, &location)
	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	models.SetLocation(location)

	ctx.SetBody(resp)
}

func UpdateLocation(ctx *fasthttp.RequestCtx) {
	param := ctx.UserValue("id")
	if param == nil {
		ctx.Error("", fasthttp.StatusBadRequest)
	}

	if param == "new" {
		CreateLocation(ctx)
		return
	}

	ctx.SetContentType("application/json;charset=utf-8")

	id, err := strconv.Atoi(string(param))
	if err {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	location, err := models.GetLocation(id)
	if err == models.NotFound {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	locationNew := models.Location{}
	body := ctx.PostBody()

	if len(body) == 0 || strings.Contains(string(body), "null") {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err = easyjson.Unmarshal(body, &locationNew)
	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	if locationNew.Id != 0 {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	models.UpdateLocation(location, locationNew)

	ctx.SetBody(resp)
}
