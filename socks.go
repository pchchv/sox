package sox

import (
	"fmt"
	"net"
)

// Constants to choose which version of SOCKS protocol to use.
const (
	SOCKS4 = iota
	SOCKS4A
	SOCKS5
)

func (cfg *config) dialFunc() func(string, string) (net.Conn, error) {
	switch cfg.Proto {
	case SOCKS5:
		return func(_, targetAddr string) (conn net.Conn, err error) {
			return cfg.dialSocks5(targetAddr)
		}
	case SOCKS4, SOCKS4A:
		return func(_, targetAddr string) (conn net.Conn, err error) {
			return cfg.dialSocks4(targetAddr)
		}
	}
	return dialError(fmt.Errorf("unknown SOCKS protocol %v", cfg.Proto))
}

func dialError(err error) func(string, string) (net.Conn, error) {
	return func(_, _ string) (net.Conn, error) {
		return nil, err
	}
}
