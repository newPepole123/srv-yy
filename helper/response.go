package helper

// create by td

import (
	"errors"
	"github.com/webpkg/web"
)

const (
	HttpStatusCodeBadRequest = 400
)


func FunnilyReturnOK(ctx *web.Context, data interface{}) (web.Data, error) {

	if !ReturnOK(ctx, data) {
		return nil, errors.New("helper.ReturnOK(ctx, data) invoke fail")
	}

	return nil, nil
}

func FunnilyReturnSimpleOK(ctx *web.Context) (web.Data, error) {

	if !ReturnSimpleOK(ctx) {
		return nil, errors.New("helper.ReturnSimpleOK(ctx) invoke fail")
	}

	return nil, nil
}

func FunnilyReturnErrorString(ctx *web.Context, message string) (web.Data, error) {

	if !ReturnErrorString(ctx, message) {
		return nil, errors.New("helper.ReturnErrorString(ctx, message) invoke fail")
	}

	return nil, nil
}

func FunnilyReturnBadRequestErrorString(ctx *web.Context, message string) (web.Data, error) {
	if !ReturnBadRequestErrorString(ctx, message) {
		return nil, errors.New("helper.ReturnBadRequestErrorString(ctx, message) invoke fail")
	}
	return nil, nil
}

func FunnilyReturnBadRequestError(ctx *web.Context, err error) (web.Data, error) {
	if !ReturnBadRequestError(ctx, err) {
		return nil, errors.New("helper.ReturnBadRequestError(ctx, err) invoke fail")
	}
	return nil, nil
}

func ReturnSimpleOK(ctx *web.Context) bool {
	return ReturnOK(ctx, nil)
}

func ReturnOK(ctx *web.Context, data interface{}) bool {
	return ReturnData(ctx, data, nil)
}

func ReturnErrorString(ctx *web.Context, message string) bool {
	return ReturnData(ctx, nil, errors.New(message))
}

func ErrorChecker(ctx *web.Context, err error) bool {
	return ReturnData(ctx, nil, err)
}

func ReturnError(ctx *web.Context, err error) bool {
	return ErrorChecker(ctx, err)
}

func ReturnErrorStringAndStatusCode(ctx *web.Context, message string, code int) bool {
	ctx.Status(code)
	return ReturnData(ctx, nil, errors.New(message))
}

func ReturnErrorAndStatusCode(ctx *web.Context, err error, code int) bool {
	ctx.Status(code)
	return ErrorChecker(ctx, err)
}

func ReturnBadRequestErrorString(ctx *web.Context, message string) bool {
	return ReturnErrorStringAndStatusCode(ctx, message, HttpStatusCodeBadRequest)
}

func ReturnBadRequestError(ctx *web.Context, err error) bool {
	return ReturnErrorAndStatusCode(ctx, err, HttpStatusCodeBadRequest)
}

func ReturnData(ctx *web.Context, data interface{}, err error) (flag bool) {
	result := NewResultData().GetSuccessfulBase()
	if err != nil {
		result = result.GetError(err.Error())
		_ = ctx.WriteJSON(result)
		flag = false
	} else {
		result = result.SetData(data)
		_ = ctx.WriteJSON(result)
		flag = false
	}
	return true
}
