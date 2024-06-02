package utils

import "math/rand/v2"

type Dice int

var Dices DiceModel

type DiceModel struct {
	D4  Dice
	D6  Dice
	D8  Dice
	D10 Dice
	D12 Dice
	D20 Dice
}

func init() {
	Dices = DiceModel{
		D4:  3,
		D6:  5,
		D8:  7,
		D10: 9,
		D12: 11,
		D20: 19,
	}
}

func ConvertToDiceEnum(dice string) Dice {
	switch dice {
	case "d4":
		return Dices.D4
	case "d6":
		return Dices.D6
	case "d8":
		return Dices.D8
	case "d10":
		return Dices.D10
	case "d12":
		return Dices.D12
	case "d20":
		return Dices.D20
	}

	return 0
}

func RollDice(dice Dice) int {
	return rand.IntN(int(dice)) + 1
}

func TestPrecision(accuracy int) bool {
	return RollDice(Dices.D8) > accuracy
}
