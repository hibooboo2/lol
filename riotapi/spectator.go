package riotapi

import (
	"fmt"
	"net/url"
	"time"

	"github.com/hibooboo2/lol/cachedclient"
	"github.com/hibooboo2/lol/constants"
)

type spectator struct {
	// /lol/spectator/v3/active-games/by-summoner/{summonerId}
	// /lol/spectator/v3/featured-games
	c *cachedclient.Client
}

func (s *spectator) Game(summonerID int64) *CurrentGameInfo {
	var g CurrentGameInfo
	err := s.c.GetObjFromAPI(fmt.Sprintf("/lol/spectator/v3/active-games/by-summoner/%d", summonerID), &g, time.Minute*15)
	if err != nil {
		return nil
	}
	if g.GameID == 0 {
		return nil
	}
	return &g
}

func (s *spectator) GameSummonerName(summonerName string) *CurrentGameInfo {
	var sum Summoner
	err := s.c.GetObjFromAPI(fmt.Sprintf("/lol/summoner/v3/summoners/by-name/%s", url.PathEscape(summonerName)), &sum, constants.DAY)
	if err != nil {
		//XXX print error... we need to either start returning errors or having error logging
		return nil
	}
	// logger.Println("trace: game: ", g)
	return s.Game(sum.ID)
}

func (s *spectator) Featured() *FeaturedGames {
	var g FeaturedGames
	err := s.c.GetObjFromAPI("/lol/spectator/v3/featured-games", &g, time.Minute*10)
	if err != nil {
		return nil
	}
	return &g
}

type FeaturedGames struct {
	GameList []struct {
		GameID            int64  `json:"gameId"`
		MapID             int    `json:"mapId"`
		GameMode          string `json:"gameMode"`
		GameType          string `json:"gameType"`
		GameQueueConfigID int    `json:"gameQueueConfigId"`
		Participants      []struct {
			TeamID        int    `json:"teamId"`
			Spell1ID      int    `json:"spell1Id"`
			Spell2ID      int    `json:"spell2Id"`
			ChampionID    int    `json:"championId"`
			ProfileIconID int    `json:"profileIconId"`
			SummonerName  string `json:"summonerName"`
			Bot           bool   `json:"bot"`
		} `json:"participants"`
		Observers struct {
			EncryptionKey string `json:"encryptionKey"`
		} `json:"observers"`
		PlatformID      string        `json:"platformId"`
		BannedChampions []interface{} `json:"bannedChampions"`
		GameStartTime   int64         `json:"gameStartTime"`
		GameLength      int           `json:"gameLength"`
	} `json:"gameList"`
	ClientRefreshInterval int `json:"clientRefreshInterval"`
}

type CurrentGameInfo struct {
	GameID            int64                 `json:"gameId"`
	MapID             int                   `json:"mapId"`
	GameMode          string                `json:"gameMode"`
	GameType          string                `json:"gameType"`
	GameQueueConfigID int                   `json:"gameQueueConfigId"`
	Participants      []LiveGameParticipant `json:"participants"`
	Observers         struct {
		EncryptionKey string `json:"encryptionKey"`
	} `json:"observers"`
	PlatformID      string     `json:"platformId"`
	BannedChampions []ChampBan `json:"bannedChampions"`
	GameStartTime   int        `json:"gameStartTime"`
	GameLength      int        `json:"gameLength"`
}

type LiveGameParticipant struct {
	BannedImg                string        `json:"-"`
	TeamID                   int           `json:"teamId"`
	Spell1ID                 int           `json:"spell1Id"`
	Spell2ID                 int           `json:"spell2Id"`
	Spell1Img                string        `json:"-"`
	Spell2Img                string        `json:"-"`
	ChampionID               int           `json:"championId"`
	ChampionImage            string        `json:"-"`
	ProfileIconID            int           `json:"profileIconId"`
	ProfileIconImage         string        `json:"_"`
	SummonerName             string        `json:"summonerName"`
	Bot                      bool          `json:"bot"`
	SummonerID               int           `json:"summonerId"`
	GameCustomizationObjects []interface{} `json:"gameCustomizationObjects"`
	Perks                    struct {
		PerkIds      []int `json:"perkIds"`
		PerkStyle    int   `json:"perkStyle"`
		PerkSubStyle int   `json:"perkSubStyle"`
	} `json:"perks"`
}

type ChampBan struct {
	TeamID   int    `json:"teamId"`
	ChampID  int    `json:"championId"`
	PickTurn int    `json:"pickTurn"`
	Img      string `json:"-"`
}
