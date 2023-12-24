# sox # work [![Go Reference](https://pkg.go.dev/badge/github.com/pchchv/sox.svg)](https://pkg.go.dev/github.com/pchchv/sox)
A Go SOCKS (SOCKS4, SOCKS4A and SOCKS5) proxy package

The package provides `socks.Dial`, which returns a TCP dial function from a socks proxy connection string.
The returned dial function can be used to establish a TCP connection via the socks proxy or used to initialize `http.Transport` for an HTTP connection.