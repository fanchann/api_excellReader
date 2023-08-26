package lib

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"

	"github.com/fanchann/excelReader/app/domain/types"
)

func ExcellReader(fileName string) ([]types.CustomersExcellModel, error) {
	assetFile := filepath.Join("uploads", fileName)
	f, err := excelize.OpenFile(assetFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.Rows("Sheet1")
	if err != nil {
		fmt.Println(err)
		os.Remove(filepath.Join("uploads", fileName))
		return []types.CustomersExcellModel{}, err
	}

	var customers []types.CustomersExcellModel

	// skip first row
	rows.Next()
	//

	for rows.Next() {
		columns, _ := rows.Columns()
		customers = append(customers, types.CustomersExcellModel{
			Customer_Name:  columns[0],
			Customer_Email: columns[1],
		})

	}
	return customers, nil
}
