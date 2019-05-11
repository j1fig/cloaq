package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:     "cloaq",
	Version: "0.1.0",
}

var Verbose bool
var ApiUrl string

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output mode")
	rootCmd.PersistentFlags().StringVarP(&ApiUrl, "api-url", "a", "https://api.openaq.org/v1", "OpenAQ API url")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
