package types

type ZoneType string

var ZoneEnum zoneTypeEnum

type zoneTypeEnum struct {
	Florest  ZoneType
	Clearing ZoneType
	River    ZoneType
	Boss     ZoneType
	Cave     ZoneType
	Way      ZoneType
	Wall     ZoneType
	Room     ZoneType
	End_Map  ZoneType
}

func Int() {
	ZoneEnum = zoneTypeEnum{
		Florest:  "floresta",
		Clearing: "clareira",
		River:    "rio",
		Boss:     "chefão",
		Cave:     "caverna",
		Wall:     "parede",
		Way:      "Caminho",
		Room:     "salão",
		End_Map:  "fim do mapa",
	}
}
