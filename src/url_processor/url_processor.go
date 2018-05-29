package url_processor

type Url string

type Urls []Url

type ResultJson struct {
	URL string `json:"url"`
	Meta struct {
		Title     string `json:"title"`
		Price     string `json:"price"`
		Image     string `json:"image"`
		Available bool   `json:"available"`
	} `json:"meta"`
}

type UrlsToProcess struct {
	Urls Urls
}

var urlsToProcess UrlsToProcess

func AddUrls(urls Urls) {

	urlsToProcess.Urls = append(urlsToProcess.Urls, urls...)

}