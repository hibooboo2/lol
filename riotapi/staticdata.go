package riotapi

import (
	"fmt"

	"github.com/hibooboo2/lol/cachedclient"
)

type staticData struct {
	c *cachedclient.Client
}

// /lol/static-data/v3/champions
func (sd *staticData) Champs() (*ChampionList, error) {
	var obj ChampionList
	err := sd.c.GetObjFromAPI("/lol/static-data/v3/champions?tags=all", &obj, cachedclient.WEEK*2)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// /lol/static-data/v3/champions/{id}
func (sd *staticData) Champ(id int) (*Champion, error) {
	var c Champion
	err := sd.c.GetObjFromAPI(fmt.Sprintf("/lol/static-data/v3/champions/%d?tags=all", id), &c, cachedclient.WEEK*2)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// /lol/static-data/v3/items
func (sd *staticData) Items() (*ItemList, error) {
	var items ItemList
	err := sd.c.GetObjFromAPI(fmt.Sprintf("/lol/static-data/v3/items"), &items, cachedclient.WEEK*2)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

// /lol/static-data/v3/items/{id}
func (sd *staticData) Item(id int) (*Item, error) {
	var i Item
	err := sd.c.GetObjFromAPI(fmt.Sprintf("/lol/static-data/v3/items/%d", id), &i, cachedclient.WEEK*2)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

// /lol/static-data/v3/language-strings
func (sd *staticData) LanguageStrings(id int) (*LangStrings, error) {
	var obj LangStrings
	err := sd.c.GetObjFromAPI(fmt.Sprintf("/lol/static-data/v3/language-strings"), &obj, cachedclient.WEEK*4)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// /lol/static-data/v3/languages
func (sd *staticData) Languages(id int) ([]string, error) {
	var langs []string
	err := sd.c.GetObjFromAPI(fmt.Sprintf("/lol/static-data/v3/languages"), &langs, cachedclient.WEEK*4)
	if err != nil {
		return nil, err
	}
	return langs, nil
}

// /lol/static-data/v3/maps
func (sd *staticData) Maps(id int) (*MapsObject, error) {
	var maps MapsObject
	err := sd.c.GetObjFromAPI(fmt.Sprintf("/lol/static-data/v3/maps"), &maps, cachedclient.WEEK*4)
	if err != nil {
		return nil, err
	}
	return &maps, nil
}

// /lol/static-data/v3/versions
func (sd *staticData) Versions() ([]string, error) {
	var versions []string
	err := sd.c.GetObjFromAPI("/lol/static-data/v3/versions", &versions, cachedclient.WEEK*1)
	return versions, err
}

// /lol/static-data/v3/realms
func (sd *staticData) Realms() (Realms, error) {
	var relm Realms
	err := sd.c.GetObjFromAPI("/lol/static-data/v3/realms", &relm, cachedclient.WEEK*1)
	return relm, err
}

//Realms represents info for a relm.
type Realms struct {
	Lg              string `json:"lg"`
	Dd              string `json:"dd"`
	DefaultLanguage string `json:"l"`
	LatestVersions  struct {
		Summoner    string `json:"summoner"`
		Map         string `json:"map"`
		Champion    string `json:"champion"`
		Language    string `json:"language"`
		Mastery     string `json:"mastery"`
		Sticker     string `json:"sticker"`
		Item        string `json:"item"`
		Rune        string `json:"rune"`
		Profileicon string `json:"profileicon"`
	} `json:"n"`
	Profileiconmax        int    `json:"profileiconmax"`
	CurrentVersionOfRealm string `json:"v"`
	Cdn                   string `json:"cdn"`
	CSS                   string `json:"css"`
}

type LangStrings struct {
	Data    map[string]string
	Version string
	Type    string
}

type MapsObject struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Data    map[string]struct {
		MapName string `json:"mapName"`
		MapID   int    `json:"mapId"`
		Image   struct {
			Full   string `json:"full"`
			Sprite string `json:"sprite"`
			Group  string `json:"group"`
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
		} `json:"image"`
	} `json:"data"`
}

type ChampionList struct {
	Data    map[string]Champion
	Keys    map[string]string `json:"keys"`
	Format  string            `json:"format"`
	Type    string            `json:"type"`
	Version string            `json:"version"`
}

type Champion struct {
	Image struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	} `json:"image"`
	Lore     string   `json:"lore"`
	Partype  string   `json:"partype"`
	Title    string   `json:"title"`
	Blurb    string   `json:"blurb"`
	Allytips []string `json:"allytips"`
	Passive  struct {
		Name                 string `json:"name"`
		Description          string `json:"description"`
		SanitizedDescription string `json:"sanitizedDescription"`
		Image                struct {
			Full   string `json:"full"`
			Sprite string `json:"sprite"`
			Group  string `json:"group"`
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
		} `json:"image"`
	} `json:"passive"`
	Tags        []string `json:"tags"`
	Recommended []struct {
		Champion string `json:"champion"`
		Title    string `json:"title"`
		Type     string `json:"type"`
		Map      string `json:"map"`
		Mode     string `json:"mode"`
		Blocks   []struct {
			Type    string `json:"type"`
			RecMath bool   `json:"recMath"`
			Items   []struct {
				ID    int `json:"id"`
				Count int `json:"count"`
			} `json:"items"`
		} `json:"blocks"`
	} `json:"recommended"`
	Skins []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Num  int    `json:"num"`
	} `json:"skins"`
	Stats struct {
		Armor                float64 `json:"armor"`
		Armorperlevel        float64 `json:"armorperlevel"`
		Attackdamage         float64 `json:"attackdamage"`
		Attackdamageperlevel float64 `json:"attackdamageperlevel"`
		Attackrange          float64 `json:"attackrange"`
		Attackspeedoffset    float64 `json:"attackspeedoffset"`
		Attackspeedperlevel  float64 `json:"attackspeedperlevel"`
		Crit                 float64 `json:"crit"`
		Critperlevel         float64 `json:"critperlevel"`
		Hp                   float64 `json:"hp"`
		Hpperlevel           float64 `json:"hpperlevel"`
		Hpregen              float64 `json:"hpregen"`
		Hpregenperlevel      float64 `json:"hpregenperlevel"`
		Movespeed            float64 `json:"movespeed"`
		Mp                   float64 `json:"mp"`
		Mpperlevel           float64 `json:"mpperlevel"`
		Mpregen              float64 `json:"mpregen"`
		Mpregenperlevel      float64 `json:"mpregenperlevel"`
		Spellblock           float64 `json:"spellblock"`
		Spellblockperlevel   float64 `json:"spellblockperlevel"`
	} `json:"stats"`
	Enemytips []string `json:"enemytips"`
	Name      string   `json:"name"`
	ID        int      `json:"id"`
	Spells    []struct {
		Name                 string `json:"name"`
		Description          string `json:"description"`
		SanitizedDescription string `json:"sanitizedDescription"`
		Tooltip              string `json:"tooltip"`
		SanitizedTooltip     string `json:"sanitizedTooltip"`
		Leveltip             struct {
			Label  []string `json:"label"`
			Effect []string `json:"effect"`
		} `json:"leveltip"`
		Image struct {
			Full   string `json:"full"`
			Sprite string `json:"sprite"`
			Group  string `json:"group"`
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
		} `json:"image"`
		Resource     string        `json:"resource"`
		Maxrank      int           `json:"maxrank"`
		Cost         []int         `json:"cost"`
		CostType     string        `json:"costType"`
		CostBurn     string        `json:"costBurn"`
		Cooldown     []float64     `json:"cooldown"`
		CooldownBurn string        `json:"cooldownBurn"`
		Effect       []interface{} `json:"effect"`
		EffectBurn   []string      `json:"effectBurn"`
		Vars         []struct {
			Key   string    `json:"key"`
			Link  string    `json:"link"`
			Coeff []float64 `json:"coeff"`
		} `json:"vars,omitempty"`
		Range     []int  `json:"range"`
		RangeBurn string `json:"rangeBurn"`
		Key       string `json:"key"`
	} `json:"spells"`
	Key  string `json:"key"`
	Info struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
}
