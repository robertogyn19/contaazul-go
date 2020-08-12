package contaazul

import (
	"fmt"
	"net/http"
	"time"
)

const (
	defaultBaseURL = "https://api.contaazul.com"
)

var (
	timeoutDefault = 30 * time.Second

	InvalidParamsError = fmt.Errorf("invalid params to create client")
)

type Client struct {
	baseURL      string
	accessToken  string
	refreshToken string
	clientID     string
	clientSecret string
	client       *http.Client
}

type Options struct {
	BaseURL      string
	AccessToken  string
	RefreshToken string
	ClientID     string
	ClientSecret string
	Timeout      time.Duration
}

func NewClient(opts Options) (*Client, error) {
	cli := new(Client)

	cli.baseURL = defaultBaseURL
	cli.accessToken = opts.AccessToken
	cli.refreshToken = opts.RefreshToken
	cli.clientID = opts.ClientID
	cli.clientSecret = opts.ClientSecret

	if cli.clientID == "" || cli.clientSecret == "" || cli.accessToken == "" || cli.refreshToken == "" {
		return cli, InvalidParamsError
	}

	if opts.BaseURL != "" {
		cli.baseURL = opts.BaseURL
	}

	if opts.Timeout.Nanoseconds() == 0 {
		opts.Timeout = timeoutDefault
	}

	cli.client = &http.Client{Timeout: opts.Timeout}
	return cli, nil
}
