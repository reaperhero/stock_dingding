package server

import (
	"fmt"
	"github.com/reaperhero/stock_dingding/model"
	"github.com/reaperhero/stock_dingding/model/repository"
	"github.com/reaperhero/stock_dingding/service/excel"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func SpilderRun() {
}

func ImportExcelToDB() {
	rows := excel.LoadFromExcel("./service/excel/example/20220530.xlsx")
	for _, row := range rows[1:] {
		lineSlice := make([]interface{}, len(row))
		for k, v := range row {
			lineSlice[k] = v
		}
		setNumToFloat([]int{3, 4, 7, 8, 9, 10, 12, 14, 15, 16, 17, 18, 19, 20, 22, 25, 28, 33, 34, 35, 36, 37, 38, 39, 40, 41}, lineSlice)

		// 万
		setNumToWan([]int{5, 6, 21, 23, 24, 26, 27}, lineSlice)
		// 亿
		setNumToYi([]int{11, 29, 30, 31, 32}, lineSlice)

		fmt.Println(lineSlice)
		now := time.Now()
		err := repository.Repository.CreateStockPriceRanking(model.StockPriceRanking{
			CreateTime:          now,
			StockCode:           lineSlice[1].(string),
			StockName:           lineSlice[2].(string),
			CurentPrice:         lineSlice[3].(float64),
			IncreasePrecent:     lineSlice[4].(float64),
			QuantumTotal:        lineSlice[5].(float64),
			QuantumCurrent:      lineSlice[6].(float64),
			BuyPrice:            lineSlice[7].(float64),
			SellPrice:           lineSlice[8].(float64),
			IncreaseSpeed:       lineSlice[9].(float64),
			Turnover:            lineSlice[10].(float64),
			Amount:              lineSlice[11].(float64),
			Pe:                  lineSlice[12].(float64),
			Subordinate:         lineSlice[13].(string),
			Highest:             lineSlice[14].(float64),
			Minimum:             lineSlice[15].(float64),
			OpenPrice:           lineSlice[16].(float64),
			Appoint:             lineSlice[17].(float64),
			YesterdayClosePrice: lineSlice[18].(float64),
			Amplitude:           lineSlice[19].(float64),
			Magnitude:           lineSlice[20].(float64),
			Committee:           lineSlice[21].(float64),
			Average:             lineSlice[22].(float64),
			InsideDish:          lineSlice[23].(float64),
			OutsideDish:         lineSlice[24].(float64),
			SideThan:            lineSlice[25].(float64),
			BuyQuantity:         lineSlice[26].(float64),
			SellQuantity:        lineSlice[27].(float64),
			PriceToBook:         lineSlice[28].(float64),
			TotalEquity:         lineSlice[29].(float64),
			TotalMarketValue:    lineSlice[30].(float64),
			CirculationCapital:  lineSlice[31].(float64),
			CurrentMarket:       lineSlice[32].(float64),
			ThreeDaysUp:         lineSlice[33].(float64),
			SixDaysUp:           lineSlice[34].(float64),
			ThreeDaysChange:     lineSlice[35].(float64),
			SixDaysChange:       lineSlice[36].(float64),
			EvenNumberDays:      lineSlice[37].(float64),
			UpThisMonth:         lineSlice[38].(float64),
			UpThisYear:          lineSlice[39].(float64),
			UpLastMonth:         lineSlice[40].(float64),
			UpPastYear:          lineSlice[41].(float64),
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func setNumToFloat(indexs []int, s []interface{}) {
	for _, index := range indexs {
		reg, _ := regexp.Compile(`\d$`)
		switch {
		case reg.MatchString(s[index].(string)):
			v, _ := strconv.ParseFloat(s[index].(string), 64)
			s[index] = v
		case strings.Contains(s[index].(string), " —"):
			s[index] = 0.0
		}
	}
}

// []int{5, 6, 21, 23, 24, 26, 27}
func setNumToWan(indexs []int, s []interface{}) {
	for _, index := range indexs {
		lineY := s[index].(string)
		switch {
		case regexp.MustCompile(`万$`).MatchString(lineY):
			i, _ := strconv.ParseFloat(lineY[:len(lineY)-3], 64)
			s[index] = i
		case regexp.MustCompile(`亿$`).MatchString(lineY):
			i, _ := strconv.ParseFloat(lineY[:len(lineY)-3], 64)
			s[index] = i * (math.Pow(10, 4))
		default:
			i, _ := strconv.ParseFloat(lineY, 64)
			s[index] = i / math.Pow(10, 4)
		}
	}
}

func setNumToYi(indexs []int, s []interface{}) {
	for _, index := range indexs {
		lineY := s[index].(string)
		switch {
		case regexp.MustCompile(`万$`).MatchString(lineY):
			i, _ := strconv.ParseFloat(lineY[:len(lineY)-3], 64)
			s[index] = i / (math.Pow(10, 4))
		case regexp.MustCompile(`亿$`).MatchString(lineY):
			i, _ := strconv.ParseFloat(lineY[:len(lineY)-3], 64)
			s[index] = i
		default:
			i, _ := strconv.ParseFloat(lineY, 64)
			s[index] = i / (math.Pow(10, 8))
		}
	}
}
