package model

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gopkg.in/mgo.v2/bson"
	"github.com/log"
	"github.com/ripple"
)

type BodyParam struct {
	NewParam bson.M `json:"newParam"`
}

func ParseBodyParamter(ctx *ripple.Context, paramNameList []string) bool {
	paramObj := new(BodyParam)
	err := ParseJsonFromUrlBody(ctx.Request.Body, paramObj)
	if nil != err {
		log.Error("file model.newParam.go, func ParseBodyParamter, err :", err, ", ParseJsonFromUrlBody failed")
		return false
	}
	ctx.NewParams = make(bson.M)
	for paramName, paramValue := range paramObj.NewParam {
		if paramValue == "" {
			return false
		}
		ctx.NewParams[paramName] = paramValue
	}
	if false == IsNewParamterValid(paramNameList, ctx.NewParams) {
		return false
	}
	return true
}

func ParseJsonFromUrlBody(body io.ReadCloser, v interface{}) error {
	bodyData, err := ioutil.ReadAll(body)
	if nil != err {
		return err
	}
	err = json.Unmarshal(bodyData, v)
	if nil != err {
		log.Error("file model.newParam.go, func ParseJsonFromUrlBody, err :", err, ", json.Unmarshal failed")
		return err
	}
	return nil
}

func IsNewParamterValid(paramNameList []string, paramMap bson.M) bool {
	for _, paraName := range paramNameList {
		_, ok := paramMap[paraName]
		if !ok {
			return false
		}
	}
	return true
}

func ParseFormParamter(ctx *ripple.Context, paramNameList []string) bool {
	r := ctx.Request
	switch r.Method {
	case "POST":
		err := r.ParseMultipartForm(100000)
		if err != nil {
			log.Error("file model.newParam.go, func ParseFormParamter, err :", err, ", r.ParseMultipartForm failed")
			ctx.Response.Body = bson.M{"状态码": "失败", "错误": http.StatusInternalServerError}
			return false
		}
		//		m = r.MultipartForm
	}
	ctx.NewParams = make(bson.M, 0)
	for paramName, paramValue := range r.MultipartForm.Value {
		ctx.NewParams[paramName] = paramValue
	}
	if false == IsNewParamterValid(paramNameList, ctx.NewParams) {
		return false
	}
	return true
}
