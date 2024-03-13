package main

import (
        "fmt"
        "log"
        "net/http"
        "github.com/gin-gonic/gin"
        "golang.org/x/exp/slog"
)

func main() {
        fmt.Println("Running http server")
        slog.Info("info")
        slog.Warn("warn")
        slog.Debug("debug")

        router := gin.Default()
        router.GET("/hello-world", getHelloWorld)
        log.Fatal(router.Run(":8080"))
}

func getHelloWorld(c *gin.Context) {
        fmt.Println("received /hello-world request\n")
        c.String(http.StatusOK,"Hello, World!")
}

