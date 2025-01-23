package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

func getPing(context *gin.Context) {
    context.JSON(
        http.StatusOK,
        gin.H {
            "message": "pong",
        },
    )
}

func main() {
    engine := gin.Default()
    gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
        log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
    }
    engine.GET("/ping", getPing)

    engine.Run()
}
