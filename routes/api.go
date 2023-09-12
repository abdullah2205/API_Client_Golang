package routes

import (
    "github.com/gin-gonic/gin"
    "API_Client_Golang/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/api/register", controllers.Register)

	router.POST("/api/login", controllers.Login)

	router.GET("/api/user", controllers.GetUser)

	router.GET("/api/buku", controllers.GetBuku)

	router.POST("/api/buku/add", controllers.AddBuku)

	router.GET("/api/buku/:id_buku", controllers.GetBukuByID)

	router.PUT("/api/buku/:id_buku", controllers.UpdateBuku)

	router.DELETE("/api/buku/:id_buku", controllers.DeleteBuku)

	router.POST("/api/logout", controllers.Logout)
}
