# GO News Hub

In order to learn how to use GO for Web Applications, I started by following a tutorial (https://freshman.tech/web-development-with-go/) for a simple app that fetches news articles matching a particular search query through the News API (https://newsapi.org/), and presents the results on the page.

After completing the tutorial, I rearranged the file structure and implemented interfaces in order to be able to use mutiple API's. The second API that I used for this exercise was the "Hacker News API" (https://hn.algolia.com/api).

A deployed version is available through Heroku at http://go-news-hub.herokuapp.com/.

Here are some pictures:

![](images/news-api.png)

![](images/hacker-news-api.png)

## Prerequisites

You need to have [Go](https://golang.org/dl/) installed on your computer. The
version I used is **1.16.3**.

## Running locally

You'll need to create a .env file with your NEWS_API_KEY.
