package domain

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 5,
}

// MakeHTTPCall is a wrapper function to aid performing an HTTP call
func MakeHTTPCall(requestMethod string, requestURL string, authorizationHeader string, rangeHeader string, skipVerifySSL bool) (*http.Response, error) {

	var response *http.Response

	// Construct the request
	request, err := http.NewRequest(requestMethod, requestURL, nil)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	VLog(fmt.Sprintf("[request] %s %s", request.Method, request.URL))

	// Set the Headers
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("User-Agent", "sba-cli/1.0")

	if authorizationHeader != "" {
		request.Header.Set("Authorization", authorizationHeader)
		VLog(fmt.Sprintf("[request] Authorization: %s", authorizationHeader))
	}

	if rangeHeader != "" {
		request.Header.Set("Range", rangeHeader)
		VLog(fmt.Sprintf("[request] Range: %s", rangeHeader))
	}

	// Handle skipping SSL verification
	if skipVerifySSL {
		VLog(">>> Skipping SSL Verification")
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient.Transport = tr
	}

	// Make the call
	return httpClient.Do(request)

}

// GenerateRequestURL is a utility function to cleanly generate a URL based on the passed baseURL and path
func GenerateRequestURL(baseURL string, pathToAppend string) (string, error) {

	var generatedRequestURL string

	u, err := url.Parse(baseURL)
	if err != nil {
		return generatedRequestURL, err
	}

	u.Path = path.Join(u.Path, pathToAppend)

	generatedRequestURL = u.String()

	return generatedRequestURL, nil

}

func ResponseBodyToStr(response *http.Response) (string, error) {

	var responseBodyStr string

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ELog(fmt.Sprintf("Error in ResponseBodyToStr error='%s'", err.Error()))
		return responseBodyStr, err
	}

	responseBodyStr = string(responseBody)

	return responseBodyStr, nil

}
