package utils

import (
	"github.com/JrSchmidtt/go-react-ssr-boilerplate/src/config"
	"github.com/gin-gonic/gin"
	gossr "github.com/natewong1313/go-react-ssr"
)

func RenderPage(file string, data any, e *gossr.Engine) gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Writer.Write(e.RenderRoute(gossr.RenderConfig{
			File:  file,
			Title: config.RenderConfig.Title,
			Props: data,
		}))
	}
}