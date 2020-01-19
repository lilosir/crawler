package controller

import (
	"context"
	"firstCrawler/engine"
	"firstCrawler/frontend/model"
	"firstCrawler/frontend/view"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/olivere/elastic/v7"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

// get the reques and handle the query
// localhost:8888/search?q=男&from=10
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.FormValue("q"))
	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}
	// fmt.Fprintf(w, "q=%s, from=%d", a, from)

	// var page model.SearchResult
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q

	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(reWriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func reWriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
