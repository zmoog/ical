/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	ics "github.com/arran4/golang-ical"
	"github.com/spf13/cobra"
	"github.com/zmoog/ws/feedback"
)

var url string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List events from a ICS calendar",
	Long:  `List events from a ICS calendar.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if url == "" {
			return fmt.Errorf("url is required")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		cal, err := ics.ParseCalendar(resp.Body)
		if err != nil {
			return err
		}

		feedback.PrintResult(Result{Calendar: cal})

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().StringVarP(&url, "url", "u", "", "URL to a ICS calendar")
}
