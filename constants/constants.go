package constants

import "time"

const (
	BR   RegionEndPoint = "br1.api.riotgames.com"
	EUNE RegionEndPoint = "eun1.api.riotgames.com"
	EUW  RegionEndPoint = "euw1.api.riotgames.com"
	JP   RegionEndPoint = "jp1.api.riotgames.com"
	KR   RegionEndPoint = "kr.api.riotgames.com"
	LAN  RegionEndPoint = "la1.api.riotgames.com"
	LAS  RegionEndPoint = "la2.api.riotgames.com"
	NA   RegionEndPoint = "na1.api.riotgames.com"
	OCE  RegionEndPoint = "oc1.api.riotgames.com"
	TR   RegionEndPoint = "tr1.api.riotgames.com"
	RU   RegionEndPoint = "ru.api.riotgames.com"
	PBE  RegionEndPoint = "pbe1.api.riotgames.com"
)
const (
	WEEK = time.Hour * 24 * 7
	DAY  = time.Hour * 24
)

type RegionEndPoint string
