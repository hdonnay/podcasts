package pocketcasts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
)

var root *url.URL

func init() {
	var err error
	root, err = url.Parse("https://play.pocketcasts.com/")
	if err != nil {
		panic(err)
	}
}

const statusErr = "pocketcasts: bad HTTP response code: %v"

type Client struct {
	c *http.Client
}

func New(user, pass string) (*Client, error) {
	j, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return nil, err
	}
	c := &http.Client{Jar: j}
	u, err := root.Parse("/users/sign_in")
	if err != nil {
		return nil, err
	}
	u.RawQuery = url.Values{
		"user[email]":       {user},
		"user[password]":    {pass},
		"user[remember_me]": {"1"},
	}.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(statusErr, res.Status)
	}

	return &Client{c: c}, nil
}

// {}
func (c *Client) All() (*All, error) {
	u, err := root.Parse("/web/podcasts/all.json")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(`{}`))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")
	res, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		io.Copy(os.Stderr, res.Body)
		return nil, fmt.Errorf(statusErr, res.Status)
	}
	defer res.Body.Close()
	v := &All{}
	if err := json.NewDecoder(res.Body).Decode(v); err != nil {
		return nil, err
	}
	return v, nil
}
