package sox

import (
	"bytes"
	"errors"
	"net"
	"strconv"
	"time"
)

type requestBuilder struct {
	bytes.Buffer
}

func (b *requestBuilder) add(data ...byte) {
	_, _ = b.Write(data)
}

func (cfg *config) readAll(conn net.Conn) (resp []byte, err error) {
	resp = make([]byte, 1024)
	if cfg.Timeout > 0 {
		if err := conn.SetReadDeadline(time.Now().Add(cfg.Timeout)); err != nil {
			return nil, err
		}
	}

	n, err := conn.Read(resp)
	resp = resp[:n]
	return
}

func (cfg *config) sendReceive(conn net.Conn, req []byte) (resp []byte, err error) {
	if cfg.Timeout > 0 {
		if err := conn.SetWriteDeadline(time.Now().Add(cfg.Timeout)); err != nil {
			return nil, err
		}
	}

	if _, err = conn.Write(req); err != nil {
		return
	}
	return cfg.readAll(conn)
}

func splitHostPort(addr string) (host string, port uint16, err error) {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return "", 0, err
	}

	portInt, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return "", 0, err
	}

	port = uint16(portInt)
	return
}

func lookupIPv4(host string) (net.IP, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}

	for _, ip := range ips {
		ipv4 := ip.To4()
		if ipv4 == nil {
			continue
		}
		return ipv4, nil
	}
	return nil, errors.New("no IPv4 address found for host: " + host)
}
