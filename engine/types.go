package engine

//Request contains a URL and ParseFunc
type Request struct {
	URL       string
	ParseFunc func([]byte) ParseResult
}

//ParseResult contains next requests and results in any type
type ParseResult struct {
	Requests []Request
	Items   []interface{}
}

//NilParser return empty ParseResult, do nothing
func NilParser([]byte) ParseResult {
	return ParseResult{}
}