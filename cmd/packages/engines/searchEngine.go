package engines

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type ISearchEgine interface {
	GetProvider() Client
}

type SearchEngine struct {
	Client Client
	Adress string
}

func (s *SearchEngine) GoGet(q string, c chan string) {
	res, err := http.Get(s.Adress + q)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		panic("Status code != 200" + res.Status)
	}
	go docParser(res, c)
}

func docParser(res *http.Response, c chan<- string) {
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	doc.Find(".mw-parser-output").Each(func(i int, s *goquery.Selection) {
		text := s.Find("p").Text()
		c <- text
	})
	close(c)
}

func (s *SearchEngine) GetProvider() Client {
	return s.Client
}

func (s *SearchEngine) PrintProvider() string {
	switch s.Client {
	case 0:
		return "Wikipedia"
	case 1:
		return "Google"
	case 2:
		return "Youtube"
	case 3:
		return "ExternalAPI"
	}
	return ""
}

type BaseURLs map[Client]string
