package cachedclient

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoCachedClient struct {
	session          *mgo.Session
	db               *mgo.Database
	requests         *mgo.Collection
	inMemoryCache    *memCache
	Debug            bool
	IgnoreExpiration bool
}

func NewMongoCachedClient(host string, port int) (*mongoCachedClient, error) {
	if host == "" {
		host = "localhost"
	}
	if port == 0 {
		port = 27017
	}

	session, err := mgo.Dial(fmt.Sprintf(`%s:%d`, host, port))
	if err != nil {
		return nil, err
	}
	mongo := &mongoCachedClient{
		session:  session,
		db:       session.DB("lol"),
		requests: session.DB("lol").C("requests"),
		inMemoryCache: &memCache{
			requests: make(map[string]request),
		},
	}
	err = mongo.requests.EnsureIndex(mgo.Index{
		Key:      []string{"url"},
		DropDups: true,
		Unique:   true,
	})
	if err != nil {
		return nil, err
	}
	return mongo, nil
}

func (db *mongoCachedClient) GetRequest(url string, expTime time.Duration) string {
	var r request
	if body := db.inMemoryCache.GetRequest(url, expTime); body != "" {
		if db.Debug {
			logger.Println("trace: From memory cache:", url)
		}
		return body
	}
	err := db.requests.Find(bson.M{"url": url}).One(&r)
	if err != nil {
		if db.Debug {
			logger.Println("err: failed to find request:", url)
		}
		return ""
	}
	if db.IgnoreExpiration {
		db.inMemoryCache.StoreRequest(r.Url, r.Body)
		if db.Debug {
			logger.Println("trace: From mongo cache:", url)
		}
		return r.Body
	}
	if time.Now().Before(r.Created.Add(expTime)) {
		if db.Debug {
			logger.Println("trace: From mongo cache:", url)
		}
		db.inMemoryCache.StoreRequest(r.Url, r.Body)
		return r.Body
	}
	err = db.requests.Remove(bson.M{"url": url})
	if db.Debug {
		if err != nil {
			logger.Println("err: failed to remove request stored:", url)
		}
		logger.Println("trace: not found:", url)
	}
	return ""
}

func (db *mongoCachedClient) StoreRequest(url string, body string) {
	var r request
	r.Url = url
	r.Body = body
	r.Created = time.Now()
	err := db.requests.Insert(&r)
	if err != nil {
		logger.Printf("err: url: %s failed to  store request body: \n %v", url, err)
		return
	}
	if db.Debug {
		logger.Println("trace: stored request: ", url)
	}
	db.inMemoryCache.StoreRequest(url, body)
}

func (db *mongoCachedClient) Close() {
	db.session.Close()
}

func (db *mongoCachedClient) SetIgnoreExpiration(ignore bool) {
	db.IgnoreExpiration = ignore
}

func (db *mongoCachedClient) SetDebug(debug bool) {
	db.Debug = debug
}
