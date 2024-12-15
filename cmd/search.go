/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Chimimouryou/JsonsStrcuts"
	"Chimimouryou/ui/bubbleTeaUi"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"net/url"
	"os/exec"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		bubbleTeaUi.Main()

		//animeId, err := searchAnime(animeName)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//episodeId, err := getAnimeInfo(animeId)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//parseJsonData(episodeId)
		//
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

// First we need to get the anime info

func parseJsonData(episodeId string) {
	jsonBody, err := fetchVideoFile(episodeId)
	if err != nil {
		fmt.Println(err)
		return
	}

	var animeStreams JsonsStrcuts.AnimeStreams
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

func fetchJsonData(fullUrl string) ([]byte, error) {

	resp, err := http.Get(fullUrl)
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

func fetchVideoFile(epidsodeId string) ([]byte, error) {
	baseURL := fmt.Sprintf("http://localhost:3000/anime/gogoanime/watch/%s", epidsodeId)

	params := url.Values{}
	params.Add("server", "streamsb")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Println(fullURL)

	jsonData, err := fetchJsonData(fullURL)
	if err != nil {
		fmt.Errorf("Failed to fetch video file: %v", err)
		return nil, err
	}

	return jsonData, nil

}

func searchAnime(name string) (string, error) {
	var fullUrl string
	fullUrl = fmt.Sprintf("http://localhost:3000/anime/gogoanime/%s", name)

	resp, err := fetchJsonData(fullUrl)

	var animeSearchQuery JsonsStrcuts.AnimeSearchQuery
	err = json.Unmarshal(resp, &animeSearchQuery)
	if err != nil {
		fmt.Errorf("Failed to parse the response body: %v", err)
		return "", nil
	}

	idList := []string{}

	for _, source := range animeSearchQuery.Results {
		// TODO: Need to handle how im going to store these results. come time to create the UI...
		idList = append(idList, source.ID)
	}

	return idList[0], nil

}

func getAnimeInfo(animeID string) (string, error) {
	url := fmt.Sprintf("http://localhost:3000/anime/gogoanime/info/%s", animeID)

	resp, err := fetchJsonData(url)

	var animeInfo JsonsStrcuts.AnimeInfo
	err = json.Unmarshal(resp, &animeInfo)
	if err != nil {
		fmt.Errorf("Failed to parse the response body: %v", err)
		return "", nil
	}

	for _, source := range animeInfo.Episodes {
		return source.ID, nil
	}

	return "", nil

}
