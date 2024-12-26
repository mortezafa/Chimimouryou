/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Chimimouryou/ui/bubbleTeaUi"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "`chimi search` is used to search up any anime",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		bubbleTeaUi.Main()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

