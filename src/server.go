package server

import (
	"github.com/JrSchmidtt/go-react-ssr-boilerplate/src/routes"
	"github.com/JrSchmidtt/go-react-ssr-boilerplate/src/utils"
	"github.com/gin-gonic/gin"
)

func Run() (error) {
	node_modules := utils.CheckIfExistNodeModules()
	if !node_modules {
		utils.NpmInstall()	
	}
	router := gin.Default()
	router.StaticFile("favicon.ico", "./public/favicon.ico")
	router.Static("/assets", "./public")
	router, err := routes.Web(router)
	if err != nil {
		return err
	}
	router, err = routes.Api(router)
	if err != nil {
		return err
	}
	if err := router.Run(":8080"); err != nil {
		return err
	}
	return nil
}