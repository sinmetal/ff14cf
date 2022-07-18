package ff14cf

import (
	"time"

	"github.com/sinmetal/ff14cf/enum"
)

type Content struct {
	// ID is URLで表示するID
	// $ReleaseVolume-Kind-Grade-英語版コンテンツの名前を適当に3文字ぐらいにしたもの
	// e.g. 極ガルーダ討滅戦 RR-TR-EX-HE
	ID string

	// Name is コンテンツ名
	Name string

	// NameID is IDに使うName
	// 英語版コンテンツの名前を適当に3文字ぐらいにしたもの
	// https://na.finalfantasyxiv.com/lodestone/playguide/db/duty/
	NameID string

	// ReleaseVersion is 新生, 蒼天など
	ReleaseVersion enum.ReleaseVersion

	// ContentKind is 討滅戦, アライアンスレイドなど
	ContentKind enum.ContentKind

	// ContentDifficulty is 極, 絶などのグレード
	// アルテマウェポンなどは極もあるが、その上の絶もあるので、それをフィルタリングするために使う
	ContentDifficulty enum.ContentDifficulty

	// Persons is 参加人数
	// アライアンスレイドは24人なのか、8人なのか・・・
	Persons int

	// SearchAlias is 検索に利用するコンテンツ名の別名
	// e.g. 極ガルーダ討滅戦 -> ガルーダ
	SearchAlias []string

	// OpenedUsers is 開放済みのUserList
	OpenedUsers []string

	// CompletedUsers is LevelSyncでクリアしているUserList
	CompletedUsers []string

	CreatedAt time.Time
}
