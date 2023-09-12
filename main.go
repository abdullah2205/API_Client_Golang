package main

import (
    "github.com/gin-gonic/gin"
    "API_Client_Golang/routes"
)

func main() {
    router := gin.Default()
    routes.SetupRoutes(router)
    router.Run(":8001")
}
