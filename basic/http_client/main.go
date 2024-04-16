package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://www.kmc.gr.jp/"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
