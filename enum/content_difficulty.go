package enum

type ContentDifficulty int

const (
	NoneContentDifficulty ContentDifficulty = iota
	Normal
	Hard
	Extreme
	Savage
	Ultimate
)

func (cd ContentDifficulty) NameJP() string {
	switch cd {
	case Normal:
		return "ノーマル"
	case Hard:
		return "真"
	case Extreme:
		return "極"
	case Savage:
		return "零式"
	case Ultimate:
		return "絶"
	default:
		return "不明のコンテンツ難易度"
	}
}

func (cd ContentDifficulty) ShortID() string {
	switch cd {
	case Normal:
		return "NR"
	case Hard:
		return "HA"
	case Extreme:
		return "EX"
	case Savage:
		return "SV"
	case Ultimate:
		return "UT"
	default:
		return "XX"
	}
}
