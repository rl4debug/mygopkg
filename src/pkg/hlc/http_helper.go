/*
	Because website HLC only requires HTTP GET request, so we build some simple helper functions to make code easier and more clear for reading
	There is nothing exported in this file
*/

package hlc

import (
	"net/http"
	"net/url"
)

//Singleton HTTP client
var httpClient *http.Client
var hlcBaseLink string
var defaultHTTPHeaders map[string]string

//Should see documentation of the init function
func init() {

	httpClient = &http.Client{}

	//may be we should move below stuffs to a configuration section. But it changes rarely, so we can place it here
	hlcBaseLink = "https://www.hellochao.vn/tu-dien-tach-ghep-am/?lang=&type=sentence&act=search"
	defaultHTTPHeaders = map[string]string{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
		"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
	}
}

//This is simple internal HTTP client which supports GET only (for website HLC)
type simpleHTTPGetClient struct {
}

func (client *simpleHTTPGetClient) Get(baseLink string, queryParams map[string]string, headersSlices ...map[string]string) (*http.Response, error) {
	httpRequest, err := client.makeGetRequest(baseLink, queryParams, headersSlices...)
	if err != nil {
		return nil, err
	}

	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, err
}

func (client *simpleHTTPGetClient) makeGetRequest(baseLink string, queryParams map[string]string, headersSlices ...map[string]string) (*http.Request, error) {
	link, err := client.makeLink(baseLink, queryParams)
	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}

	for _, header := range headersSlices {
		for name, value := range header {
			httpRequest.Header.Add(name, value)
		}
	}

	return httpRequest, nil
}

func (client *simpleHTTPGetClient) makeLink(baseLink string, queryParams map[string]string) (string, error) {
	baseURL, err := url.ParseRequestURI(baseLink)
	if err != nil {
		return "", err
	}

	query := baseURL.Query()
	for name, value := range queryParams {
		query.Add(name, value)
	}
	baseURL.RawQuery = query.Encode()

	return baseURL.String(), nil
}

// func makeSimpleHTTPGetRequest(url, cookie string) (request *http.Request, err error) {
// 	request, err = http.NewRequest("GET", url, nil)

// 	for name, value := range defaultHTTPHeaders {
// 		request.Header.Add(name, value)
// 	}
// 	request.Header.Add("Cookie", cookie)

// 	//unadorned return (this is an idiom in Go language, we can understand it similar to 'naked return')
// 	return
// }

// func makeTranslatingHTTPRequest(textSearch, cookie string) (request *http.Request, err error) {
// 	link := makeTranslatingLink(textSearch)
// 	return makeSimpleHTTPGetRequest(link, cookie)
// }

// func makeTranslatingLink(textSearch string) string {
// 	//ignore error, because we use hardcode link
// 		hlcBaseURL, _ := url.ParseRequestURI(hlcBaseLink)
// 		query := hlcBaseURL.Query()
// 		query.Add("sct", textSearch)
// 		hlcBaseURL.RawQuery = query.Encode()
// 		return hlcBaseURL.String()
// }

// func getTranslatingResponse(r *http.Request) (*http.Response, error) {
// 	return httpClient.Do(r)
// }
