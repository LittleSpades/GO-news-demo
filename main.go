package main //the code from main.go file belongs to the main package

import (
	"log"

	//provides http client and server implementations for use in the app
	"net/http"

	//access operating system functionality
	"os"

	"github.com/LittleSpades/GO-news-hub/news"
	"github.com/LittleSpades/GO-news-hub/transport"
)

func main() {

	//The Load method reads the .env file and loads the set variables into the environment so that they can be accessed through the os.Getenv() method
	//err := godotenv.Load()
	//if err != nil {
	//	log.Println("Error loading .env file")
	//}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		log.Fatal("Env: apiKey must be set")
	}

	newsAPIURL := news.New("https://newsapi.org/v2")
	hackerNewsAPIURL := news.New("http://hn.algolia.com/api/v1")

	fs := http.FileServer(http.Dir("assets"))

	//create an http request multiplexer
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//Registering http request handlers
	mux.HandleFunc("/search", transport.SearchNewsAPI(newsAPIURL, apiKey))
	mux.HandleFunc("/search_hn", transport.SearchHackerNewsAPI(hackerNewsAPIURL))

	mux.HandleFunc("/", transport.IndexHandler)
	http.ListenAndServe(":"+port, mux)
}
