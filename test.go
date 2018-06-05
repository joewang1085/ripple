package service

import (
	"fmt"

	"github.com/ripple"
)

func Test(ctx *ripple.Context) {

	t := ctx.Newparam["test"].(string)
	fmt.Println("ripple param test is ", t)

}
