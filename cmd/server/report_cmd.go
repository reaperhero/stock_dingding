package server

import (
	"github.com/spf13/cobra"
	"time"
)

var (
	day = ""
	reportCmd = &cobra.Command{
		Use:   "report",
		Short: "show every harden",
		Long:  "show every harden",
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{"--day", "test"}, cobra.ShellCompDirectiveDefault
		},
		Run: func(cmd *cobra.Command, args []string) {
			reportDailyLimitStatisticsStock(day)
		},
	}
)

func init()  {
	reportCmd.Flags().StringVar(&day, "day", "", time.Now().Format("2006-01-02"))

}