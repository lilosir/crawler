package model

type SearchResult struct {
	Hits  int64
	Start int
	// Items []engine.Item
	Items    []interface{}
	Query    string
	PrevFrom int
	NextFrom int
}
