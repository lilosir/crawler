package fetcher

import (
	"io/ioutil"
	"net/http"
)

//Fetch all the data from website into byte slice
func Fetch(url string) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

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
