package entrypoint

import (
	_ "go-http/docs"
	"net/http"
)

// RootHandler - Returns all the available APIs
// @Summary This API can be used as health check for this application.
// @Description Tells if the chi-swagger APIs are working or not.
// @Tags chi-swagger
// @Accept  json
// @Produce  json
// @Success 200 {string} response "api response"
// @Router / [get]
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome123"))
}
