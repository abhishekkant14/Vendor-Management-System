package main

import (
	"fmt"

	"Vendor_Management_System/config"
	"Vendor_Management_System/models"
	"Vendor_Management_System/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	err := config.DB.AutoMigrate(&models.Vendor{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Vendor table migrated successfully")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Vendor Management System API",
		})
	})

	routes.VendorRoutes(router)

	router.Run(":9091")
}
