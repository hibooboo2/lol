package lol

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"time"

	"github.com/hibooboo2/lol/riotapi"
)

type Team struct {
	TeamID               int    `json:"teamId"`
	Win                  string `json:"win"`
	FirstBlood           bool   `json:"firstBlood"`
	FirstTower           bool   `json:"firstTower"`
	FirstInhibitor       bool   `json:"firstInhibitor"`
	FirstBaron           bool   `json:"firstBaron"`
	FirstDragon          bool   `json:"firstDragon"`
	FirstRiftHerald      bool   `json:"firstRiftHerald"`
	TowerKills           int    `json:"towerKills"`
	InhibitorKills       int    `json:"inhibitorKills"`
	BaronKills           int    `json:"baronKills"`
	DragonKills          int    `json:"dragonKills"`
	VilemawKills         int    `json:"vilemawKills"`
	RiftHeraldKills      int    `json:"riftHeraldKills"`
	DominionVictoryScore int    `json:"dominionVictoryScore"`
	Bans                 []Ban  `json:"bans"`
}

type Ban struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}

type Game struct {
	GameID                int64                 `json:"gameId"`
	PlatformID            string                `json:"platformId"`
	GameCreation          int64                 `json:"gameCreation"`
	GameDuration          int                   `json:"gameDuration"`
	QueueID               int                   `json:"queueId"`
	MapID                 int                   `json:"mapId"`
	SeasonID              int                   `json:"seasonId"`
	GameVersion           string                `json:"gameVersion"`
	GameMode              string                `json:"gameMode"`
	GameType              string                `json:"gameType"`
	Teams                 []Team                `json:"teams"`
	Participants          []Participant         `json:"participants"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	Cached                bool                  `json:"-"`
}

func (g *Game) Created() time.Time {
	return time.Unix(0, g.GameCreation*int64(time.Millisecond))
}

func (g *Game) Length() time.Duration {
	return time.Duration(g.GameDuration) * time.Second
}

type Games []Game

func (g Games) Len() int {
	return len(g)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (g Games) Less(i, j int) bool {
	return g[i].GameCreation < g[j].GameCreation
}

// Swap swaps the elements with indexes i and j.
func (g Games) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g Games) Sort() {
	sort.Sort(g)
}

func (g *Games) Add(game Game) {
	*g = append(*g, game)
}

type ParticipantIdentity struct {
	ParticipantID int    `json:"participantId"`
	Player        Player `json:"player"`
}

// WebMatch circumvent riots api throttling. Or at least attepmt to. This is using the endpoint that the web ui uses. No docs for it.
func WebMatch(gameID int64, currentPlatformID string, useCache bool) (*Game, error) {
	var game Game
	var err error
	if useCache {
		logger.Println("Trying cache")
		game, err = c.cache.GetGame(gameID, currentPlatformID)
		if err == nil {
			game.Cached = true
			return &game, nil
		}
	}
	logger.Println("Requesting")
	query := make(url.Values)
	// query.Add("visibleAccountId", fmt.Sprintf(`%d`, accountID))
	query.Add("visiblePlatformId", currentPlatformID)
	//https://acs.leagueoflegends.com/v1/stats/game/NA1/2591856267?visiblePlatformId=NA1&visibleAccountId=237823602

	reqURL, err := url.Parse(fmt.Sprintf("https://acs.leagueoflegends.com/v1/stats/game/%s/%d", currentPlatformID, gameID))
	reqURL.RawQuery = query.Encode()
	if err != nil {
		return nil, err
	}
	resp, err := c.Get(reqURL.String(), false)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Expected 200 code got: %d", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&game)
	if err != nil {
		return nil, err
	}
	err = c.cache.SaveGame(game, currentPlatformID)
	if err != nil {
		// logger.Println("err: Failed to save game to db / cache", err)
	}
	return &game, nil
}

func HaveMatch(gameID int64) bool {
	return c.cache.lolCache.HaveGame(gameID)
}

type Player struct {
	PlatformID        string `json:"platformId"`
	AccountID         int64  `json:"accountId"`
	SummonerName      string `json:"summonerName"`
	SummonerID        int64  `json:"summonerId"`
	CurrentPlatformID string `json:"currentPlatformId"`
	CurrentAccountID  int64  `json:"currentAccountId"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	ProfileIcon       int    `json:"profileIcon"`
}

func (p Player) ToSummoner() riotapi.Summoner {
	return riotapi.Summoner{
		ID:            p.SummonerID,
		Name:          p.SummonerName,
		AccountID:     p.AccountID,
		ProfileIconID: p.ProfileIcon,
	}
}
