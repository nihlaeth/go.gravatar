package gravatar

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Client is a container with gravatar
//settings
type Client struct {
	rating, defImg   string
	forceDef, secure bool
	size             int16
}

//NewClient initializes the client
//with default values
func NewClient() *Client {
	//initialize stuff
	client := &Client{}
	client.rating = "g"
	client.defImg = "404"
	client.forceDef = false
	client.secure = true
	client.size = 100
	return client
}

//SetRating sets the lowest rating allowed. 
//Options are g, pg, r and x.
func (c *Client) SetRating(rating string) {
	c.rating = rating
}

//SetDefaultImage sets the policy for default images. 
//Options are 404 (error), [url],
//mm, identicon, monsterid, wavatar, retro and blank.
func (c *Client) SetDefaultImage(defImg string) {
	c.defImg = defImg
}

//SetForceDefault: if this is set to true, the default image will always be returned.
func (c *Client) SetForceDefault(forceDef bool) {
	c.forceDef = forceDef
}

//SetSecure determins if ssl should be used.
func (c *Client) SetSecure(secure bool) {
	c.secure = secure
}

//SetSize sets the size of the returned avatar. Must be an integer between 1 and 2048.
func (c *Client) SetSize(size int16) {
	c.size = size
}

//URL returns the url of the gravatar.
func (c *Client) URL(email string) string {
	//create hash
	hash := md5.New()
	//hash.Write([]byte(strings.ToLower(strings.TrimSpace(email))))
	io.WriteString(hash, strings.ToLower(strings.TrimSpace(email)))
	md5str := string(hash.Sum([]byte{}))
	//build url
	gurl := "http"
	if c.secure {
		gurl += fmt.Sprintf("s://secure.gravatar.com/avatar/%x?", md5str)
	} else {
		gurl += fmt.Sprintf("://gravatar.com/avatar/%x?", md5str)
	}
	gurl += fmt.Sprintf("s=%d&d=%s&r=%s", c.size, url.QueryEscape(c.defImg), c.rating)
	if c.forceDef {
		gurl += fmt.Sprintf("&f=%t", c.forceDef)
	}
	return gurl
}

//Image fetches the raw image data of the gravatar.
func (c *Client) Image(email string) []byte {
	//get url
	gurl := c.URL(email)
	//fetch avatar
	req, err := http.Get(gurl)
	if err != nil {
		//handle error
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		//handle error
	}
	//now return the body data
	return body
}
