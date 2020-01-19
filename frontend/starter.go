package main

import (
	"firstCrawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./view")))

	searchResultHandler := controller.CreateSearchResultHandler("./view/template.html")
	http.Handle("/search", searchResultHandler)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
