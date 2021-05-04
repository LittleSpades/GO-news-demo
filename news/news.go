package news

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const httpTimeout = 10 * time.Second

type NewsFetcher interface {
	FetchNews(string, string) ([]byte, error)
}

// Client is ...
type Client struct {
	http     *http.Client
	BaseURL  string
	PageSize int
}

func New(baseURL string) NewsFetcher {
	return &Client{
		http:    &http.Client{Timeout: httpTimeout},
		BaseURL: baseURL,
	}
}

// FetchNews is ...
func (c *Client) FetchNews(endpoint string, parameters string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s?%s", c.BaseURL, endpoint, parameters)
	resp, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	return body, err
}
