package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
	"tablero/webserver/controllers/core"
	"tablero/webserver/controllers/status"
	"tablero/webserver/validation/errors"
)

var Router *gin.Engine

func configureRuntime() {
	//Sets the max quantity of cores
	coreNum := runtime.NumCPU()
	runtime.GOMAXPROCS(coreNum)
}

func configureRouter() {
	//Stats the gin engine
	Router = gin.Default()
	Router.NoRoute(noRouteHandler())
	Router.RedirectFixedPath = false
	Router.RedirectTrailingSlash = false
}

// export the starting of the application
func StartApp() {
	configureRouter()
	configureRuntime()

	// mappings
	Router.GET("/ping", status.Ping)
	Router.GET("/", status.Ok)
	Router.GET("/expensives/:id", core.ExpensiveById)
	Router.GET("/expensives", core.Expensives)
	Router.POST("/expensives", core.CreateExpensive)

	Router.Run(":8080")
}

func main() {
	StartApp()
}

// handles resource not found error
func noRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		errors.RespondError(c, errors.NotFoundApiError(fmt.Sprintf("Resource %s not found.", c.Request.URL.Path)))
	}
}
