package utils

import (
	"fmt"
	"io"
	"net/http"
)

func FetchJsonData(fullUrl string) ([]byte, error) {

	resp, err := http.Get(fullUrl)
	if err != nil {
		fmt.Errorf("Failed to make the request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("Request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Failed to read the response body: %v", err)
	}
	return body, nil
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}		
}
