package cachedclient

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/comail/colog"
	"github.com/hibooboo2/lol"
)

func main2() {
	// defer profile.Start().Stop()
	SetLogLevel(colog.LTrace)

	mongo, err := NewMongoCachedClient("", 0)
	if err != nil {
		panic(err)
	}

	mongo.Debug = true
	mongo.IgnoreExpiration = false

	c := &Client{
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
	}
	var s lol.Summoner
	err = c.GetObjFromAPI(fmt.Sprintf("/lol/summoner/v3/summoners/by-name/%s", "sirfxwright"), &s, time.Hour*24)
	if err == nil {
		log.Println(s)
	} else {
		log.Println(err)
	}
	err = c.GetObjFromAPI(fmt.Sprintf("/lol/summoner/v3/summoners/by-name/%s", "sirfxwright"), &s, time.Hour*24)

}
