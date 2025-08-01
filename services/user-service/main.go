package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
    shutdown := InitProvider()
    defer shutdown()

    r := gin.Default()
    r.Use(otelgin.Middleware("user-service"))

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    r.GET("/users", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"users": []string{"Alice", "Bob"}})
    })

    r.Run(":8080")
}