package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getDevices(ctx *gin.Context) {
	devices, err := server.queries.GetAllDevices(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, devices)
}

func (server *Server) getTop100Devices(ctx *gin.Context) {
	devices, err := server.queries.GetLastXDevices(ctx, 100)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, devices)
}
