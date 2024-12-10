/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"

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
		fmt.Println("searching... One moment please.")
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

type AnimeStreams struct {
	Sources []struct {
		Url     string `json:"url"`
		IsM3U8  bool   `json:"isM3U8"`
		Quality string `json:"quality"`
	} `json:"sources"`
}

func parseJsonData() {
	jsonBody, err := fetchJsonBody()
	if err != nil {
		fmt.Println(err)
		return
	}

	var animeStreams AnimeStreams
	err = json.Unmarshal(jsonBody, &animeStreams)
	if err != nil {
		fmt.Println(err)
		return
	}

	urls := []string{}

	for _, source := range animeStreams.Sources {
		urls = append(urls, source.Url)
	}

	fmt.Println(urls[3])

	playVid := exec.Command("vlc", urls[3])

	err = playVid.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = playVid.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("COMMAND DONE YAY")

}

func fetchJsonBody() ([]byte, error) {
	baseURL := "http://localhost:3000/anime/gogoanime/watch/tokyo-ghoul-episode-1"

	params := url.Values{}
	params.Add("server", "vidstreaming")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Errorf("Failed to make the request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("Request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Failed to read the response body: %v", err)
	}
	return body, nil
}
