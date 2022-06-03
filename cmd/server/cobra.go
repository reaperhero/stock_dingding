package server

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
	rootCmd.AddCommand(createCmd, reportCmd, classIfication,searchCmd)
}


