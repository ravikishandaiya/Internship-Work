package client

import (
	"fmt"
	"io"
	"net/http"
)


client := &http.Client{
	checkRedirect: re
}

func main() {
	resp, err := http.Get("https://api.zoom.us/v2/users/")
	defer resp.Body.Close()

	fmt.Print(err)

	body,  err := io.ReadAll(resp.Body)

	fmt.Print(err)

	if err != nil {
		fmt.Print("Error:\n")
		fmt.Print(err)
	} else {
		fmt.Print("Responce:\n")
		fmt.Print(body)
	}
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
