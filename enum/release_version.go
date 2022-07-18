package enum

type ReleaseVersion int

const (
	NoneReleaseVersion ReleaseVersion = iota
	RealmReborn
	Heavensward
	Stormblood
	Shadowbringers
	Endwalker
)

func (rv ReleaseVersion) NameJP() string {
	switch rv {
	case RealmReborn:
		return "新生"
	case Heavensward:
		return "蒼天"
	case Stormblood:
		return "紅蓮"
	case Shadowbringers:
		return "漆黒"
	case Endwalker:
		return "暁月"
	default:
		return "不明のリリースバージョン"
	}
}

func (rv ReleaseVersion) ShortID() string {
	switch rv {
	case RealmReborn:
		return "RR"
	case Heavensward:
		return "HW"
	case Stormblood:
		return "SB"
	case Shadowbringers:
		return "SB"
	case Endwalker:
		return "EW"
	default:
		return "XX"
	}
}
