package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://YOUR_DOMAIN/.well-known/jwks.json")
}


/*
// holds informations to connnect to api
type Client struct {
	address string
	token string
}

func NewClient(address string, token string) *client {
	return &Client{
		address: address,
		token: token,
		httpClient: &http.client{},
	}
}

func (c *client) GetAll() (*map[string]server.User, error) {

}
*/ 