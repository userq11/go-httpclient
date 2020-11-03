package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/userq11/go-httpclient/gohttp"
)

var githubHttpClient = getGithubClient()

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer asdjkp12p3o4emp2")
	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
}

func getUrls() {
	headers := make(http.Header)

	response, err := githubHttpClient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
