package cachedclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	YEAR  = time.Hour * 24 * 365
	MONTH = time.Hour * 24 * 28
	WEEK  = time.Hour * 24 * 7
	DAY   = time.Hour * 24
)

type Client struct {
	client           *http.Client
	Debug            bool
	IgnoreCache      bool
	IgnoreExpiration bool
	baseURL          string
	cache            RequestCache
	Auth             func(*http.Request)
}

func (c *Client) GetObjFromAPI(url string, val interface{}, expTime time.Duration) error {
	url = path.Join(string(c.baseURL), url)
	body, err := c.GetBody(url, true, expTime)
	if err != nil {
		return err
	}
	buff := &bytes.Buffer{}
	io.Copy(buff, body)
	cp := buff.String()
	err = json.NewDecoder(buff).Decode(val)
	if err != nil || c.Debug {
		logger.Println("trace: body: ", cp)
	}
	return err
}

func (c *Client) GetObjNoBase(url string, val interface{}, expTime time.Duration) error {
	body, err := c.GetBody(url, false, expTime)
	if err != nil {
		return err
	}
	buff := &bytes.Buffer{}
	io.Copy(buff, body)
	cp := buff.String()
	err = json.NewDecoder(buff).Decode(val)
	if err != nil && c.Debug {
		logger.Println("trace: body: ", cp)
	}
	return err
}

func (c *Client) GetBody(url string, auth bool, expTime time.Duration) (io.Reader, error) {
	if !c.IgnoreCache {
		body := c.cache.GetRequest(url, expTime)
		if body != "" {
			if c.Debug {
				logger.Println("used cache: ", url)
			}
			return strings.NewReader(body), nil
		}
	}
	resp, err := c.Get(url, auth)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buff := &bytes.Buffer{}
	io.Copy(buff, resp.Body)
	if !c.IgnoreCache {
		c.cache.StoreRequest(url, buff.String())
	}
	return buff, nil
}

func (c *Client) Get(url string, auth bool) (*http.Response, error) {
	if !strings.HasPrefix(url, "http") {
		url = fmt.Sprintf("https://%s", url)
	}
	logger.Printf("trace: GET: %s", url)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if auth {
		c.Auth(r)
	}
	resp, err := c.client.Do(r)
	if err != nil {
		if c.Debug {
			logger.Printf("err: failed request: %v", err)
		}
		return resp, err
	}
	switch resp.StatusCode {
	case http.StatusTooManyRequests:
		if c.Debug {
			logger.Println("debug: Headers on 429 request:", resp.Header)
		}
		if resp.Header.Get("Retry-After") != "" {
			wait := resp.Header.Get("Retry-After")
			waitn, _ := strconv.Atoi(wait)
			if waitn == 0 {
				waitn = 2
			}
			time.Sleep(time.Second * time.Duration(waitn))
		}
		if c.Debug {
			logger.Println("trace: slow down charlie.\r")
		}
		resp, err = c.Get(url, auth)
		return resp, err
	case http.StatusNotFound:
		logger.Println("err: not found", url)
		return nil, fmt.Errorf("err: object not found: %s", url)
	case http.StatusOK, http.StatusAccepted:
		return resp, err
	default:
		logger.Println("err: Code not expected:", resp.StatusCode)
		return nil, fmt.Errorf("err: %d %s", resp.StatusCode, url)
	}
}

//SwapCache change the requestcache used by the Client.
func (c *Client) SwapCache(cache RequestCache) error {
	if cache == nil {
		return fmt.Errorf("must not provide nil cache")
	}
	c.cache = cache
	return nil
}

//NewClient make a new cached http client for interfacing with riots api.
func NewClient(baseUrl string, cache RequestCache, auth func(r *http.Request)) *Client {
	c := cache
	if c == nil {
		c = &memCache{
			requests: make(map[string]request),
		}
	}
	return &Client{
		Auth:    auth,
		cache:   c,
		baseURL: baseUrl,
		client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}
