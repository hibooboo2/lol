package cachedclient

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hibooboo2/lol"
)

type RequestCache interface {
	StoreRequest(url string, body string)
	GetRequest(url string, expTime time.Duration) string
	SetIgnoreExpiration(ignoreExpiration bool)
	SetDebug(debug bool)
}

//DefaultClient is the default cached client for the package. It has an auth function for riots api and
//uses na server as well as an in memory cache.
func DefaultClient() *Client {
	return &Client{
		Auth: func(r *http.Request) {
			r.Header.Add("X-Riot-Token", os.Getenv("X_Riot_Token"))
		},
		cache: &memCache{
			requests: make(map[string]request),
		},
		baseURL: string(lol.NA),
		client: &http.Client{
			Timeout: time.Second * 5,
		},
		IgnoreExpiration: true,
		Debug:            true,
	}
}

func DefaultMongoClient() (*Client, error) {
	mongo, err := NewMongoCachedClient("", 0)
	if err != nil {
		log.Println("err: failed to connect to mongo. Using in memory cache for default client: ", err)
		return nil, err
	}
	mongo.IgnoreExpiration = true

	return &Client{
		Auth: func(r *http.Request) {
			r.Header.Add("X-Riot-Token", os.Getenv("X_Riot_Token"))
		},
		cache:   mongo,
		baseURL: string(lol.NA),
		client: &http.Client{
			Timeout: time.Second * 5,
		},
		// IgnoreCache:      true,
		IgnoreExpiration: true,
		Debug:            true,
	}, nil
}
