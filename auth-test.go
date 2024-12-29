package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://localhost:8080/videos"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Basic cm9vdDpyb290")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode == 200 {
		fmt.Println(string(body))
	} else {
		fmt.Println(res.StatusCode)
		fmt.Println(res.Request)
		fmt.Println(string(body))
	}
}
