package fetcher

import (
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimit = time.Tick(50 * time.Millisecond)

//Fetch all the data from website into byte slice
func Fetch(url string) ([]byte, error) {
	<-rateLimit
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (BB10; Touch) AppleWebKit/537.1+ (KHTML, like Gecko) Version/10.0.0.1337 Mobile Safari/537.1+")

	client := &http.Client{}
	resp, err := client.Do(req)
	// resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return all, nil
}
