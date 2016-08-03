package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
	"taxpanel/webserver/controllers/core"
	"taxpanel/webserver/controllers/status"
	"taxpanel/webserver/validation/errors"
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
	Router.PUT("/expensives/:id/active", core.ActiveExpensive)
	Router.PUT("/expensives/:id/inactive", core.InactiveExpensive)
	Router.DELETE("/expensives/:id", core.DeleteExpensive)
	Router.GET("/expensives", core.Expensives)
	Router.POST("/expensives", core.CreateExpensive)

	Router.GET("/users", core.Users)
	Router.GET("/users/:id", core.UserById)
	Router.POST("/users", core.CreateUser)

	Router.GET("/logs", core.Logs)
	Router.GET("/logs/:id", core.LogById)
	Router.POST("/logs", core.CreateLog)

	Router.GET("/salary", core.Salary)

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
