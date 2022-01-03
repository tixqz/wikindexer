package main

import (
	"fmt"
	"io"
	"net/http"
)

const WIKI_API = "wikipedia.org/w/api.php"

type Client struct {
	url    string
	locale string
}

func DefaultClient() *Client {
	return &Client{
		url:    WIKI_API,
		locale: "en",
	}
}

func (c *Client) FetchAllPages(curr string) {
	url := fmt.Sprintf("https://%s.%s", c.locale, c.url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Errorf("Can't get response from %s. Error: %v", url, err)
	}
	err = c.SerializeData(res.Body)
}

func (c *Client) SerializeData(data io.Reader) error {

	return nil
}
