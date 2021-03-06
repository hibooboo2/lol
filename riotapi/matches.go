package riotapi

type Matchlist struct {
	Matches    []MatchReference `json:"matches"`
	StartIndex int              `json:"startIndex"`
	EndIndex   int              `json:"endIndex"`
	TotalGames int              `json:"totalGames"`
}

type MatchReference struct {
	PlatformID string `json:"platformId"`
	GameID     int64  `json:"gameId"`
	Champion   int    `json:"champion"`
	Queue      int    `json:"queue"`
	Season     int    `json:"season"`
	Timestamp  int64  `json:"timestamp"`
	Role       string `json:"role"`
	Lane       string `json:"lane"`
}

type Match struct {
	GameCreation          int64  `json:"gameCreation"`
	GameDuration          int64  `json:"gameDuration"`
	GameID                int64  `json:"gameId"`
	GameMode              string `json:"gameMode"`
	GameType              string `json:"gameType"`
	GameVersion           string `json:"gameVersion"`
	MapID                 int64  `json:"mapId"`
	ParticipantIdentities []struct {
		ParticipantID int64 `json:"participantId"`
		Player        struct {
			AccountID         int64  `json:"accountId"`
			CurrentAccountID  int64  `json:"currentAccountId"`
			CurrentPlatformID string `json:"currentPlatformId"`
			MatchHistoryURI   string `json:"matchHistoryUri"`
			PlatformID        string `json:"platformId"`
			ProfileIcon       int64  `json:"profileIcon"`
			SummonerID        int64  `json:"summonerId"`
			SummonerName      string `json:"summonerName"`
		} `json:"player"`
	} `json:"participantIdentities"`
	Participants []struct {
		ChampionID                int64  `json:"championId"`
		HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
		ParticipantID             int64  `json:"participantId"`
		Spell1Id                  int64  `json:"spell1Id"`
		Spell2Id                  int64  `json:"spell2Id"`
		Stats                     struct {
			Assists                        int64 `json:"assists"`
			ChampLevel                     int64 `json:"champLevel"`
			CombatPlayerScore              int64 `json:"combatPlayerScore"`
			DamageDealtToObjectives        int64 `json:"damageDealtToObjectives"`
			DamageDealtToTurrets           int64 `json:"damageDealtToTurrets"`
			DamageSelfMitigated            int64 `json:"damageSelfMitigated"`
			Deaths                         int64 `json:"deaths"`
			DoubleKills                    int64 `json:"doubleKills"`
			FirstBloodAssist               bool  `json:"firstBloodAssist"`
			FirstBloodKill                 bool  `json:"firstBloodKill"`
			FirstInhibitorAssist           bool  `json:"firstInhibitorAssist"`
			FirstInhibitorKill             bool  `json:"firstInhibitorKill"`
			FirstTowerAssist               bool  `json:"firstTowerAssist"`
			FirstTowerKill                 bool  `json:"firstTowerKill"`
			GoldEarned                     int64 `json:"goldEarned"`
			GoldSpent                      int64 `json:"goldSpent"`
			InhibitorKills                 int64 `json:"inhibitorKills"`
			Item0                          int64 `json:"item0"`
			Item1                          int64 `json:"item1"`
			Item2                          int64 `json:"item2"`
			Item3                          int64 `json:"item3"`
			Item4                          int64 `json:"item4"`
			Item5                          int64 `json:"item5"`
			Item6                          int64 `json:"item6"`
			KillingSprees                  int64 `json:"killingSprees"`
			Kills                          int64 `json:"kills"`
			LargestCriticalStrike          int64 `json:"largestCriticalStrike"`
			LargestKillingSpree            int64 `json:"largestKillingSpree"`
			LargestMultiKill               int64 `json:"largestMultiKill"`
			LongestTimeSpentLiving         int64 `json:"longestTimeSpentLiving"`
			MagicDamageDealt               int64 `json:"magicDamageDealt"`
			MagicDamageDealtToChampions    int64 `json:"magicDamageDealtToChampions"`
			MagicalDamageTaken             int64 `json:"magicalDamageTaken"`
			NeutralMinionsKilled           int64 `json:"neutralMinionsKilled"`
			ObjectivePlayerScore           int64 `json:"objectivePlayerScore"`
			ParticipantID                  int64 `json:"participantId"`
			PentaKills                     int64 `json:"pentaKills"`
			Perk0                          int64 `json:"perk0"`
			Perk0Var1                      int64 `json:"perk0Var1"`
			Perk0Var2                      int64 `json:"perk0Var2"`
			Perk0Var3                      int64 `json:"perk0Var3"`
			Perk1                          int64 `json:"perk1"`
			Perk1Var1                      int64 `json:"perk1Var1"`
			Perk1Var2                      int64 `json:"perk1Var2"`
			Perk1Var3                      int64 `json:"perk1Var3"`
			Perk2                          int64 `json:"perk2"`
			Perk2Var1                      int64 `json:"perk2Var1"`
			Perk2Var2                      int64 `json:"perk2Var2"`
			Perk2Var3                      int64 `json:"perk2Var3"`
			Perk3                          int64 `json:"perk3"`
			Perk3Var1                      int64 `json:"perk3Var1"`
			Perk3Var2                      int64 `json:"perk3Var2"`
			Perk3Var3                      int64 `json:"perk3Var3"`
			Perk4                          int64 `json:"perk4"`
			Perk4Var1                      int64 `json:"perk4Var1"`
			Perk4Var2                      int64 `json:"perk4Var2"`
			Perk4Var3                      int64 `json:"perk4Var3"`
			Perk5                          int64 `json:"perk5"`
			Perk5Var1                      int64 `json:"perk5Var1"`
			Perk5Var2                      int64 `json:"perk5Var2"`
			Perk5Var3                      int64 `json:"perk5Var3"`
			PerkPrimaryStyle               int64 `json:"perkPrimaryStyle"`
			PerkSubStyle                   int64 `json:"perkSubStyle"`
			PhysicalDamageDealt            int64 `json:"physicalDamageDealt"`
			PhysicalDamageDealtToChampions int64 `json:"physicalDamageDealtToChampions"`
			PhysicalDamageTaken            int64 `json:"physicalDamageTaken"`
			PlayerScore0                   int64 `json:"playerScore0"`
			PlayerScore1                   int64 `json:"playerScore1"`
			PlayerScore2                   int64 `json:"playerScore2"`
			PlayerScore3                   int64 `json:"playerScore3"`
			PlayerScore4                   int64 `json:"playerScore4"`
			PlayerScore5                   int64 `json:"playerScore5"`
			PlayerScore6                   int64 `json:"playerScore6"`
			PlayerScore7                   int64 `json:"playerScore7"`
			PlayerScore8                   int64 `json:"playerScore8"`
			PlayerScore9                   int64 `json:"playerScore9"`
			QuadraKills                    int64 `json:"quadraKills"`
			SightWardsBoughtInGame         int64 `json:"sightWardsBoughtInGame"`
			TimeCCingOthers                int64 `json:"timeCCingOthers"`
			TotalDamageDealt               int64 `json:"totalDamageDealt"`
			TotalDamageDealtToChampions    int64 `json:"totalDamageDealtToChampions"`
			TotalDamageTaken               int64 `json:"totalDamageTaken"`
			TotalHeal                      int64 `json:"totalHeal"`
			TotalMinionsKilled             int64 `json:"totalMinionsKilled"`
			TotalPlayerScore               int64 `json:"totalPlayerScore"`
			TotalScoreRank                 int64 `json:"totalScoreRank"`
			TotalTimeCrowdControlDealt     int64 `json:"totalTimeCrowdControlDealt"`
			TotalUnitsHealed               int64 `json:"totalUnitsHealed"`
			TripleKills                    int64 `json:"tripleKills"`
			TrueDamageDealt                int64 `json:"trueDamageDealt"`
			TrueDamageDealtToChampions     int64 `json:"trueDamageDealtToChampions"`
			TrueDamageTaken                int64 `json:"trueDamageTaken"`
			TurretKills                    int64 `json:"turretKills"`
			UnrealKills                    int64 `json:"unrealKills"`
			VisionScore                    int64 `json:"visionScore"`
			VisionWardsBoughtInGame        int64 `json:"visionWardsBoughtInGame"`
			Win                            bool  `json:"win"`
		} `json:"stats"`
		TeamID   int64 `json:"teamId"`
		Timeline struct {
			CreepsPerMinDeltas          map[string]float64 `json:"creepsPerMinDeltas"`
			CsDiffPerMinDeltas          map[string]float64 `json:"csDiffPerMinDeltas"`
			DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
			DamageTakenPerMinDeltas     map[string]float64 `json:"damageTakenPerMinDeltas"`
			GoldPerMinDeltas            map[string]float64 `json:"goldPerMinDeltas"`
			Lane                        string             `json:"lane"`
			ParticipantID               int64              `json:"participantId"`
			Role                        string             `json:"role"`
			XpDiffPerMinDeltas          map[string]float64 `json:"xpDiffPerMinDeltas"`
			XpPerMinDeltas              map[string]float64 `json:"xpPerMinDeltas"`
		} `json:"timeline"`
	} `json:"participants"`
	PlatformID string `json:"platformId"`
	QueueID    int64  `json:"queueId"`
	SeasonID   int64  `json:"seasonId"`
	Teams      []struct {
		Bans                 []interface{} `json:"bans"`
		BaronKills           int64         `json:"baronKills"`
		DominionVictoryScore int64         `json:"dominionVictoryScore"`
		DragonKills          int64         `json:"dragonKills"`
		FirstBaron           bool          `json:"firstBaron"`
		FirstBlood           bool          `json:"firstBlood"`
		FirstDragon          bool          `json:"firstDragon"`
		FirstInhibitor       bool          `json:"firstInhibitor"`
		FirstRiftHerald      bool          `json:"firstRiftHerald"`
		FirstTower           bool          `json:"firstTower"`
		InhibitorKills       int64         `json:"inhibitorKills"`
		RiftHeraldKills      int64         `json:"riftHeraldKills"`
		TeamID               int64         `json:"teamId"`
		TowerKills           int64         `json:"towerKills"`
		VilemawKills         int64         `json:"vilemawKills"`
		Win                  string        `json:"win"`
	} `json:"teams"`
}
