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
	c           *cachedclient.Client
	champsByID  map[int]Champion
	itemsByID   map[int]Item
	itemsByName map[string]Item
	itemNames   []string
	realm       lol.Realms
	one         sync.Once
}

var dclient sync.Once

//DefaultClient the default ddragon client
func DefaultClient() *client {
	var c *client
	dclient.Do(func() {
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
		err = cc.GetObjNoBase("http://ddragon.leagueoflegends.com/realms/na.json", &relm, cachedclient.WEEK*1)

		if err != nil {
			log.Println(err)
			return
		}
		c = NewClient(cc, relm)
	})
	return c
}

func NewClient(cc *cachedclient.Client, realms lol.Realms) *client {
	c := &client{c: cc, realm: realms}
	c.init()
	return c
}

func (c *client) init() {
	champs, err := c.Champs()
	if err != nil {
		log.Println(err)

	} else {
		c.champsByID = make(map[int]Champion)
		for _, champ := range champs.Data {
			id, err := strconv.Atoi(champ.Key)
			if err == nil {
				c.champsByID[id] = champ
			}
		}
	}

	items, err := c.GetItems()
	if err != nil {
		logger.Println("err: failed to initialize items: ", err)

	} else {
		c.itemsByID = make(map[int]Item)
		c.itemsByName = make(map[string]Item)
		for key, item := range items.Items {
			id, err := strconv.Atoi(key)
			if err != nil {
				logger.Println("err: failed to parse item id: ", key, err)
				continue
			}
			c.itemsByName[item.Name] = item
			c.itemsByID[id] = item
			c.itemNames = append(c.itemNames, item.Name)
		}
	}
}
