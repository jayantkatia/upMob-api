package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/jayantkatia/backend_upcoming_mobiles/db/sqlc"
)

type Server struct {
	queries *db.Queries
	router  *gin.Engine
}

func NewServer(queries *db.Queries) *Server {
	server := &Server{queries: queries}
	router := gin.Default()

	router.GET("/devices", server.getDevices)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
