package routes

import (
	"math/rand"
	"github.com/JrSchmidtt/go-react-ssr-boilerplate/src/config"
	"github.com/JrSchmidtt/go-react-ssr-boilerplate/src/models"
	"github.com/gin-gonic/gin"
	gossr "github.com/natewong1313/go-react-ssr"
)

// Web is the function that defines all the routes for the web application.
func Web(g *gin.Engine) (*gin.Engine, error) {
	engine, err := gossr.New(config.SsrConfig)
	if err != nil {
		return nil, err
	}
	g.GET("/", func(c *gin.Context) {
		renderPage("Home.tsx", models.IndexPage{
			Number: rand.Intn(100),
		}, c, engine)
	})
	return g, nil
}

func renderPage(file string, data any, c *gin.Context, e *gossr.Engine) {
	c.Writer.Write(e.RenderRoute(gossr.RenderConfig{
		File:  file,
		Title: config.RenderConfig.Title,
		Props: &data,
	}))
}