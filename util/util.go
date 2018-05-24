package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostJson(url string, jsonBytes []byte, headers map[string]string) {
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(jsonBytes),
	)

	if err != nil {
		fmt.Println("request create error: ", err)
		return
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("requeset error: ", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
}
