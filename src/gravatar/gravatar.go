package gravatar

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	rating, defImg   string
	forceDef, secure bool
	size             int16
}

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

//This sets the lowest rating allowed. Options are g, pg, r and x.
func (c *Client) SetRating(rating string) {
	c.rating = rating
}

//Sets the policy for default images. Options are 404 (error), [url],
//mm, identicon, monsterid, wavatar, retro and blank.
func (c *Client) SetDefaultImage(defImg string) {
	c.defImg = defImg
}

//If this is set to true, the default image will always be returned.
func (c *Client) SetForeDefault(forceDef bool) {
	c.forceDef = forceDef
}

//Use ssl?
func (c *Client) SetSecure(secure bool) {
	c.secure = secure
}

//Size of the returned avatar. Must be an integer between 1 and 2048.
func (c *Client) SetSize(size int16) {
	c.size = size
}

//This function returns the url of the gravatar
func (c *Client) Url(email string) string {
	//create hash
	hash := md5.New()
	hash.Write([]byte(strings.ToLower(strings.TrimSpace(email))))
	md5str := string(hash.Sum([]byte{}))
	//build url
	gurl := "http"
	if c.secure {
		gurl += fmt.Sprintf("s://secure.gravatar.com/avatar/%x?", md5str)
	} else {
		gurl += fmt.Sprintf("://gravatar.com/avatar/%x?", md5str)
	}
	gurl += fmt.Sprintf("s=%d&d=%s&f=%t&r=%s", c.size, url.QueryEscape(c.defImg), c.forceDef, c.rating)
	return gurl
}

func (c *Client) Img(email string) []byte {
	//get url
	gurl := c.Url(email)
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
