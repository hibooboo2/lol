package ddragon

import (
	"fmt"
	"log"

	"github.com/hibooboo2/lol/cachedclient"
)

func (c *client) SummonerSpells() (*SummonerSpellList, error) {
	var summoners SummonerSpellList
	err := c.c.GetObjNoBase(c.RealmLink(c.realm.LatestVersions.Summoner, "data/"+c.realm.DefaultLanguage, "summoner.json"), &summoners, cachedclient.WEEK*2)
	if err != nil {
		return nil, err
	}
	return &summoners, nil
}

func (c *client) SummonerSpell(id int) (*SummonerSpell, error) {
	spells, err := c.SummonerSpells()
	if err != nil {
		return nil, err
	}
	for _, s := range spells.Data {
		if s.Key == fmt.Sprint(id) {
			log.Println(s.ID)
			return &s, nil
		}
	}
	return nil, fmt.Errorf("Unable to find summonerspell: %d", id)
}

type SummonerSpellList struct {
	Data    map[string]SummonerSpell `json:"data"`
	Type    string                   `json:"type"`
	Version string                   `json:"version"`
}

type SummonerSpell struct {
	Description   string `json:"description"`
	ID            string `json:"id"`
	Key           string `json:"key"`
	Name          string `json:"name"`
	SummonerLevel int    `json:"summonerLevel"`
}
