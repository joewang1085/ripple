package controllers

import (
	"service"
	//	"github.com/gopkg.in/mgo.v2/bson"
	"github.com/ripple"
)

type AppController struct {
}

func NewAppController() *AppController {
	output := new(AppController)
	return output
}

func (this *AppController) Get(ctx *ripple.Context) {
}

func (this *AppController) Post(ctx *ripple.Context) {
}

func (this *AppController) Put(ctx *ripple.Context) {
}

func (this *AppController) Get1(ctx *ripple.Context) {
	ctx.Response.Body = 111
}

func (this *AppController) GetTest(ctx *ripple.Context) {
	urlParaNameList := []string{"test"}
	if false == model.ParseBodyParamter(ctx, urlParaNameList) {
		ctx.Response.Body = bson.M{"状态码": "失败", "错误": "参数不合法"}
		return
	}
	service.Test(ctx)
}
