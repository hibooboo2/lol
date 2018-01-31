package cachedclient

import (
	"sync"
	"time"
)

type memCache struct {
	requests         map[string]request
	lock             sync.RWMutex
	IgnoreExpiration bool
	Debug            bool
}

type request struct {
	Body    string
	Url     string
	Created time.Time
}

func (c *memCache) GetRequest(url string, expTime time.Duration) string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	r, ok := c.requests[url]
	if !ok {
		return ""
	}
	if c.IgnoreExpiration {
		if c.Debug {
			logger.Println("debug: From inMemorycache ignoredExpiration:", url)
		}
		return r.Body
	}
	if expTime == 0 {
		return r.Body
	}
	if time.Now().Before(r.Created.Add(expTime)) {
		if c.Debug {
			logger.Println("debug: From inMemorycache:", url)
		}
		return r.Body
	}
	c.lock.RUnlock()
	c.lock.Lock()
	delete(c.requests, url)
	c.lock.Unlock()
	c.lock.RLock()
	return ""
}

func (c *memCache) StoreRequest(url string, body string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.requests[url] = request{
		Url:     url,
		Body:    body,
		Created: time.Now(),
	}
}

func (c *memCache) SetIgnoreExpiration(ignore bool) {
	c.IgnoreExpiration = ignore
}

func (c *memCache) SetDebug(debug bool) {
	c.Debug = debug
}
