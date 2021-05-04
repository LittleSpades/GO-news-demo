package news

import (
	"fmt"
	"time"
)

// Article is ...
type ArticleNewsAPI struct {
	Source struct {
		ID   interface{} `json:"id"`
		Name string      `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

//FormatPublishedDate is ...
func (a *ArticleNewsAPI) FormatPublishedDate() string {
	year, month, day := a.PublishedAt.Date()
	return fmt.Sprintf("%v %d, %d", month, day, year)
}

// Results is ...
type ResultsNewsAPI struct {
	Status       string           `json:"status"`
	TotalResults int              `json:"totalResults"`
	Articles     []ArticleNewsAPI `json:"articles"`
}

// Article is ...
type ArticleHackerNewsAPI struct {
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	PublishedAt time.Time `json:"created_at"`
}

// Results is ...
type ResultsHackerNewsAPI struct {
	TotalResults int                    `json:"nbHits"`
	Hits         []ArticleHackerNewsAPI `json:"hits"`
}
