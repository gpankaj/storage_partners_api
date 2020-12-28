package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gpankaj/storage_partners_api/logger"
)

var (
	//With small case, router is private variable and is availble only within app package.
	router  = gin.Default()
)

type frameworkInterface interface {

}

func StartApplication() {
	mapUrls()

	logger.Info("About to start application")
	router.Run(":8080")
}
