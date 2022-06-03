package server

import (
	"github.com/reaperhero/stock_dingding/service/stock"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "search stock",
		Long:  "search stock",
		Run: func(cmd *cobra.Command, args []string) {
			reportDailyLimitStatisticsStock()
		},
	}
	hanyeSearchCmd = &cobra.Command{
		Use:   "hanye",
		Short: "show every harden",
		Long:  "show every harden",
		Run: func(cmd *cobra.Command, args []string) {
			for _, arg := range args {
				list, err := stock.GetStockBySubordinate(arg)
				if err != nil {
					log.Fatal(err)
				}
				echoStock(list,sortWithSubordinateMarkValue)
			}
		},
	}
)

func init() {
	searchCmd.AddCommand(hanyeSearchCmd)
}
