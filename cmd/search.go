/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
		parseJsonData()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

// First we need to get the anime info

type AnimeInfo struct {
	ID            string `json:"id"`
	title         string `json:"title"`
	url           string `json:"url"`
	image         string `json:"image"`
	totalEpisodes int32  `json:"totalEpisodes"`
}

type AnimeEpisodesContent struct {
	ID     string `json:"id"`
	number int32  `json:"number"`
	url    string `json:"url"`
}

func parseJsonData() {
	baseURL := "http://localhost:3000/anime/gogoanime/watch/tokyo-ghoul-episode-1"

	params := url.Values{}
	params.Add("server", "vidstreaming")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		log.Fatalf("Failed to make the request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}
	fmt.Println(string(body))
}
