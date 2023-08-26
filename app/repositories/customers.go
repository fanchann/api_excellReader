package repositories

import "github.com/fanchann/excelReader/app/domain/models"

type ICustomerRepository interface {
	GetAll() []models.Customers
	BulkInsert(data []models.Customers) ([]models.Customers, error)
	FindById(id uint) (models.Customers, error)
}
