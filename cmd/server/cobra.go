package server

import (
	"fmt"
	"github.com/reaperhero/stock_dingding/model"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"strings"
)

var rootCmd = &cobra.Command{}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	classIfication = &cobra.Command{
		Use:   "classification",
		Short: "A stock classification type",
		Long:  "A stock classification type",
		Run: func(cmd *cobra.Command, args []string) {
			reportChinaAllStock()
		},
	}
)

func init() {
	rootCmd.AddCommand(createCmd, reportCmd, classIfication)
}
func sortStock(source []model.Stock) {
	sort.Slice(source, func(i, j int) bool {
		if strings.Compare(source[i].Subordinate, source[j].Subordinate) > 0 {
			return true
		}
		if strings.Compare(source[i].Subordinate, source[j].Subordinate) < 0 {
			return false
		}
		if source[i].Pe >= source[j].Pe {
			return false
		}
		return true
	})
}
