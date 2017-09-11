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
	ctx.SetContentType("application/json;charset=utf-8")

	location := &models.Location{}
	var err error
	// check params
	err = easyjson.Unmarshal(ctx.PostBody(), location)

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

	ctx.SetBody([]byte("{}"))
}

func UpdateLocation(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json;charset=utf-8")

	param := ctx.UserValue("id")
	var location *models.Location

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

	if !models.ValidateUserParams(params, "update") {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	models.UpdateLocation(location, params)

	ctx.SetBody([]byte("{}"))
}
