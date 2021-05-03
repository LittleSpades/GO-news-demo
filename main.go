package main //the code from main.go file belongs to the main package

import (
	"bytes"
	"html/template"
	"math"

	"log"

	//provides http client and server implementations for use in the app
	"net/http"

	"net/url"

	//access operating system functionality
	"os"

	"strconv"

	"time"

	"github.com/freshman-tech/news-demo-starter-files/news"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

//Search is ...
type Search struct {
	Query      string
	NextPage   int
	TotalPages int
	Results    *news.Results
}

//IsLastPage is ...
func (s *Search) IsLastPage() bool {
	return s.NextPage >= s.TotalPages
}

//CurrentPage is ...
func (s *Search) CurrentPage() int {
	if s.NextPage == 1 {
		return s.NextPage
	}

	return s.NextPage - 1
}

//PreviousPage is ...
func (s *Search) PreviousPage() int {
	return s.CurrentPage() - 1
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

func searchHandler(newsapi *news.Client) http.HandlerFunc {
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

		results, err := newsapi.FetchEverything(searchQuery, page)
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
			TotalPages: int(math.Ceil(float64(results.TotalResults / newsapi.PageSize))),
			Results:    results,
		}

		if ok := !search.IsLastPage(); ok {
			search.NextPage++
		}

		buf := &bytes.Buffer{}
		err = tpl.Execute(buf, search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(w)
	}
}

func main() {

	//The Load method reads the .env file and loads the set variables into the environment so that they can be accessed through the os.Getenv() method
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	apikey := os.Getenv("NEWS_API_KEY")
	if apikey == "" {
		log.Fatal("Env: apikey must be set")
	}

	myClient := &http.Client{Timeout: 10 * time.Second}
	newsapi := news.NewClient(myClient, apikey, 20)

	fs := http.FileServer(http.Dir("assets"))

	//create an http request multiplexer
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	//Registering http request handlers
	mux.HandleFunc("/search", searchHandler(newsapi))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
