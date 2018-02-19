package main

import (
	"fmt"
	"os"
	"strings"
)

// GetAllGames Gets all the games forever for an account.
func GetAllGames(accountID int64, platformID string) (Games, error) {
	var games Games
	var info *GamesInfoWebUiResponse
	var err error
	info, err = c.WebMatchHistory(accountID, platformID, 0)
	if err != nil {
		switch platformID {
		case "NA":
			platformID = NA1
		case NA1:
			platformID = "NA"
		}
		info, err = c.WebMatchHistory(accountID, platformID, 0)
	}
	if err != nil {
		return nil, err
	}

	for info.Games.GameIndexEnd < info.Games.GameCount-1 {
		games = append(games, info.Games.Games...)
		info, err = c.WebMatchHistory(accountID, platformID, info.Games.GameIndexEnd)
		if err != nil {
			return nil, err
		}
		var player string
		for _, sum := range info.Games.Games[0].ParticipantIdentities {
			if sum.Player.AccountID == accountID {
				player = sum.Player.SummonerName
				break
			}
		}
		if Debug {
			fmt.Fprintf(os.Stdout, "Len Games: %d IndexStart: %d IndexEnd: %d GamesCount: %d Player: %s\r", len(games), info.Games.GameIndexBegin, info.Games.GameIndexEnd, info.Games.GameCount, player)
		}
	}
	return games, nil
}

// GetAllGamesLimitPatch gets all the games for a player from current patch.
func GetAllGamesLimitPatch(accountID int64, platformID string, patch string, limitAmt int) (Games, error) {
	var games Games
	var info *GamesInfoWebUiResponse
	var err error
	info, err = c.WebMatchHistory(accountID, platformID, 0)
	if err != nil {
		switch platformID {
		case "NA":
			platformID = NA1
		case NA1:
			platformID = "NA"
		}
		info, err = c.WebMatchHistory(accountID, platformID, 0)
	}
	if err != nil {
		return nil, err
	}

	for info.Games.GameIndexEnd < info.Games.GameCount-1 {
		for _, game := range info.Games.Games {
			if !strings.HasPrefix(game.GameVersion, patch) || len(games) > limitAmt {
				return games, nil
			}
			games = append(games, game)
		}
		info, err = c.WebMatchHistory(accountID, platformID, info.Games.GameIndexEnd)
		if err != nil {
			return nil, err
		}
		var player string
		for _, sum := range info.Games.Games[0].ParticipantIdentities {
			if sum.Player.AccountID == accountID {
				player = sum.Player.SummonerName
				break
			}
		}
		if Debug {
			fmt.Fprintf(os.Stdout, "Len Games: %d IndexStart: %d IndexEnd: %d GamesCount: %d Player: %s\r", len(games), info.Games.GameIndexBegin, info.Games.GameIndexEnd, info.Games.GameCount, player)
		}
	}
	return games, nil
}
