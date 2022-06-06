package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	reportCmd = &cobra.Command{
		Use:   "report",
		Short: "report every day care",
		Long:  "report every day care",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				log.Fatal("report every day care args example 5 2")
			}
			var (
				day, count int
			)
			day, _ = strconv.Atoi(args[0])
			count, _ = strconv.Atoi(args[1])
			fmt.Printf("%d 日内，%d 涨停记录",day,count)
			reportCareAboutStock(day, count) //
		},
	}
	hardenTodayCmd = &cobra.Command{
		Use:   "today",
		Short: "report every harden stock",
		Long:  "report every harden stock",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("今日涨停：")
			reportDailyLimitStatisticsStock()
		},
	}
)

func init() {
	reportCmd.AddCommand(hardenTodayCmd)
}
