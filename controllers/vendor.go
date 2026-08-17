package controllers

import (
	"Vendor_Management_System/config"
	"Vendor_Management_System/models"

	"github.com/gin-gonic/gin"
)

func CreateVendor(c *gin.Context) {

	var vendor models.Vendor

	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := config.DB.Create(&vendor)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Vendor created successfully",
		"vendor":  vendor,
	})
}

func GetVendors(c *gin.Context) {

	var vendors []models.Vendor

	result := config.DB.Find(&vendors)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"vendors": vendors,
	})
}

func GetVendorByID(c *gin.Context) {

	id := c.Param("id")

	var vendor models.Vendor

	result := config.DB.First(&vendor, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"error": "Vendor not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"vendor": vendor,
	})
}

func UpdateVendor(c *gin.Context) {

	id := c.Param("id")

	var vendor models.Vendor

	result := config.DB.First(&vendor, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"error": "Vendor not found",
		})
		return
	}

	var updateData struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
		City  string `json:"city"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	vendor.Name = updateData.Name
	vendor.Email = updateData.Email
	vendor.Phone = updateData.Phone
	vendor.City = updateData.City

	result = config.DB.Save(&vendor)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Vendor updated successfully",
		"vendor":  vendor,
	})
}

func DeleteVendor(c *gin.Context) {

	id := c.Param("id")

	var vendor models.Vendor

	result := config.DB.First(&vendor, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"error": "Vendor not found",
		})
		return
	}

	result = config.DB.Delete(&vendor)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Vendor deleted successfully",
	})
}
