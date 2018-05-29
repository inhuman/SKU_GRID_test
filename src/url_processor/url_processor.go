package url_processor

import (
	"crypto/md5"
	"net/http"
	"utils"
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"unicode/utf8"
	"regexp"
	"strings"
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

type UrlsToProcess struct {
	Urls map[[16]byte]string
}

type UrlsDone struct {
	Urls map[[16]byte]ResultJson
}

var urlsToProcess UrlsToProcess

func init() {
	urlsToProcess.Urls = make(map[[16]byte]string)
}

func AddUrls(urls Urls) {

	for _, url := range urls {
		urlsToProcess.Urls[md5.Sum([]byte(url))] = url
		ProcessUrl(url)
	}
}

func ProcessUrl(url string) {

	var result ResultJson

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
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
