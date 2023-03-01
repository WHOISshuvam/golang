/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the ip",
	Long:  `tarce the ip`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Println(ip)
			}
		} else {
			fmt.Printf("Error")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

// func showData() {

// }

// func getdata() {
// 	url := "http://google.com"
// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Println("Unable to get the response")
// 	}

// 	responseByte, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Println("Unable to get the response")

// 	}
// }
