package socks

import "time"

type auth struct {
	Username string
	Password string
}

type config struct {
	Proto   int
	Host    string
	Auth    *auth
	Timeout time.Duration
}
