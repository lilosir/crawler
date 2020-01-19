package view

import (
	"firstCrawler/frontend/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	template := template.Must(template.ParseFiles(filename))
	return SearchResultView{template}
}

func (s SearchResultView) Render(wr io.Writer, data model.SearchResult) error {
	return s.template.Execute(wr, data)
}
