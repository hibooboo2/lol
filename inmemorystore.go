package lol

import (
	"sync"
	"time"

	"golang.org/x/sync/syncmap"
)

type memCache struct {
	games    syncmap.Map
	toVisit  syncmap.Map
	visit    chan int64
	visited  syncmap.Map
	requests map[string]request
	lock     sync.RWMutex
}

func (c *memCache) HaveGame(gameID int64) bool {
	_, ok := c.games.Load(gameID)
	return ok
}

func (c *memCache) AddGame(gameID int64) {
	c.games.Store(gameID, struct{}{})
}

func (c *memCache) Player(accountID int64) {
	_, ok := c.visited.Load(accountID)
	if ok {
		return
	}
	c.toVisit.Store(accountID, struct{}{})
}

func (c *memCache) VisitedPlayer(accountID int64) {
	c.toVisit.Delete(accountID)
	c.visited.Store(accountID, struct{}{})
}

func (c *memCache) GetPlayerToVisit() int64 {
	var id int64
	var ok bool
	c.toVisit.Range(func(key interface{}, value interface{}) bool {
		id, ok = key.(int64)
		return false
	})
	if id == 0 {
		return 0
	}
	c.VisitedPlayer(id)
	return id
}

func (c *memCache) HaveVisitedPlayer(accountID int64) bool {
	_, ok := c.visited.Load(accountID)
	return ok
}

func (c *memCache) GetRequest(url string, expTime time.Duration) string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	r, ok := c.requests[url]
	if !ok {
		return ""
	}
	if CacheNoExpire {
		return r.Body
	}
	if time.Now().Before(r.Created.Add(expTime)) {
		if Debug {
			logger.Println("trace: From inmemory cache:", url)
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

func (c *memCache) StoreRequest(r request) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.requests[r.Url] = r
	if Debug {
		logger.Println("trace: stored request: ", r.Url)
	}
}
