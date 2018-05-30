package url_processor

import (
	"crypto/md5"
	"net/http"
	"utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"unicode/utf8"
	"regexp"
	"strings"
	"encoding/hex"
)

type Urls []string

type ResultJson struct {
	URL string `json:"url"`
	Meta struct {
		Title     string  `json:"title"`
		Price     float64 `json:"price"`
		Currency  string  `json:"currency"`
		Image     string  `json:"image"`
		Available bool    `json:"available"`
	} `json:"meta"`
}

var urlsToProcess []UrlToProcess

type UrlToProcess struct {
	Md5 string
	Url string
}

var urlsChan chan UrlToProcess

var UrlsDone map[string]ResultJson

func init() {
	UrlsDone = make(map[string]ResultJson)
	urlsChan = make(chan UrlToProcess, 10)
}

func AddUrls(urls Urls) []UrlToProcess {

	for _, url := range urls {

		hasher := md5.New()
		hasher.Write([]byte(url))

		urlToProcess := UrlToProcess{}
		urlToProcess.Md5 = hex.EncodeToString(hasher.Sum(nil))
		urlToProcess.Url = url

		urlsToProcess = append(urlsToProcess, urlToProcess)

		urlsChan <- urlToProcess
	}

	return urlsToProcess
}

func Start() {

	fmt.Println("Start url processor")

	for {
		select {
		case u := <-urlsChan:
			UrlsDone[u.Md5] = ProcessUrl(u.Url)
		}
	}
}


func ProcessUrl(url string) ResultJson {

	fmt.Println("Processing url", url)

	var result ResultJson

	result.URL = url

	res, err := http.Get(url)
	utils.CheckError(err)

	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Printf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckError(err)

	// find tile
	doc.Find("#productTitle").Each(func(i int, s *goquery.Selection) {
		result.Meta.Title = s.Text()
	})

	// find price and currency
	doc.Find("span .offer-price").Each(func(i int, s *goquery.Selection) {

		curArr := []rune(s.Text())
		result.Meta.Currency = string(curArr[0])

		price, err := strconv.ParseFloat(trimFirstRune(s.Text()), 64)
		utils.CheckError(err)
		result.Meta.Price = price
	})

	// find image url
	doc.Find("#imgBlkFront").Each(func(i int, s *goquery.Selection) {
		img, exists := s.Attr("data-a-dynamic-image")
		if exists {
			result.Meta.Image = GetStringFromQuotes(img)
		}
	})

	// find availability
	doc.Find("#availability").Each(func(i int, s *goquery.Selection) {
		avail := s.Find("span").Text()
		result.Meta.Available = strings.TrimSpace(avail) == "In stock."
	})

	fmt.Printf("%+v\n", result)
	return result
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func GetStringFromQuotes(str string) string {
	re := regexp.MustCompile(`("[^"]+")`)

	founds := re.FindAllString(str, 1)
	if len(founds) > 0 {
		return re.FindAllString(str, 1)[0]
	} else {
		return ""
	}

}
