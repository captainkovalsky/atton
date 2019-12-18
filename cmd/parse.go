/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	tools "github.com/captainkovalsky/atton/common"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	InvestmentAccountsTitle = "Investment accounts"
	titleSelector           = "h3"
	totalCountSelector      = "h1"
	newTodaySelector        = "span"
	statBlockSelector       = ".stats_block"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := colly.NewCollector()

		// Set HTML callback
		// Won't be called if error occurs
		c.OnHTML(statBlockSelector, func(e *colly.HTMLElement) {
			title := e.ChildText(titleSelector)
			if title == InvestmentAccountsTitle {
				totalCount, err := tools.TryParseNumber(e.ChildText(totalCountSelector))

				if err != nil {
					log.Error(err)
				}

				newTodayCount, err := tools.TryParseNumber(e.ChildText(newTodaySelector))

				if err != nil {
					log.Error(err)
				}

				log.WithFields(log.Fields{
					"totalCount":    totalCount,
					"newTodayCount": newTodayCount,
				}).Info("register information")
			}
		})

		// Set error handler
		c.OnError(func(r *colly.Response, err error) {
			log.WithFields(log.Fields{
				"URL":   r.Request.URL,
				"Error": err.Error(),
			}).Panic("failed with processing response")
		})

		// Start scraping
		c.Visit("https://attonbank.com/en/servicestatus/")

	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
