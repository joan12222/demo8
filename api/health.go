package api

import (
	"cloudnative/srv"
	"net/http"
)

// functions user need to write
func HealthCheck(c *srv.Context) {
	resp_json := &commonResponse{
		Data: "healthy",
	}
	_ = c.WriteJson(http.StatusOK, resp_json)
}
