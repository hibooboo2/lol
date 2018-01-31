package ddragon

import (
	"fmt"

	"github.com/hibooboo2/lol/cachedclient"
)

// /lol/static-data/v3/champions?tags=all
//Swaped with ddragon so it doesn't die.
func (c *client) Champs() (*ChampionList, error) {
	var champs ChampionList
	err := c.c.GetObjNoBase(c.RealmLink(c.realm.LatestVersions.Champion, "data/"+c.realm.DefaultLanguage, "champion.json"), &champs, cachedclient.WEEK*1)
	if err != nil {
		return nil, err
	}
	return &champs, nil
}

// /lol/static-data/v3/champions/{id} Actually using datadragon...
func (c *client) Champ(id int) (*Champion, error) {
	champ, ok := c.champsByID[id]
	if !ok {
		err := fmt.Sprintf("Cannot find champ %d", id)
		return &Champion{Name: err}, fmt.Errorf(err)
	}
	return &champ, nil
}

type ChampionList struct {
	Data    map[string]Champion
	Format  string `json:"format"`
	Type    string `json:"type"`
	Version string `json:"version"`
}

type Champion struct {
	Version string `json:"version"`
	ID      string `json:"id"`
	Key     string `json:"key"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Blurb   string `json:"blurb"`
	Info    struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
	Image struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	} `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   struct {
		Hp                   float64 `json:"hp"`
		Hpperlevel           float64 `json:"hpperlevel"`
		Mp                   float64 `json:"mp"`
		Mpperlevel           float64 `json:"mpperlevel"`
		Movespeed            float64 `json:"movespeed"`
		Armor                float64 `json:"armor"`
		Armorperlevel        float64 `json:"armorperlevel"`
		Spellblock           float64 `json:"spellblock"`
		Spellblockperlevel   float64 `json:"spellblockperlevel"`
		Attackrange          float64 `json:"attackrange"`
		Hpregen              float64 `json:"hpregen"`
		Hpregenperlevel      float64 `json:"hpregenperlevel"`
		Mpregen              float64 `json:"mpregen"`
		Mpregenperlevel      float64 `json:"mpregenperlevel"`
		Crit                 float64 `json:"crit"`
		Critperlevel         float64 `json:"critperlevel"`
		Attackdamage         float64 `json:"attackdamage"`
		Attackdamageperlevel float64 `json:"attackdamageperlevel"`
		Attackspeedoffset    float64 `json:"attackspeedoffset"`
		Attackspeedperlevel  float64 `json:"attackspeedperlevel"`
	} `json:"stats"`
}
