package domain

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/spf13/viper"
)

var httpClient = &http.Client{
	Timeout: time.Second * 5,
}

func MakeHTTPCall(requestMethod string, requestURL string, authorizationHeader string) (string, error) {

	var responseBodyStr string

	// Construct the request
	request, err := http.NewRequest(requestMethod, requestURL, nil)
	if err != nil {
		fmt.Println(err)
		return responseBodyStr, err
	}

	// Add the Headers
	request.Header.Add("Content-Type", "application/json; charset=utf-8")

	if authorizationHeader != "" {
		request.Header.Add("Authorization", authorizationHeader)
	}

	if viper.GetBool("verbose") {
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

	return responseBodyStr, nil

}

func GenerateRequestURL(requestBase string, pathToAppend string) (string, error) {

	var generatedRequestURL string

	u, err := url.Parse(requestBase)
	if err != nil {
		return generatedRequestURL, err
	}

	u.Path = path.Join(u.Path, pathToAppend)

	generatedRequestURL = u.String()

	return generatedRequestURL, nil

}
