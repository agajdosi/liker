package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	for {
		url := os.Getenv("GET_URL")
		resp, _ := http.Get(url)

		fmt.Println("getting -", url, "respond:", resp)

		time.Sleep(time.Second * 5)
	}
}
