package parse

import "net/url"

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

func (c *Client) queryObject(className, whereJson, limit, skip, order, keys string) (*result, error) {
	p := url.Values{}
	if whereJson != "" {
		p.Add("where", whereJson)
	}
	if limit != "" {
		p.Add("limit", limit)
	}
	if skip != "" {
		p.Add("skip", skip)
	}
	if order != "" {
		p.Add("order", order)
	}
	if keys != "" {
		p.Add("keys", keys)
	}
	url := c.makeClassURL(className)
	return c.httpGet(url, p)
}
