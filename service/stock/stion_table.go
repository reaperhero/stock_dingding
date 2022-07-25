package stock

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	defaultTimeLayout = "2006-01-02"
)

type reportTable struct {
	table  *table.Table
	calDay int
}

func newReportTable(calDay int) (*reportTable, error) {
	title := []string{"行业", "代码", "名称", "涨幅统计"}
	for k := calDay - 1; k >= 0; k-- {
		title = append(title, fmt.Sprintf("%d", k))
	}
	tab, err := gotable.Create(title...)
	if err != nil {
		return nil, err
	}
	return &reportTable{table: tab, calDay: calDay}, nil
}

func (r *reportTable) AddRow(s Siton) {
	l := len(s.Increases)
	switch {
	case l >= r.calDay:
		row := []string{s.Subordinate, s.StockCode, s.StockName, fmt.Sprintf("%.1f%%", s.IncreasePrecent)}
		for k := r.calDay; k > 0; k-- {
			row = append(row, fmt.Sprintf("%.1f", s.Increases[l-k]))
		}

		err := r.table.AddRow(row)
		if err != nil {
			log.Errorf("[reportTable] AddRow calDay r.table.AddRow %v", err)
			return
		}
	case l < r.calDay:
		row := []string{s.Subordinate, s.StockCode, s.StockName, fmt.Sprintf("%.1f%%", s.IncreasePrecent)}
		for k := r.calDay - l; k > 0; k-- {
			row = append(row, " ")
		}
		for _, increase := range s.Increases {
			row = append(row, fmt.Sprintf("%.1f", increase))
		}

		err := r.table.AddRow(row)
		if err != nil {
			log.Errorf("[reportTable] AddRow Increases r.table.AddRow %v", err)
			return
		}
	}
}

func (r *reportTable) writeTofile() {

	file, _ := os.OpenFile(fmt.Sprintf(fileRtName, time.Now().Format("2006-01-02")), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer func() {
		_ = file.Close()
	}()
	_, err := file.WriteString(r.table.String())
	if err != nil {
		log.Errorf("[reportTable] writeTofile write.WriteString %v", err)
		return
	}
}
