package lol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const (
	BR   RegionEndPoint = "br1.api.riotgames.com"
	EUNE RegionEndPoint = "eun1.api.riotgames.com"
	EUW  RegionEndPoint = "euw1.api.riotgames.com"
	JP   RegionEndPoint = "jp1.api.riotgames.com"
	KR   RegionEndPoint = "kr.api.riotgames.com"
	LAN  RegionEndPoint = "la1.api.riotgames.com"
	LAS  RegionEndPoint = "la2.api.riotgames.com"
	NA   RegionEndPoint = "na1.api.riotgames.com"
	OCE  RegionEndPoint = "oc1.api.riotgames.com"
	TR   RegionEndPoint = "tr1.api.riotgames.com"
	RU   RegionEndPoint = "ru.api.riotgames.com"
	PBE  RegionEndPoint = "pbe1.api.riotgames.com"
)
const (
	WEEK = time.Hour * 24 * 7
	DAY  = time.Hour * 24
)

type RegionEndPoint string

type client struct {
	client            *http.Client
	baseURL           RegionEndPoint
	cache             *lolMongo
	requestsMade      *int64
	requestsSucceeded *int64
}

// Debug weather or not debug is enabled for the riot package.
var Debug bool

// NewClient returns a new client for performing operations using riots api.
func NewClient(region RegionEndPoint) (RiotClient, error) {
	// cache, err := NewLolMongo("dev.jhrb.us", 27217)
	cache, err := NewLolMongo("", 0)
	// cache, err := NewLolMongo("192.168.1.170", 27027)
	if err != nil {
		return nil, err
	}
	var x, y int64
	return &client{
		client: &http.Client{
			Timeout: time.Second * 5,
		},
		requestsMade:      &x,
		requestsSucceeded: &y,
		cache:             cache,
		baseURL:           region,
	}, nil
}

func (c *client) GetCache() *lolMongo {
	return c.cache
}

func (c *client) GetObjRiot(url string, val interface{}, expTime time.Duration) error {
	url = path.Join(string(c.baseURL), url)
	body, err := c.GetBody(url, true, expTime)
	if err != nil {
		return err
	}
	buff := &bytes.Buffer{}
	io.Copy(buff, body)
	cp := buff.String()
	err = json.NewDecoder(buff).Decode(val)
	if err != nil {
		logger.Println("trace: body: ", cp)
	}
	return err
}

func (c *client) GetObjUnauthedRiot(url string, val interface{}, expTime time.Duration) error {
	body, err := c.GetBody(url, false, expTime)
	if err != nil {
		return err
	}
	buff := &bytes.Buffer{}
	io.Copy(buff, body)
	cp := buff.String()
	err = json.NewDecoder(buff).Decode(val)
	if err != nil && Debug {
		logger.Println("trace: body: ", cp)
	}
	return err
}

func (c *client) GetBody(url string, auth bool, expTime time.Duration) (io.Reader, error) {
	if !IgnoreCache {
		body := c.cache.GetRequest(url, expTime)
		if body != "" {
			logger.Println(url)
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
	if !IgnoreCache {
		c.cache.StoreRequest(url, buff.String())
	}
	return buff, nil
}

func (c *client) Get(url string, auth bool) (*http.Response, error) {
	if !strings.HasPrefix(url, "http") {
		url = fmt.Sprintf("https://%s", url)
	}
	logger.Printf("trace: GET: %s", url)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if auth && os.Getenv("X_Riot_Token") != "" {
		r.Header.Add("X-Riot-Token", os.Getenv("X_Riot_Token"))
	}
	resp, err := c.client.Do(r)
	if err != nil {
		if Debug {
			logger.Printf("err: failed request: %v", err)
		}
		return resp, err
	}
	switch resp.StatusCode {
	case http.StatusTooManyRequests:
		if Debug {
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
		if Debug {
			logger.Println("trace: slow down charlie.\r")
		}
		resp, err = c.Get(url, auth)
		return resp, err
	case http.StatusNotFound:
		logger.Println("err: not found", url)
		return nil, fmt.Errorf("err: object not found: %s", url)
	case http.StatusOK, http.StatusAccepted:
		atomic.AddInt64(c.requestsSucceeded, 1)
		if Debug {
			fmt.Fprintf(os.Stdout, "RequestsSucceeded: %d\r", atomic.LoadInt64(c.requestsSucceeded))
		}
		return resp, err
	default:
		logger.Println("err: Code not expected:", resp.StatusCode)
		return nil, fmt.Errorf("err: %d %s", resp.StatusCode, url)
	}
}

func (c *client) Close() {
	c.cache.Close()
}
