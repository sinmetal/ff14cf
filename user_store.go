package ff14cf

import "time"

type User struct {
	// UserID is UUIDでいいか？
	UserID string

	// FirstName is Game内FirstName
	FirstName string

	// LastName is Game内LastName
	LastName string

	// HomeWorld is Game内HomeWorld
	HomeWorld string

	// Salt is
	Salt string

	// OpenedContents is 開放済みコンテンツ一覧
	OpenedContents []string

	CreatedAt time.Time
}
