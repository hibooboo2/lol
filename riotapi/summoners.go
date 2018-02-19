package riotapi

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/hibooboo2/lol/cachedclient"
	"github.com/hibooboo2/lol/constants"
)

// /lol/summoner/v3/summoners/{summonerId}
// /lol/summoner/v3/summoners/by-account/{accountId}
// /lol/summoner/v3/summoners/by-name/{summonerName}
type summoners struct {
	c *cachedclient.Client
}

func (s *summoners) Arg(arg string) (*Summoner, error) {
	sum := s.ByName(arg)
	if sum != nil {
		return sum, nil
	}
	id, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		return nil, err
	}
	sum = s.ByAccountID(id)
	sum2 := s.ID(id)
	switch {
	case sum != nil && sum2 != nil:
		return nil, fmt.Errorf("Not sure if using accountID or summonerID: %s", arg)
	case sum != nil:
		return sum, nil
	case sum2 != nil:
		return sum2, nil
	default:
		return nil, fmt.Errorf("Cannot locate summoner: %s", arg)
	}
}

// /lol/summoner/v3/summoners/{summonerId}
func (s *summoners) ID(summonerID int64) *Summoner {
	var sum Summoner
	err := s.c.GetObjFromAPI(fmt.Sprintf("/lol/summoner/v3/summoners/%d", summonerID), &sum, constants.DAY)
	if err != nil {
		return nil
	}
	return &sum
}

// /lol/summoner/v3/summoners/by-account/{accountId}
func (s *summoners) ByAccountID(accountID int64) *Summoner {
	var sum Summoner
	err := s.c.GetObjFromAPI(fmt.Sprintf(`/lol/summoner/v3/summoners/by-account/%d`, accountID), &sum, constants.DAY)
	if err != nil {
		return nil
	}
	return &sum
}

// /lol/summoner/v3/summoners/by-name/{summonerName}
func (s *summoners) ByName(summonerName string) *Summoner {
	var sum Summoner
	err := s.c.GetObjFromAPI(fmt.Sprintf("/lol/summoner/v3/summoners/by-name/%s", url.PathEscape(summonerName)), &sum, constants.DAY)
	if err != nil {
		// logger.Println("err: ", err)
		return nil
	}
	return &sum
}

type Summoner struct {
	ID            int64  `json:"id"`
	AccountID     int64  `json:"accountId"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}
