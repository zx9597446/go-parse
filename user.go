package parse

import "net/url"

const userBaseURL = "users"
const userClass = "_User"

type User struct {
	Object
}

func NewUser() *User {
	o := NewObject()
	return &User{*o}
}

func (u *User) Register(c *Client, username, password, email, phone string) (*result, error) {
	u.Set("username", username)
	u.Set("password", password)
	u.Set("email", email)
	u.Set("mobilePhoneNumber", phone)
	url := c.makeURLPrefix(userBaseURL)
	return c.httpPost(url, u.Encode())
}

func (u *User) Login(c *Client, username, password string) (*result, error) {
	p := url.Values{}
	p.Add("username", username)
	p.Add("password", password)
	uri := c.makeURLPrefix("login")
	r, err := c.httpGet(uri, p)
	if err != nil {
		return r, err
	}
	o, err := r.Decode()
	if err == nil {
		u.Data = o.Data
	}
	return r, err
}

func (u *User) fetchFrom(c *Client, objectId string) (*result, error) {
	url := c.makeURLPrefix(userBaseURL, objectId)
	r, err := c.httpGet(url, nil)
	if err != nil {
		return r, err
	}
	u1, err := r.Decode()
	if err == nil {
		u.Data = u1.Data
	}
	return r, err
}

func FetchUser(c *Client, userId string) (*User, error) {
	u := NewUser()
	_, err := u.fetchFrom(c, userId)
	return u, err
}
