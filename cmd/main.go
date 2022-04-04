package main

import (
	"cloudnative/api"
	"cloudnative/srv"
	"net/http"
)

func main() {
	myServer := srv.NewHttpServer("Demo Http Server")
	myServer.Route(http.MethodGet, "/healthz", api.HealthCheck)
	myServer.Route(http.MethodPost, "/signup", api.SignUp)
	myServer.Start(":8080")
}

type commonResponse struct {
	Data interface{}
}
