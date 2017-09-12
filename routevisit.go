package main

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"highload/models"
	"strconv"
)

func GetVisit(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

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

		if err == models.NotFound {
			ctx.Error("", fasthttp.StatusNotFound)
			return
		}

		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	response, err := easyjson.Marshal(visit)
	if err != nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	ctx.SetBody(response)
}

func CreateVisit(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	visit := &models.Visit{}

	// check params
	err := easyjson.Unmarshal(ctx.PostBody(), visit)

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

	if !models.ValidateVsitParams(params, "insert") {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	//go func() {
		models.SetVisit(visit)
	//}()

	ctx.SetBody(resp)
}

func UpdateVisit(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	param := ctx.UserValue("id")
	var visitNew *models.Visit
	var visit *models.Visit

	if param == nil {
		ctx.Error("", fasthttp.StatusBadRequest)
	}

	if param == "new" {
		CreateVisit(ctx)
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

	visit, err = models.GetVisit(id)
	if err != nil {
		if err == models.NotFound {
			ctx.Error("", fasthttp.StatusNotFound)
			return
		}

		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	body := ctx.PostBody()

	err = json.Unmarshal(body, &visitNew)
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

	if !models.ValidateVsitParams(params, "update") {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	//go func() {
		models.UpdateVisit(visit, visitNew)
	//}()

	ctx.SetBody(resp)
}
