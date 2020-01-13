package engine

//Request contains a URL and ParseFunc
type Request struct {
	URL       string
	ParseFunc ParseFunc
}

// ParseFunc ..
type ParseFunc func(contents []byte, url string) ParseResult

//ParseResult contains next requests and results in any type
type ParseResult struct {
	Requests []Request
	Items    []Item
}

//Item can have varying info, but at least id and url
type Item struct {
	ID      string
	URL     string
	TYPE    string
	Payload interface{}
}

//NilParser return empty ParseResult, do nothing
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
