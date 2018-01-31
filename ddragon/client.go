package ddragon

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/hibooboo2/lol"
	"github.com/hibooboo2/lol/cachedclient"
)

type client struct {
	c          *cachedclient.Client
	champsByID map[int]Champion
	itemsByID  map[int]Item
	realm      lol.Realms
	one        sync.Once
}

var DefaultClient *client

func NewClient(cc *cachedclient.Client, realms lol.Realms) *client {
	c := &client{c: cc, realm: realms}
	c.init()
	return c
}

func (c *client) init() {
	champs, err := c.Champs()
	if err != nil {
		log.Println(err)
		return
	}
	c.champsByID = make(map[int]Champion)
	for _, champ := range champs.Data {
		id, err := strconv.Atoi(champ.Key)
		if err == nil {
			c.champsByID[id] = champ
		}
	}
}

func init() {
	mongo, err := cachedclient.NewMongoCachedClient("", 0)
	if err != nil {
		panic(err)
	}

	mongo.Debug = true
	mongo.IgnoreExpiration = false

	cc := cachedclient.NewClient(string(lol.NA), mongo, func(r *http.Request) {
		r.Header.Add("X-Riot-Token", os.Getenv("X_Riot_Token"))
	})

	var relm lol.Realms
	err = cc.GetObjFromAPI("/lol/static-data/v3/realms", &relm, cachedclient.WEEK*1)

	if err != nil {
		log.Println(err)
		return
	}
	DefaultClient = NewClient(cc, relm)
}
