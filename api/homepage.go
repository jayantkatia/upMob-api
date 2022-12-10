package api

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shurcooL/github_flavored_markdown"
)

func (server *Server) homePage(ctx *gin.Context) {
	mdFile, err := ioutil.ReadFile("README.md")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	mdHTML := github_flavored_markdown.Markdown(mdFile)
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write([]byte(mdHTML))
	return
}
