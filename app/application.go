package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	//With small case, router is private variable and is availble only within app package.
	router  = gin.Default()
)

type frameworkInterface interface {

}

func StartApplication() {
	router.Use(cors.Default())
	mapUrls()
	router.Run(":8080")
}
