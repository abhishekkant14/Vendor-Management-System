package routes

import (
	"Vendor_Management_System/controllers"

	"github.com/gin-gonic/gin"
)

func VendorRoutes(router *gin.Engine) {

	router.POST("/vendors", controllers.CreateVendor)

	router.GET("/vendors", controllers.GetVendors)

	router.GET("/vendors/:id", controllers.GetVendorByID)

	router.PUT("/vendors/:id", controllers.UpdateVendor)

	router.DELETE("/vendors/:id", controllers.DeleteVendor)
}
