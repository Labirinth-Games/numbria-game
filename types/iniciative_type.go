package types

type Iniciative int

var IniciativeType iniciativeEnum

type iniciativeEnum struct {
	None   Iniciative
	Enemy  Iniciative
	Player Iniciative
}

func init() {
	const (
		none Iniciative = iota + 1
		creature
		player
	)

	IniciativeType = iniciativeEnum{
		None:   none,
		Enemy:  creature,
		Player: player,
	}
}
