package server

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
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
	rootCmd.AddCommand(createCmd, reportCmd, classIfication, searchCmd)
}
