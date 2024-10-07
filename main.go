package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
        "github.com/gin-gonic/gin"
)

func main() {
        fmt.Println(os.Environ())
        port := os.Getenv("APP_PORT")
        fmt.Printf("Running http server on port :%s", port)
        fmt.Println()

        router := gin.Default()
        router.GET("/hello-world", getHelloWorld)
        log.Fatal(router.Run(":"+port))
}

func getHelloWorld(c *gin.Context) {
        fmt.Println("received /hello-world request\n")
        c.String(http.StatusOK,"Hello, World!")
}

