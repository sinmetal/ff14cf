package ff14cf

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
)

type User struct {
	// UserID is LodeStoneCharacterIDでいいか？
	// https://jp.finalfantasyxiv.com/lodestone/character/%d/
	UserID string `firestore:"-" json:"userID"`

	// ViewID is URLとして利用する値
	// FirstName-LastName-適当な4文字
	ViewID string `json:"viewID"`

	// FirstName is Game内FirstName
	FirstName string `json:"firstName"`

	// LastName is Game内LastName
	LastName string `json:"lastName"`

	// HomeWorld is Game内HomeWorld
	HomeWorld string `json:"homeWorld"`

	// OpenedContents is 開放済みコンテンツ一覧
	OpenedContents []string `json:"openedContents"`

	// CompletedContents is クリア済みコンテンツ一覧
	CompletedContents []string `json:"completedContents"`

	CreatedAt time.Time `firestore:",serverTimestamp" json:"createdAt"`
}

func NewUserStore(ctx context.Context, fs *firestore.Client) (*UserStore, error) {
	return &UserStore{
		fs: fs,
	}, nil
}

type UserStore struct {
	fs *firestore.Client
}

func (s *UserStore) CollectionName() string {
	return "Users"
}

func (s *UserStore) CollectionRef() *firestore.CollectionRef {
	return s.fs.Collection(s.CollectionName())
}

func (s *UserStore) NewViewID(ctx context.Context, firstName string, lastName string) (string, error) {
	v, err := MakeRandomStr(4)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s-%s-%s", firstName, lastName, v), nil
}

func (s *UserStore) Create(ctx context.Context, user *User) (*User, error) {
	wr, err := s.CollectionRef().Doc(user.UserID).Create(ctx, user)
	if err != nil {
		return nil, err
	}
	user.CreatedAt = wr.UpdateTime
	return user, nil
}
