package api

import (
	"cloudnative/srv"
	"net/http"
)

// functions user need to write
func HealthCheck(c *srv.Context) {
	c.WriteJson(http.StatusOK, "Healthy")
}
