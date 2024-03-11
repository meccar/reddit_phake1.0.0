package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET files from public
// func publicHandler(c *gin.Context) {
// 	http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))).ServeHTTP(c.Writer, c.Request)
// }

// PATHS handlers
// func GetHandler(pageName string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		web.Render(c.Writer, pageName, nil)
// 	}
// }

func GetHandler(pageName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, pageName, nil)
		c.Next()
	}
}

func (server *Server) handlerWrapper(handler func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c)
		c.Next()
	}
}
