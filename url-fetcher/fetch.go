package url_fetcher

import "net/http"

func Fetch(url string) (*http.Response, error){
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}
