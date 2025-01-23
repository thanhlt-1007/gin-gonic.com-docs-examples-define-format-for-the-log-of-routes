package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
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

    engine.GET("/ping", getPing)

    engine.Run()
}
