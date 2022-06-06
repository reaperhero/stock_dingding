package server

import (
	"fmt"
	"github.com/reaperhero/stock_dingding/service/stock"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "search stock",
		Long:  "search stock",
	}
	hanyeSearchCmd = &cobra.Command{
		Use:   "sub",
		Short: "show stock in subordinate",
		Long:  "show stock in subordinate",
		Run: func(cmd *cobra.Command, args []string) {
			for _, arg := range args {
				list, err := stock.GetStockBySubordinate(arg)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(EchoStock(list, SortWithSubordinateSixDaysChange))
			}
		},
	}
)

func init() {
	searchCmd.AddCommand(hanyeSearchCmd)
}
