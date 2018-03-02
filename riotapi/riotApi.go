package riotapi

import (
	"net/http"
	"os"

	"github.com/hibooboo2/lol/cachedclient"
	"github.com/hibooboo2/lol/constants"
)

// RiotClient is the interface that defines the interactions with riots api.
type RiotClient interface {
	Mastery() *champMastery
	Spectator() *spectator
	StaticData() *staticData
	// /lol/league/v3/challengerleagues/by-queue/{queue}
	// /lol/league/v3/leagues/by-summoner/{summonerId}
	// /lol/league/v3/masterleagues/by-queue/{queue}
	// /lol/league/v3/positions/by-summoner/{summonerId}
	// /lol/match/v3/matches/{matchId}
	// /lol/match/v3/matches/{matchId}/by-tournament-code/{tournamentCode}
	// /lol/match/v3/matches/by-tournament-code/{tournamentCode}/ids
	// /lol/match/v3/matchlists/by-account/{accountId}
	// /lol/match/v3/matchlists/by-account/{accountId}/recent
	// /lol/match/v3/timelines/by-match/{matchId}
	// /lol/platform/v3/champions
	// /lol/platform/v3/champions/{id}
	// /lol/platform/v3/masteries/by-summoner/{summonerId}
	// /lol/platform/v3/masteries/by-summoner/{summonerId}
	// /lol/platform/v3/runes/by-summoner/{summonerId}
	// /lol/static-data/v3/masteries
	// /lol/static-data/v3/masteries/{id}
	// /lol/static-data/v3/profile-icons
	// /lol/static-data/v3/runes
	// /lol/static-data/v3/runes/{id}
	// /lol/static-data/v3/summoner-spells
	// /lol/static-data/v3/summoner-spells/{id}

	// /lol/status/v3/shard-data
	Summoners() *summoners

	//GetCache Expose Cache
	GetCache() cachedclient.RequestCache
}

type api struct {
	client *cachedclient.Client
}

// NewClient returns a new client for performing operations using riots api.
func NewClient(region constants.RegionEndPoint) (RiotClient, error) {
	// cache, err := NewLolMongo("dev.jhrb.us", 27217)
	cache, err := cachedclient.NewMongoCachedClient("", 0)
	// cache, err := NewLolMongo("192.168.1.170", 27027)
	if err != nil {
		return nil, err
	}
	client := cachedclient.NewClient(string(region), cache, func(r *http.Request) {
		r.Header.Add("X-Riot-Token", os.Getenv("X_Riot_Token"))
	})
	if err != nil {
		return nil, err
	}
	return &api{
		client: client,
	}, nil
}

func (a *api) GetCache() cachedclient.RequestCache {
	return a.client.GetCache()
}
func (a *api) Mastery() *champMastery {
	return &champMastery{a.client}
}

func (a *api) Summoners() *summoners {
	return &summoners{a.client}
}

func (a *api) Spectator() *spectator {
	return &spectator{a.client}
}

func (a *api) StaticData() *staticData {
	return &staticData{a.client}
}
