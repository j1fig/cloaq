package cmd

import (
	"fmt"
	"github.com/j1fig/cloaq/oaq"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

// citiesCmd represents the cities command.
var citiesCmd = &cobra.Command{
	Use:   "cities",
	Short: "list all available cities",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := oaq.Config{ApiUrl}
		url := fmt.Sprintf("%v/%v", cfg.Url, "cities")
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			log.Println(bodyString)
		}
	},
	Args: cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(citiesCmd)
}
