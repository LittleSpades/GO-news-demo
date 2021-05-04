package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/freshman-tech/news-demo-starter-files/news"
)

func SearchHackerNewsAPI(newsAPI news.NewsFetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := u.Query()
		searchQuery := params.Get("q")
		page := params.Get("page")
		if page == "" {
			page = "1"
		}

		endpoint := "search_by_date"
		parameters := fmt.Sprintf("query=%s&page=%s&hitsPerPage=%d&tags=story", url.QueryEscape(searchQuery), page, 20)

		body, err := newsAPI.FetchNews(endpoint, parameters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		results := &news.ResultsHackerNewsAPI{}

		err = json.Unmarshal(body, results)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		nextPage, err := strconv.Atoi(page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		search := &Search{
			Query:      searchQuery,
			NextPage:   nextPage,
			TotalPages: int(math.Ceil(float64(results.TotalResults / 20))),
			Results:    results,
		}

		if ok := !search.IsLastPage(); ok {
			search.NextPage++
		}

		buf := &bytes.Buffer{}

		tpl := template.Must(template.ParseFiles("index_hn.html"))

		err = tpl.Execute(buf, search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(w)

		//body, err = json.Marshal(search)
		//w.Header().Set("Content-Type", "application/json")
		//w.Write(body)
	}
}
