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
func MakeHTTPCall(requestMethod string, requestURL string, authorizationHeader string, skipVerifySSL bool) (string, error) {

	var responseBodyStr string

	// Construct the request
	request, err := http.NewRequest(requestMethod, requestURL, nil)
	if err != nil {
		fmt.Println(err)
		return responseBodyStr, err
	}

	// Set the Headers
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("User-Agent", "sba-cli/1.0")

	if authorizationHeader != "" {
		request.Header.Add("Authorization", authorizationHeader)
	}

	// Handle skipping SSL verification
	if skipVerifySSL {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient.Transport = tr
	}

	// Explicitly print out the outgoing HTTP call
	if CLIConfig.Verbose {
		fmt.Printf(">>> %s %s \n", request.Method, request.URL)
	}

	// Make the call
	response, err := httpClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return responseBodyStr, err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return responseBodyStr, err
	}

	responseBodyStr = string(responseBody)

	if CLIConfig.Verbose {
		fmt.Printf(">>> %s %s \n", response.Status, response.Proto)
	}

	return responseBodyStr, nil

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
