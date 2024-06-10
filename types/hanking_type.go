package types

type HankingType string

var HankingTypeEnum hankingTypeEnum

type hankingTypeEnum struct {
	S HankingType
	A HankingType
	B HankingType
	C HankingType
}

func init() {
	HankingTypeEnum = hankingTypeEnum{
		S: "S",
		A: "A",
		B: "B",
		C: "C",
	}
}

func (h hankingTypeEnum) ToInt(hank HankingType) int {
	switch hank {
	case HankingTypeEnum.S:
		return 8
	case HankingTypeEnum.A:
		return 6
	case HankingTypeEnum.B:
		return 2
	case HankingTypeEnum.C:
		return 0
	}

	return 0
}
