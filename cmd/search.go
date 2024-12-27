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
	Short: "Used to search up any anime",
	Long:  `
	Chimimouryou is quite simple as of now. The only key you need to know is how to quit Chimimouryou: simply press Ctrl + C at any time.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		bubbleTeaUi.Main()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

