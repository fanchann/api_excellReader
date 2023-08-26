package services

import "github.com/fanchann/excelReader/app/domain/types"

type ICustomerService interface {
	GetAllData() []types.CustomerResponse
	FindDataById(id uint) (types.CustomerResponse, error)
	FileUpload(fileName string) error
}
