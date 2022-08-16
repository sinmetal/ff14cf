package ff14cf

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/sinmetal/ff14cf/enum"
)

type Content struct {
	// ContentID is URLで表示するID
	// $ReleaseVolume-Kind-Grade-英語版コンテンツの名前を適当に3文字ぐらいにしたもの
	// e.g. 極ガルーダ討滅戦 RR-TR-EX-HE
	ContentID string `firestore:"-"`

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

	// エオルゼアデータベースURL JP
	EorzeaDatabaseJPURL string

	// Persons is 参加人数
	// アライアンスレイドは24人なのか、8人なのか・・・
	Persons int

	// MinItemLevel 最低平均アイテムレベル
	MinItemLevel int

	// SearchAlias is 検索に利用するコンテンツ名の別名
	// e.g. 極ガルーダ討滅戦 -> ガルーダ
	SearchAlias []string

	// OpenedUsers is 開放済みのUserList
	OpenedUsers []string

	// CompletedUsers is LevelSyncでクリアしているUserList
	CompletedUsers []string

	CreatedAt time.Time
}

func (c *Content) NewContentID() string {
	return fmt.Sprintf("%s-%s-%s-%s", c.ReleaseVersion.ShortID(), c.ContentKind.ShortID(), c.ContentDifficulty.ShortID(), c.NameID)
}

type ContentStore struct {
	fs *firestore.Client
}

func NewContentStore(ctx context.Context, fs *firestore.Client) (*ContentStore, error) {
	return &ContentStore{
		fs: fs,
	}, nil
}

func (s *ContentStore) CollectionName() string {
	return "Contents"
}

func (s *ContentStore) CollectionRef() *firestore.CollectionRef {
	return s.fs.Collection(s.CollectionName())
}

func (s *ContentStore) Create(ctx context.Context, value *Content) (*Content, error) {
	wr, err := s.CollectionRef().Doc(value.ContentID).Create(ctx, value)
	if err != nil {
		return nil, err
	}
	value.CreatedAt = wr.UpdateTime
	return value, nil
}
