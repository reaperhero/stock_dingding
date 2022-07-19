package stock

import (
	"encoding/json"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/reaperhero/stock_dingding/model"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"os"
	"time"
)

var (
	fileRtName = "situation/situation_report_%s.txt"
	fileRdName = "situation/situation_record.txt"
)

type Siton struct {
	StockCode       string    `json:"stock_code"`
	StockName       string    `json:"stock_name"`
	Increases       []float64 `json:"increases"`
	IncreasePrecent float64   `json:"-"`
	Subordinate     string    `json:"subordinate"`
}

type SitonManage struct {
	M *hashmap.Map
}

func NewSitonManage() *SitonManage {
	sitonManage := &SitonManage{M: hashmap.New()}

	file, err := ioutil.ReadFile(fileRdName)
	if err != nil {
		return sitonManage
	}

	var stions []Siton
	err = json.Unmarshal(file, &stions)
	if err != nil {
		log.Errorf("[NewSitonManage] json.Unmarshal %v", err)
		return sitonManage
	}

	for _, siton := range stions {
		sitonManage.M.Put(siton.StockCode, siton)
	}

	return sitonManage
}

// 获取situation_record.txt文件中的stock code
func (s *SitonManage) GetStockSet() []string {
	ss := make([]string, 0, 30)
	for _, i := range s.M.Keys() {
		ss = append(ss, i.(string))
	}
	return ss
}

func (s *SitonManage) AddTodayStock(ss []model.Stock) {
	for _, stock := range ss {
		v, ok := s.M.Get(stock.StockCode)
		if ok {
			hisSiton := v.(Siton)
			hisSiton.Increases = append(hisSiton.Increases, stock.IncreasePrecent)
			s.M.Put(stock.StockCode, hisSiton)
			continue
		}
		s.M.Put(stock.StockCode, Siton{
			StockCode:   stock.StockCode,
			StockName:   stock.StockName,
			Increases:   []float64{stock.IncreasePrecent},
			Subordinate: stock.Subordinate,
		})
	}
}
func (s *SitonManage) RecordFile(reci int) {
	var res = make([]Siton, 0, 100)
	for _, v := range s.M.Values() {
		s, ok := v.(Siton)
		lDay := len(s.Increases)
		if ok && lDay > reci {
			s.Increases = s.Increases[lDay-reci:]
		}
		res = append(res,s)
	}
	file, _ := os.OpenFile(fileRdName, os.O_WRONLY|os.O_CREATE, 0644)
	defer func() {
		_ = file.Close()
	}()

	fileInfo ,_ := file.Stat()
	if fileInfo.ModTime().Day() == time.Now().Day() {
		return
	}

	marshal, err := json.MarshalIndent(res,""," ")
	if err != nil {
		log.Errorf("[SitonManag] RecordFile json.Marshal %v",err)
		return
	}
	file.Write(marshal)
}

func RecordFileWithStions(ss []Siton) {
	file, _ := os.OpenFile(fileRdName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer func() {
		_ = file.Close()
	}()

	marshal, err := json.MarshalIndent(ss,""," ")
	if err != nil {
		log.Errorf("[SitonManag] RecordFile json.Marshal %v",err)
		return
	}
	file.Write(marshal)
}

func (s *SitonManage) ReportFile(calDay int) {
	heap := binaryheap.NewWith(func(a, b interface{}) int {
		as := a.(Siton).IncreasePrecent
		bs := b.(Siton).IncreasePrecent

		if math.Max(as, bs) == as {
			return -1
		}
		if math.Max(as, bs) == bs {
			return 1
		}
		return 0
	})
	for _, v := range s.M.Values() {
		s := v.(Siton)
		var (
			initIncrese float64 = 100
			increaseDay         = len(s.Increases)
		)

		switch {
		case increaseDay >= calDay:
			for k := calDay; k > 0; k-- {
				increase := s.Increases[increaseDay-k]
				initIncrese *= (100 + increase) / 100
			}
		case increaseDay < calDay:
			for _, increase := range s.Increases {
				initIncrese *= (100 + increase) / 100
			}
		}
		s.IncreasePrecent = initIncrese
		heap.Push(s)
	}

	b, err := newReportTable(calDay)
	if err != nil {
		log.Errorf("[SitonManage] reportFile newReportTable %v", err)
		return
	}
	for {
		v, ok := heap.Pop()
		if !ok {
			b.writeTofile()
			return
		}
		b.AddRow(v.(Siton))
	}
}

