package main

import (
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"hl/models"
	"strconv"
	"strings"
)

func GetVisit(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	param := string(ctx.UserValue("id"))
	id, err := strconv.Atoi(param)
	if err {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	visit, err := models.GetVisit(id)
	if err == models.NotFound {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	response, err := easyjson.Marshal(visit)
	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	ctx.SetBody(response)
}

func CreateVisit(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	visit := models.Visit{}

	body := ctx.PostBody()
	if len(body) == 0 || strings.Contains(string(body), "null") {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err := easyjson.Unmarshal(body, &visit)
	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	models.SetVisit(visit)

	ctx.SetBody(resp)
}

func UpdateVisit(ctx *fasthttp.RequestCtx) {
	param := ctx.UserValue("id")
	if param == nil {
		ctx.Error("", fasthttp.StatusBadRequest)
	}

	if param == "new" {
		CreateVisit(ctx)
		return
	}

	ctx.SetContentType("application/json;charset=utf-8")

	id, err := strconv.Atoi(string(param))
	if err {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	visit, err := models.GetVisit(id)
	if err == models.NotFound {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	visitNew := models.Visit{}
	body := ctx.PostBody()

	if len(body) == 0 || strings.Contains(string(body), "null") {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err = easyjson.Unmarshal(body, &visitNew)
	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	if visitNew.Id != 0 {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	models.UpdateVisit(visit, visitNew)

	ctx.SetBody(resp)
}
