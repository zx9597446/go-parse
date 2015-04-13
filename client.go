package parse

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const apiURL = "https://api.parse.com/1"

type Client struct {
	AppId         string
	AppKey        string
	MasterKey     string
	UsingMaster   bool
	DebugRequest  bool
	DebugResponse bool
}

func NewClient() *Client {
	return &Client{
		DebugRequest:  false,
		DebugResponse: false,
	}
}

func (c *Client) request(url, method, body string) (*result, error) {
	if c.DebugRequest {
		log.Println(method, url)
	}
	r, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-Parse-Application-Id", c.AppId)
	if c.UsingMaster {
		r.Header.Set("X-Parse-Master-Key", c.MasterKey)
	} else {
		r.Header.Set("X-Parse-REST-API-Key", c.AppKey)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	sbody, _ := ioutil.ReadAll(res.Body)
	if c.DebugResponse {
		log.Println(sbody)
	}
	var location string
	if u, err := res.Location(); err == nil {
		location = u.String()
	}
	ret := &result{res.StatusCode, location, string(sbody)}
	if !ret.CheckStatusCode() {
		return ret, errors.New(ret.Response)
	}
	return ret, nil
}

func (c *Client) httpGet(url string, param url.Values) (*result, error) {
	withQuery := fmt.Sprintf("%s?%s", url, param.Encode())
	return c.request(withQuery, "GET", "")
}

func (c *Client) httpPut(url, body string) (*result, error) {
	return c.request(url, "PUT", body)
}

func (c *Client) httpDelete(url string) (*result, error) {
	return c.request(url, "DELETE", "")
}

func (c *Client) httpPost(url, body string) (*result, error) {
	return c.request(url, "POST", body)
}

func (c *Client) makeURL(parts ...string) (url string) {
	var path string
	if len(parts) == 0 {
		log.Panicln("can not make url", parts)
	} else if len(parts) == 1 {
		path = parts[0]
	} else {
		path = strings.Join(parts, "/")
	}
	url = fmt.Sprintf("%s/%s", apiURL, path)
	return
}

func (c *Client) makeURLPrefix(prefix string, parts ...string) string {
	tmp := make([]string, 0)
	tmp = append(tmp, prefix)
	tmp = append(tmp, parts...)
	return c.makeURL(tmp...)
}
