package routes

import "github.com/gin-gonic/gin"

// Api is the function that defines all the routes for the api.
func Api(g *gin.Engine) (*gin.Engine, error) {
	g.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	return g, nil
}