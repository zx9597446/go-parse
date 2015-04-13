package parse

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

const classBaseURL = "classes"

func (c *Client) makeClassURL(parts ...string) string {
	return c.makeURLPrefix(classBaseURL, parts...)
}

func (c *Client) createObject(className, jsonData string) (*result, error) {
	return c.httpPost(c.makeClassURL(className), jsonData)
}

func (c *Client) getObject(className, objectId, include string) (*result, error) {
	p := url.Values{}
	p.Add("include", include)
	url := c.makeClassURL(className, objectId)
	return c.httpGet(url, p)
}

func (c *Client) getObjectDirectly(location string) (*result, error) {
	return c.httpGet(location, nil)
}

func (c *Client) updateObject(className, objectId, jsonData string) (*result, error) {
	url := c.makeClassURL(className, objectId)
	return c.httpPut(url, jsonData)
}

func (c *Client) deleteObject(className, objectId string) (*result, error) {
	url := c.makeClassURL(className, objectId)
	return c.httpDelete(url)
}

func (c *Client) queryObjects(options QueryOptions) (*result, error) {
	if options.Class == "" {
		return nil, errors.New("require class name to query")
	}
	p := url.Values{}
	where, err := json.Marshal(options.Where)
	if err != nil {
		return nil, err
	}
	p.Add("where", string(where))
	if options.Limit <= 0 {
		options.Limit = DefaultLimit
	}
	p.Add("limit", strconv.Itoa(options.Limit))
	p.Add("skip", strconv.Itoa(options.Skip))
	if options.Order != "" {
		p.Add("order", options.Order)
	}
	if options.Keys != "" {
		p.Add("keys", options.Keys)
	}
	if options.Include != "" {
		p.Add("include", options.Include)
	}
	if options.Count {
		p.Add("count", "1")
	}
	url := c.makeClassURL(options.Class)
	return c.httpGet(url, p)
}
