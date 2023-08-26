package repositories

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"github.com/fanchann/excelReader/app/domain/models"
)

type CustomersRepoImpl struct {
	db *gorm.DB
}

func NewCustomersRepoImpl(db *gorm.DB) ICustomerRepository {
	return &CustomersRepoImpl{db: db}
}

func (repo *CustomersRepoImpl) GetAll() []models.Customers {
	var customers []models.Customers
	repo.db.Select("customer_id", "customer_name").Find(&customers)
	return customers
}
func (repo *CustomersRepoImpl) BulkInsert(data []models.Customers) ([]models.Customers, error) {
	if err := repo.db.Create(&data).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return []models.Customers{}, errors.New("duplicate data detected")
		}

		return []models.Customers{}, err
	}

	return data, nil
}

func (repo *CustomersRepoImpl) FindById(id uint) (models.Customers, error) {
	var customer models.Customers
	if err := repo.db.First(&customer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Customers{}, errors.New("data not found")
		}
		return models.Customers{}, err
	}
	return customer, nil
}
