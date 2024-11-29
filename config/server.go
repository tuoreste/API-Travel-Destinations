package config

import "github.com/gin-gonic/gin"

// other web apps may interact this API thru server.go
// e.g authentication
// ports, timeouts etc could be defined from here

func SetupServer() *gin.Engine {
    router := gin.Default()
    return router
}
