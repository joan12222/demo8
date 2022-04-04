package api

import (
	"cloudnative/srv"
	"fmt"
	"net/http"
)

type signUpReq struct {
	Email             string
	Password          string
	ConfirmedPassword string
}

func SignUp(ctx *srv.Context) {
	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf(ctx.W, "err: %v", err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("write repsonse error", err)
	}
}
