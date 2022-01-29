package helper

import (
	"encoding/json"
	"net/http"
)

// create by td

const (
	ErrorCode      = -1
	SuccessfulCode = 1
)

func NewResultData() *ResultData {
	return new(ResultData)
}

type ResultData struct {
	StatusCode   int         `json:"statusCode"`
	Data         interface{} `json:"data"`
	ErrorMessage interface{}  `json:"errorMessage"`
}

func (this *ResultData) GetErrorBase() *ResultData {
	return this.SetStatusCode(ErrorCode)
}

func (this *ResultData) GetError(message interface{}) *ResultData {
	return this.GetErrorBase().SetErrorMessage(message)
}

func (this *ResultData) GetSuccessfulBase() *ResultData {
	return this.SetStatusCode(SuccessfulCode)
}

func (this *ResultData) GetSuccessful(data interface{}) *ResultData {
	return this.GetSuccessfulBase().SetData(data)
}

func (this *ResultData) SetStatusCode(code int) *ResultData {
	this.StatusCode = code
	return this
}

func (this *ResultData) GetStatusCode() int {
	return this.StatusCode
}

func (this *ResultData) SetErrorMessage(message interface{}) *ResultData {
	this.ErrorMessage = message
	return this
}

func (this *ResultData) GetErrorMessage() interface{} {
	return this.ErrorMessage
}

func (this *ResultData) SetData(data interface{}) *ResultData {
	this.Data = data
	return this
}

func (this *ResultData) GetDate() interface{} {
	return this.Data
}


// util methods

func TryParseResponseBody(response *http.Response) (*ResultData, error) {
	defer response.Body.Close()

	var result *ResultData
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}