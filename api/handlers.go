package api

import (
	"net/http"
	"encoding/json"

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

func GetJSONHandler(pageName string, apiURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetch JSON data from the API endpoint
		resp, err := http.Get(apiURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
			return
		}
		defer resp.Body.Close()

		// Decode the JSON response
		var data interface{}
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode JSON"})
			return
		}

		// Render the HTML page with the specified name and pass the JSON data
		c.HTML(http.StatusOK, pageName, gin.H{"data": data})
	}
}

func (server *Server) handlerWrapper(handler func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c)
		c.Next()
	}
}
