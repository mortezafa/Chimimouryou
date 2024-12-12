/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Chimimouryou/JsonsStrcuts"
	"Chimimouryou/ui/bubbleTeaUi"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("searching... One moment please.")

		//if len(args) < 0 {
		//	panic("You didnt specify an Anime Name!")
		//}
		//
		//animeId, err := searchAnime(args)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//episodeId, err := getAnimeInfo(animeId)
		//if err != nil {
		//	fmt.Println(err)
		//}	
		//fmt.Println(episodeId)
		//parseJsonData(episodeId)
		
		bubbleTeaUi.Main()
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

func fetchVideoFile (epidsodeId string) ([]byte, error) {
	baseURL := fmt.Sprintf("http://localhost:3000/anime/gogoanime/watch/%s", epidsodeId)

	params := url.Values{}
	params.Add("server", "vidstreaming")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	
	jsonData, err := fetchJsonData(fullURL)
	if err != nil {
		fmt.Errorf("Failed to fetch video file: %v", err)
		return nil, err
	}
	
	return jsonData, nil

}

func searchAnime(name []string) (string, error) {
	var fullUrl string
	if len(name) == 1 {
		fullUrl = fmt.Sprintf("http://localhost:3000/anime/gogoanime/%s", name[0])
	} else {
		joinedUrl := strings.Join(name, "%20")
		fmt.Println(joinedUrl)
		fullUrl = fmt.Sprintf("http://localhost:3000/anime/gogoanime/%s", joinedUrl)
	}
	
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
	fmt.Println("HELLO???")
	fmt.Println(strings.Join(idList, "\n"))

	return idList[0], nil

}

func getAnimeInfo(animeID string) (string, error)  {
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