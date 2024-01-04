package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type excel struct {
	file *excelize.File
}

var _ ExcelEngine = (*excel)(nil)

func NewExcel() ExcelEngine {
	ex := &excel{}

	ex.file = excelize.NewFile()

	return ex
}

func (e *excel) GetFile() *excelize.File {
	return e.file
}

func (e *excel) SetHeader() {
	e.file.SetCellValue("Sheet1", "A1", "Campaign")
	e.file.SetCellValue("Sheet1", "B1", "Tanggal")
	e.file.SetCellValue("Sheet1", "C1", "Harga")
	e.file.SetCellValue("Sheet1", "D1", "Nama")
	e.file.SetCellValue("Sheet1", "E1", "Email")
	e.file.SetCellValue("Sheet1", "F1", "Phone")
	e.file.SetCellValue("Sheet1", "G1", "KTP")
	e.file.SetCellValue("Sheet1", "H1", "Address")
	e.file.SetCellValue("Sheet1", "I1", "Status")

	style, err := e.file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   true,
			Family: "Times New Roman",
			Size:   12,
		},
	})

	if err != nil {
		fmt.Println(err)
	}

	e.file.SetRowStyle("Sheet1", 1, 1, style)
}

func (e *excel) SaveFile(fileName string) {
	if err := e.file.SaveAs(string(fileName) + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}

func (e *excel) Close() {
	if e != nil {
		e.file.Close()
	}
}
