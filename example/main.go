package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pchchv/sox"
)

func main() {
	// Create a SOCKS proxy dialing function.
	dialSocksProxy := sox.Dial("socks5://127.0.0.1:1080?timeout=5s") // user/password authentication
	tr := &http.Transport{Dial: dialSocksProxy}
	httpClient := &http.Client{Transport: tr}

	resp, err := httpClient.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.StatusCode)
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf))
}
