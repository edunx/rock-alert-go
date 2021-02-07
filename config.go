package alert

import (
	resty "github.com/go-resty/resty/v2"
)

type Config struct {
	Url      string
	Origin   string
	Notifier  string
	Resolver string
}

type Alert struct {
	C  Config

	Client *resty.Client
}
