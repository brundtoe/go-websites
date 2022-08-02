package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	resp, err := client.Get(
		"http://leonora/api/customers/7")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
