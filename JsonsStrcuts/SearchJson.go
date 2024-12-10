package JsonsStrcuts

type AnimeSearchQuery struct {
	CurrentPage int  `json:"current_page"`
	HasNextPage bool `json:"hasNextPage"`
	Results     []struct {
		ID string `json:"id"`
		Title string `json:"title"`
		URL   string `json:"url"`
		Image string `json:"image"`
		ReleaseDate string `json:"releaseDate"`
		SubOrDub string `json:"subOrDub"`
	}
}

type AnimeInfo struct{
	Episodes []struct {
		ID string `json:"id"`
		Number int `json:"number"`
		Url string `json:"url"`
	}
}

type AnimeStreams struct {
	Sources []struct {
		Url     string `json:"url"`
		IsM3U8  bool   `json:"isM3U8"`
		Quality string `json:"quality"`
	} `json:"sources"`
}