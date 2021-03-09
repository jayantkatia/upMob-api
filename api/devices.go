package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getDevices(ctx *gin.Context) {
	devices, err := server.queries.GetDevices(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, devices)
}
