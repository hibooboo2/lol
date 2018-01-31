package ddragon

import (
	"net/http"
	"os"
	"testing"

	"github.com/comail/colog"
	"github.com/hibooboo2/lol"
	"github.com/hibooboo2/lol/cachedclient"
)

func TestGetChamps(t *testing.T) {
	mongo, err := cachedclient.NewMongoCachedClient("", 0)
	if err != nil {
		panic(err)
	}

	mongo.Debug = true
	mongo.IgnoreExpiration = false

	cc := cachedclient.NewClient(string(lol.NA), mongo, func(r *http.Request) {
		r.Header.Add("X-Riot-Token", os.Getenv("X_Riot_Token"))
	})
	cc.Debug = true
	cachedclient.SetLogLevel(colog.LTrace)

	var relm lol.Realms
	err = cc.GetObjFromAPI("/lol/static-data/v3/realms", &relm, cachedclient.WEEK*1)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	c := NewClient(cc, relm)

	_, err = c.Champ(420)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestGetItem(t *testing.T) {
	boot, err := DefaultClient.GetItem(1001)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(boot)
	if boot.Name != "Boots of Speed" {
		t.Fatal("did not get right name for item")
	}
	sr := 0
	for _, item := range DefaultClient.itemsByID {
		sell := float64(item.Gold.Sell) / float64(item.Gold.Total)
		if item.Maps["11"] && item.Gold.Purchasable && sell > 0.69 && len(item.Into) == 1 && len(item.From) == 1 {
			t.Log(item.Gold.Total, item.Gold.Sell, sell, item.Name, item.Into, item.From)
			sr++
		}
	}
	t.Log(sr)
}
