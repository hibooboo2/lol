package ddragon

import (
	"log"
	"strconv"
	"sync"

	"github.com/hibooboo2/lol/cachedclient"
	"github.com/hibooboo2/lol/riotapi"
)

type client struct {
	c           *cachedclient.Client
	champsByID  map[int]Champion
	itemsByID   map[int]Item
	itemsByName map[string]Item
	itemNames   []string
	realm       riotapi.Realms
	one         sync.Once
}

var dclient sync.Once

//DefaultClient the default ddragon client
func DefaultClient() *client {
	var c *client
	dclient.Do(func() {
		cc := cachedclient.DefaultClient()

		var relm riotapi.Realms
		err := cc.GetObjNoBase("http://ddragon.leagueoflegends.com/realms/na.json", &relm, cachedclient.WEEK*1)

		if err != nil {
			log.Println(err)
			return
		}
		c = NewClient(cc, relm)
	})
	return c
}

func NewClient(cc *cachedclient.Client, realms riotapi.Realms) *client {
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
