package server

import (
	"github.com/spf13/cobra"
)

var (
	reportCmd = &cobra.Command{
		Use:   "report",
		Short: "show every harden",
		Long:  "show every harden",
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{"--day", "test"}, cobra.ShellCompDirectiveDefault
		},
		Run: func(cmd *cobra.Command, args []string) {
			reportDailyLimitStatisticsStock()
		},
	}
	hardenTodayCmd = &cobra.Command{
		Use:   "h_today",
		Short: "show every harden",
		Long:  "show every harden",
		Run: func(cmd *cobra.Command, args []string) {
			reportDailyLimitStatisticsStock()
		},
	}
)

func init() {
	reportCmd.AddCommand(hardenTodayCmd)
}
