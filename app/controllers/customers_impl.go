package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/fanchann/excelReader/app/domain/types"
	"github.com/fanchann/excelReader/app/services"
	"github.com/fanchann/excelReader/utils"
)

type CustomerCntrlImpl struct {
	service services.ICustomerService
}

func NewCustomerCntrlImpl(service services.ICustomerService) CustomerCntrlImpl {
	return CustomerCntrlImpl{service: service}
}

func (cntrl *CustomerCntrlImpl) Route(app *fiber.App) {
	app.Get("/api/customers", cntrl.GetAllData)
	app.Get("/api/customer/:idCustomer", cntrl.FindDataById)
	app.Post("/api/uploads", cntrl.UploadData)
}

func (cntrl *CustomerCntrlImpl) GetAllData(c *fiber.Ctx) error {
	customerData := cntrl.service.GetAllData()
	return c.JSON(types.WebResponse{Status: http.StatusFound, Message: "success get data", Data: customerData})
}

func (cntrl *CustomerCntrlImpl) UploadData(c *fiber.Ctx) error {
	file, err := c.FormFile("document")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(types.WebResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
	}

	newFileName := fmt.Sprintf("%v_%v", utils.DateLogHistory(), file.Filename)

	if errSaveFile := c.SaveFile(file, fmt.Sprintf("./uploads/%s", newFileName)); errSaveFile != nil {
		return c.Status(http.StatusNotAcceptable).JSON(types.WebResponse{Status: http.StatusNotAcceptable, Message: errSaveFile.Error(), Data: nil})
	}

	errReadFile := cntrl.service.FileUpload(newFileName)
	if errReadFile != nil {
		os.Remove(filepath.Join("uploads", newFileName))
		return c.Status(http.StatusNotAcceptable).JSON(types.WebResponse{Status: http.StatusNotAcceptable, Message: errReadFile.Error(), Data: nil})
	}

	return c.SendStatus(http.StatusOK)
}

func (cntrl *CustomerCntrlImpl) FindDataById(c *fiber.Ctx) error {
	idCustomer := c.Params("idCustomer")
	id, _ := strconv.Atoi(idCustomer)
	customerData, errCustomerNotFound := cntrl.service.FindDataById(uint(id))
	if errCustomerNotFound != nil {
		return c.JSON(types.WebResponse{Status: http.StatusNotFound, Message: errCustomerNotFound.Error(), Data: nil})
	}
	return c.JSON(types.WebResponse{Status: http.StatusFound, Message: "data found", Data: customerData})
}
