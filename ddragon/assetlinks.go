package ddragon

import "fmt"

func (c *client) SummonerSpellLink(id int) string {
	s, err := c.SummonerSpell(id)
	if err != nil {
		logger.Println("err: when getting summonerspell: ", err)
		return ""
	}
	return c.RealmLink(c.realm.LatestVersions.Summoner, "img/spell", s.ID+".png")
}

func (c *client) ChampionSpriteLink(id int) string {
	champ, err := c.Champ(id)
	if err != nil {
		logger.Println("err: when getting champion: ", err)
		return ""
	}
	return c.RealmLink(c.realm.LatestVersions.Champion, "img/champion", champ.ID+".png")
}

func (c *client) RealmLink(resourceVersion, resource, key string) string {
	return fmt.Sprintf("%s/%s/%s/%s", c.realm.Cdn, resourceVersion, resource, key)
}

func (c *client) ProfileIconLink(id int) string {
	return c.RealmLink(c.realm.LatestVersions.Profileicon, "img/profileicon", fmt.Sprintf("%d.png", id))
}

func (c *client) ItemSpriteLink(id int) string {
	return c.RealmLink(c.realm.LatestVersions.Item, "img/item", fmt.Sprintf("%d.png", id))
}
