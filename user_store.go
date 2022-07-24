package ff14cf

import "time"

type User struct {
	// UserID is LodeStoneCharacterIDでいいか？
	// https://jp.finalfantasyxiv.com/lodestone/character/%d/
	UserID string

	// ViewID is URLとして利用する値
	// FirstName-LastName-適当な4文字
	ViewID string

	// FirstName is Game内FirstName
	FirstName string

	// LastName is Game内LastName
	LastName string

	// HomeWorld is Game内HomeWorld
	HomeWorld string

	// OpenedContents is 開放済みコンテンツ一覧
	OpenedContents []string

	// CompletedContents is クリア済みコンテンツ一覧
	CompletedContents []string

	CreatedAt time.Time
}
