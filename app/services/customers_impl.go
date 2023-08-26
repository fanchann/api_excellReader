package services

import (
	"errors"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/fanchann/excelReader/app/domain/models"
	"github.com/fanchann/excelReader/app/domain/types"
	"github.com/fanchann/excelReader/app/repositories"
	"github.com/fanchann/excelReader/lib"
)

type CustomerServiceImpl struct {
	repo  repositories.ICustomerRepository
	cache *cache.Cache
}

func NewCustomerServiceImpl(repo repositories.ICustomerRepository, cache *cache.Cache) ICustomerService {
	return &CustomerServiceImpl{repo: repo, cache: cache}
}

func (service *CustomerServiceImpl) GetAllData() []types.CustomerResponse {
	if cachedData, found := service.cache.Get("allCustomers"); found {
		return cachedData.([]types.CustomerResponse)
	}

	customersData := service.repo.GetAll()
	customers := make([]types.CustomerResponse, len(customersData))

	for i, customer := range customersData {
		customers[i] = modelToTypes(customer)
	}
	service.cache.Set("allCustomers", customers, time.Duration(time.Minute*5))

	return customers

}

func (service *CustomerServiceImpl) FindDataById(id uint) (types.CustomerResponse, error) {
	data, err := service.repo.FindById(id)
	if err != nil {
		return types.CustomerResponse{}, err
	}

	return modelToTypes(data), nil

}

func (service *CustomerServiceImpl) FileUpload(fileName string) error {
	alloWedExt := fileUploadFilter(fileName)
	if !alloWedExt {
		return errors.New("file extension must .xlsx")
	}

	dataCustomer, err := lib.ExcellReader(fileName)
	if err != nil {
		return err
	}

	customerModels := dataCustomerExcelToModel(dataCustomer)
	_, errBulkInsert := service.repo.BulkInsert(customerModels)
	if errBulkInsert != nil {
		return errBulkInsert
	}
	return nil

}

func fileUploadFilter(fileName string) bool {
	if !strings.HasSuffix(fileName, ".xlsx") {
		return false
	}
	return true
}

func modelToTypes(data models.Customers) types.CustomerResponse {
	return types.CustomerResponse{Id: data.ID, Name: data.Name}
}

func dataCustomerExcelToModel(data []types.CustomersExcellModel) []models.Customers {
	customers := make([]models.Customers, len(data))
	for i, item := range data {
		customers[i] = models.Customers{Name: item.Customer_Name, Email: item.Customer_Email}
	}
	return customers
}
