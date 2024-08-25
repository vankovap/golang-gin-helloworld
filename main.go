package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
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
        log.Fatal(router.Run(os.Getenv("APP_PORT")))
}

func getHelloWorld(c *gin.Context) {
        fmt.Println("received /hello-world request\n")
        c.String(http.StatusOK,"Hello, World!")
}

