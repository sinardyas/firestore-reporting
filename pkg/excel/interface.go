package excel

import (
	"github.com/xuri/excelize/v2"
)

type ExcelEngine interface {
	GetFile() *excelize.File
	SetHeader()
	SaveFile(string)
	Close()
}
