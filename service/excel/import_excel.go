package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	log "github.com/sirupsen/logrus"
)

func LoadFromExcel(filePath string) [][]string {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Error(err)
		return nil
	}
	rows := f.GetRows("Sheet1")
	return rows[1:]
}
