package models

type Vendor struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone" binding:"required,len=10"`
	City  string `json:"city" binding:"required"`
}
