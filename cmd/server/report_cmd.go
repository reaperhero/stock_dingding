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
			fmt.Printf("%d 日内，%d 涨停记录", day, count)
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
	fileCmd = &cobra.Command{
		Use:   "tofile",
		Short: "report to docs folder",
		Long:  "report to docs folder",
		Run: func(cmd *cobra.Command, args []string) {
			reportCareAboutStockTofile()
		},
	}
	stionCmd = &cobra.Command{
		Use:   "stion",
		Short: "stion to situation folder",
		Long:  "stion to situation folder",
		Run: func(cmd *cobra.Command, args []string) {
			trendStock()
		},
	}
	stionInitCmd = &cobra.Command{
		Use:   "init",
		Short: "stion init to situation folder",
		Long:  "stion init to situation folder",
		Run: func(cmd *cobra.Command, args []string) {
			initTrendStock()
		},
	}
)

func init() {
	reportCmd.AddCommand(hardenTodayCmd, fileCmd, stionCmd)
	stionCmd.AddCommand(stionInitCmd)
}
