# sox [![Go Reference](https://pkg.go.dev/badge/github.com/pchchv/sox.svg)](https://pkg.go.dev/github.com/pchchv/sox)
A Go SOCKS (SOCKS4, SOCKS4A and SOCKS5) proxy package

The package provides `socks.Dial`, which returns a TCP dial function from a socks proxy connection string.
The returned dial function can be used to establish a TCP connection via the socks proxy or used to initialize `http.Transport` for an HTTP connection.


## Usage

```go
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
```