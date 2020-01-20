package client

import (
	"github.com/fobilow/import-cycle/problem/user"
	"net/http"
)

type Client struct {
	baseUrl string
	httpClient *http.Client


	User *user.UserService
}

func (c *Client) Execute(uri string, method string, payload []byte) {
	//make request
}

func NewClient() *Client {
	c := &Client{}
	c.User = &user.UserService{c}
	return c
}
