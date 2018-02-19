package ddragon

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strconv"
	"strings"

	"github.com/hibooboo2/lol/cachedclient"
	"github.com/renstrom/fuzzysearch/fuzzy"
	"golang.org/x/net/html"
)

func (c *client) GetItems() (*ItemData, error) {
	var items ItemData
	err := c.c.GetObjNoBase(c.RealmLink(c.realm.LatestVersions.Item, "data/"+c.realm.DefaultLanguage, "item.json"), &items, cachedclient.WEEK*1)
	if err != nil {
		return nil, err
	}
	for key, item := range items.Items {
		id, _ := strconv.Atoi(key)
		item.Img = c.ItemSpriteLink(id)
		item.FixDesc()
		items.Items[key] = item
	}
	return &items, nil
}

func (c *client) GetItem(id int) (*Item, error) {
	if c.itemsByID == nil {
		return nil, fmt.Errorf("Items not found") //Probably should be a panic... Or just call at start of client init to make sure it works.
	}
	if item, ok := c.itemsByID[id]; ok {
		return &item, nil
	}
	err := fmt.Sprintf("Cannot find item %d", id)
	return &Item{Name: err}, fmt.Errorf(err)
}

func (c *client) SearchItems(itemName string) []Item {
	itemNames := fuzzy.FindFold(itemName, c.itemNames)
	items := []Item{}
	for _, name := range itemNames {
		items = append(items, c.itemsByName[name])
	}
	return items
}

func (i *Item) FixDesc() {
	node, err := html.Parse(strings.NewReader(string(i.Description)))
	if err != nil {
		log.Println(err)
		i.Description = ""
		return
	}
	buff := &bytes.Buffer{}
	err = html.Render(buff, node)
	if err != nil {
		log.Println(err)
		i.Description = ""
		return
	}

	newDesc := strings.NewReplacer("<html>", "", "</html>", "", "<head>", "", "</head>", "", "<body>", "", "</body>", "").Replace(buff.String())
	i.Description = template.HTML(newDesc)
}

type ItemData struct {
	Type    string          `json:"type"`
	Version string          `json:"version"`
	Items   map[string]Item `json:"data"`
	Groups  []ItemGroup     `json:"groups"`
	Tree    []ItemTree      `json:"tree"`
}

type ItemGroup struct {
	ID              string `json:"id"`
	MaxGroupOwnable string `json:"MaxGroupOwnable"`
}
type ItemTree struct {
	Header string   `json:"header"`
	Tags   []string `json:"tags"`
}
type ItemName string

func (i ItemName) String() string {
	id, err := strconv.Atoi(string(i))
	if err != nil {
		return string(i)
	}
	return DefaultClient().itemsByID[id].Name
}

func (i ItemName) Html() template.HTML {
	const itemTemplate = `<img src="%s" height="38"/>`
	id, err := strconv.Atoi(string(i))
	if err != nil {
		return template.HTML(i)
	}
	return template.HTML(fmt.Sprintf(itemTemplate, DefaultClient().itemsByID[id].Img))
}

type Item struct {
	Name             string            `json:"name"`
	Img              string            `json:"-"`
	Rune             RuneDTO           `json:"rune"`
	Gold             GoldDTO           `json:"gold"`
	Group            string            `json:"group"`
	Description      template.HTML     `json:"description"`
	Colloq           string            `json:"colloq"`
	Plaintext        string            `json:"plaintext"`
	Consumed         bool              `json:"consumed"`
	Stacks           int               `json:"stacks"`
	Depth            int               `json:"depth"`
	ConsumeOnFull    bool              `json:"consumeOnFull"`
	From             []ItemName        `json:"from"`
	Into             []ItemName        `json:"into"`
	SpecialRecipe    int               `json:"specialRecipe"`
	InStore          bool              `json:"inStore"`
	HideFromAll      bool              `json:"hideFromAll"`
	RequiredChampion string            `json:"requiredChampion"`
	Stats            StatsDTO          `json:"stats"`
	Tags             []string          `json:"tags"`
	Maps             map[string]bool   `json:"maps"`
	Effect           map[string]string `json:"effect"`
}

type GoldDTO struct {
	Base        int  `json:"base"`
	Total       int  `json:"total"`
	Sell        int  `json:"sell"`
	Purchasable bool `json:"purchasable"`
}

type RuneDTO struct {
	Isrune bool   `json:"isrune"`
	Tier   int    `json:"tier"`
	Type   string `json:"type"`
}
type StatsDTO struct {
	FlatHPPoolMod                       int     `json:"FlatHPPoolMod"`
	RFlatHPModPerLevel                  int     `json:"rFlatHPModPerLevel"`
	FlatMPPoolMod                       int     `json:"FlatMPPoolMod"`
	RFlatMPModPerLevel                  int     `json:"rFlatMPModPerLevel"`
	PercentHPPoolMod                    float64 `json:"PercentHPPoolMod"`
	PercentMPPoolMod                    float64 `json:"PercentMPPoolMod"`
	FlatHPRegenMod                      float64 `json:"FlatHPRegenMod"`
	RFlatHPRegenModPerLevel             int     `json:"rFlatHPRegenModPerLevel"`
	PercentHPRegenMod                   float64 `json:"PercentHPRegenMod"`
	FlatMPRegenMod                      int     `json:"FlatMPRegenMod"`
	RFlatMPRegenModPerLevel             int     `json:"rFlatMPRegenModPerLevel"`
	PercentMPRegenMod                   float64 `json:"PercentMPRegenMod"`
	FlatArmorMod                        int     `json:"FlatArmorMod"`
	RFlatArmorModPerLevel               int     `json:"rFlatArmorModPerLevel"`
	PercentArmorMod                     float64 `json:"PercentArmorMod"`
	RFlatArmorPenetrationMod            int     `json:"rFlatArmorPenetrationMod"`
	RFlatArmorPenetrationModPerLevel    int     `json:"rFlatArmorPenetrationModPerLevel"`
	RPercentArmorPenetrationMod         int     `json:"rPercentArmorPenetrationMod"`
	RPercentArmorPenetrationModPerLevel int     `json:"rPercentArmorPenetrationModPerLevel"`
	FlatPhysicalDamageMod               int     `json:"FlatPhysicalDamageMod"`
	RFlatPhysicalDamageModPerLevel      int     `json:"rFlatPhysicalDamageModPerLevel"`
	PercentPhysicalDamageMod            float64 `json:"PercentPhysicalDamageMod"`
	FlatMagicDamageMod                  int     `json:"FlatMagicDamageMod"`
	RFlatMagicDamageModPerLevel         int     `json:"rFlatMagicDamageModPerLevel"`
	PercentMagicDamageMod               float64 `json:"PercentMagicDamageMod"`
	FlatMovementSpeedMod                int     `json:"FlatMovementSpeedMod"`
	RFlatMovementSpeedModPerLevel       int     `json:"rFlatMovementSpeedModPerLevel"`
	PercentMovementSpeedMod             float64 `json:"PercentMovementSpeedMod"`
	RPercentMovementSpeedModPerLevel    int     `json:"rPercentMovementSpeedModPerLevel"`
	FlatAttackSpeedMod                  int     `json:"FlatAttackSpeedMod"`
	PercentAttackSpeedMod               float64 `json:"PercentAttackSpeedMod"`
	RPercentAttackSpeedModPerLevel      int     `json:"rPercentAttackSpeedModPerLevel"`
	RFlatDodgeMod                       int     `json:"rFlatDodgeMod"`
	RFlatDodgeModPerLevel               int     `json:"rFlatDodgeModPerLevel"`
	PercentDodgeMod                     float64 `json:"PercentDodgeMod"`
	FlatCritChanceMod                   float64 `json:"FlatCritChanceMod"`
	RFlatCritChanceModPerLevel          int     `json:"rFlatCritChanceModPerLevel"`
	PercentCritChanceMod                float64 `json:"PercentCritChanceMod"`
	FlatCritDamageMod                   int     `json:"FlatCritDamageMod"`
	RFlatCritDamageModPerLevel          int     `json:"rFlatCritDamageModPerLevel"`
	PercentCritDamageMod                float64 `json:"PercentCritDamageMod"`
	FlatBlockMod                        int     `json:"FlatBlockMod"`
	PercentBlockMod                     float64 `json:"PercentBlockMod"`
	FlatSpellBlockMod                   int     `json:"FlatSpellBlockMod"`
	RFlatSpellBlockModPerLevel          int     `json:"rFlatSpellBlockModPerLevel"`
	PercentSpellBlockMod                float64 `json:"PercentSpellBlockMod"`
	FlatEXPBonus                        int     `json:"FlatEXPBonus"`
	PercentEXPBonus                     float64 `json:"PercentEXPBonus"`
	RPercentCooldownMod                 int     `json:"rPercentCooldownMod"`
	RPercentCooldownModPerLevel         int     `json:"rPercentCooldownModPerLevel"`
	RFlatTimeDeadMod                    int     `json:"rFlatTimeDeadMod"`
	RFlatTimeDeadModPerLevel            int     `json:"rFlatTimeDeadModPerLevel"`
	RPercentTimeDeadMod                 int     `json:"rPercentTimeDeadMod"`
	RPercentTimeDeadModPerLevel         int     `json:"rPercentTimeDeadModPerLevel"`
	RFlatGoldPer10Mod                   int     `json:"rFlatGoldPer10Mod"`
	RFlatMagicPenetrationMod            int     `json:"rFlatMagicPenetrationMod"`
	RFlatMagicPenetrationModPerLevel    int     `json:"rFlatMagicPenetrationModPerLevel"`
	RPercentMagicPenetrationMod         int     `json:"rPercentMagicPenetrationMod"`
	RPercentMagicPenetrationModPerLevel int     `json:"rPercentMagicPenetrationModPerLevel"`
	FlatEnergyRegenMod                  int     `json:"FlatEnergyRegenMod"`
	RFlatEnergyRegenModPerLevel         int     `json:"rFlatEnergyRegenModPerLevel"`
	FlatEnergyPoolMod                   int     `json:"FlatEnergyPoolMod"`
	RFlatEnergyModPerLevel              int     `json:"rFlatEnergyModPerLevel"`
	PercentLifeStealMod                 float64 `json:"PercentLifeStealMod"`
	PercentSpellVampMod                 float64 `json:"PercentSpellVampMod"`
}
