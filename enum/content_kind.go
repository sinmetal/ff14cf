package enum

type ContentKind int

const (
	NoneContentKind ContentKind = iota
	Dungeon
	Trial
	AllianceRaid
	NormalRaid
)

func (ck ContentKind) NameJP() string {
	switch ck {
	case Dungeon:
		return "ダンジョン"
	case Trial:
		return "討滅戦"
	case AllianceRaid:
		return "アライアンスレイド"
	case NormalRaid:
		return "ノーマルレイド"
	default:
		return "不明のコンテンツの種類"
	}
}

func (ck ContentKind) ShortID() string {
	switch ck {
	case Dungeon:
		return "DU"
	case Trial:
		return "TR"
	case AllianceRaid:
		return "AR"
	case NormalRaid:
		return "NR"
	default:
		return "XX"
	}
}
