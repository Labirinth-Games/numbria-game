package types

type BattleStatus string

var BattleStatusType battleStatusEnum

type battleStatusEnum struct {
	Sucess BattleStatus
	Fail   BattleStatus
}

func init() {
	BattleStatusType = battleStatusEnum{
		Sucess: "sucess",
		Fail:   "fail",
	}
}
