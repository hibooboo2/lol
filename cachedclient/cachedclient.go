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

var DefaultClient *Client

func init() {
	mongo, err := NewMongoCachedClient("", 0)
	if err != nil {
		log.Println("err: failed to connect to mongo. Using in memory cache for default client: ", err)
	}
	mongo.IgnoreExpiration = true

	DefaultClient = &Client{
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
		// Debug:            true,
	}
	if mongo == nil {
		DefaultClient.cache = &memCache{
			requests: make(map[string]request),
		}
	}
}
