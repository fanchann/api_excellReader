package models

type Customers struct {
	ID    uint   `gorm:"column:customer_id"`
	Name  string `gorm:"column:customer_name"`
	Email string `gorm:"unique;column:customer_email"`
}
