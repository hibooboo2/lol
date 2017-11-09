package lol

import mgo "gopkg.in/mgo.v2"

func (db *lolMongo) transferToAnother(host string, port int) error {
	db2, err := NewLolMongo(host, port)
	if err != nil {
		return err
	}
	logger.Println("Starting transfer")
	batchSize := 100
	totalGames, _ := db.games.Find(nil).Count()
	var count int
	for count < totalGames {
		var games []Game
		err = db.games.Find(nil).Skip(count).Limit(batchSize).All(&games)
		if err != nil {
			logger.Println("err:", err)
			return err
		}
		count += len(games)
		logger.Println("Got batch", count)
		b := db2.games.Bulk()
		for _, game := range games {
			b.Insert(game)
		}
		res, err := b.Run()
		if err != nil {
			logger.Println("err:", err)
			return err
		}
		logger.Println("Inserted batch", count)
		if res.Matched+res.Modified != len(games) {
			logger.Printf("May have skipped games. Games: %d Matched+Mod: %d", len(games), res.Matched+res.Modified)
		}
		games = nil
	}
	logger.Println("Moved ", count, "Games")

	var players []Player
	totalPlayers, _ := db.players.Find(nil).Count()
	count = 0
	for count < totalPlayers {
		err = db.players.Find(nil).Skip(count).Batch(batchSize).All(&players)
		if err != nil {
			logger.Println("err:", err)
			return err
		}
		count += len(players)
		b := db2.players.Bulk()
		for _, player := range players {
			b.Insert(player)
		}
		res, err := b.Run()
		if err != nil {
			logger.Println("err:", err)
			return err
		}
		if res.Matched+res.Modified != len(players) {
			logger.Printf("May have skipped players. Players: %d Matched+Mod: %d", len(players), res.Matched+res.Modified)
		}
		players = nil
	}

	totalPlayers, _ = db.playersVisited.Find(nil).Count()
	count = 0
	for count < totalPlayers {
		err = db.playersVisited.Find(nil).Skip(count).Batch(batchSize).All(&players)
		if err != nil {
			logger.Println("err:", err)
			return err
		}
		count += len(players)
		b := db2.playersVisited.Bulk()
		for _, player := range players {
			b.Insert(player)
		}
		res, err := b.Run()
		if err != nil {
			logger.Println("err:", err)
			return err
		}
		if res.Matched+res.Modified != len(players) {
			logger.Printf("May have skipped players. Players: %d Matched+Mod: %d", len(players), res.Matched+res.Modified)
		}
		players = nil
	}

	return nil
}

func (db *lolMongo) EnsureIndexes() error {
	db.requests.EnsureIndex(mgo.Index{
		Key:      []string{"url"},
		DropDups: true,
		Unique:   true,
	})

	err := db.games.EnsureIndex(mgo.Index{
		Key:      []string{"gameid"},
		DropDups: true,
		Unique:   true,
	})
	if err != nil {
		return err
	}
	return db.players.EnsureIndex(mgo.Index{
		Key:      []string{"accountid"},
		DropDups: true,
		Unique:   true,
	})
}
