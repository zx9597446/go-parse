package parse

func (c *Client) callfunc(fn, jsonParam string) (*result, error) {
	requestURL := c.makeURLPrefix("functions", fn)
	r, err := c.httpPost(requestURL, jsonParam)
	return r, err
}

func CallFunction(c *Client, fn, jsonParam string) (*Object, error) {
	r, err := c.callfunc(fn, jsonParam)
	if err != nil {
		return nil, err
	}
	return r.Decode()
}
