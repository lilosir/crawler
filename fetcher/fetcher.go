package fetcher

import (
	"io/ioutil"
	"net/http"
)

//Fetch all the data from website into byte slice
func Fetch(url string) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1 (KHTML, like Gecko) CriOS/79.0.3945.88 Mobile/13B143 Safari/601.1.46")

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
