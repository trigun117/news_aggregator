package fetch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

var link = os.Getenv("LINK")

// Sources struct contains information about article source
type Sources struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Article struct contains information about news
type Article struct {
	Author      string  `json:"author"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	URL         string  `json:"url"`
	URLToImage  string  `json:"urlToImage"`
	PublishedAt string  `json:"publishedAt"`
	Source      Sources `json:"source"`
}

// News struct contains articles
type News struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

// NewsArticles contains data about news
var NewsArticles News

// FreshNews fetching news from api
func FreshNews() error {
	resonse, _ := http.Get(link)
	body, _ := ioutil.ReadAll(resonse.Body)
	defer resonse.Body.Close()
	json.Unmarshal(body, &NewsArticles)
	return nil
}
