package repository

import (
	"Vendor_Management_System/config"
	"Vendor_Management_System/models"
)

func CreateUser(vendor *models.Vendor) error {
	return config.DB.Create(vendor).Error

}
