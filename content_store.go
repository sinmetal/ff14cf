package ff14cf

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/sinmetal/ff14cf/enum"
	"github.com/sinmetal/ff14cf/errpkg"
	"google.golang.org/api/iterator"
)

type Content struct {
	// ContentID is URLで表示するID
	// $ReleaseVolume-Kind-Grade-英語版コンテンツの名前を適当に3文字ぐらいにしたもの
	// e.g. 極ガルーダ討滅戦 RR-TR-EX-HE
	ContentID string `firestore:"-" json:"contentID"`

	// Name is コンテンツ名
	Name string `json:"name"`

	// NameID is IDに使うName
	// 英語版コンテンツの名前を適当に3文字ぐらいにしたもの
	// https://na.finalfantasyxiv.com/lodestone/playguide/db/duty/
	NameID string `json:"nameID"`

	// ReleaseVersion is 新生, 蒼天など
	ReleaseVersion enum.ReleaseVersion `json:"releaseVersion"`

	// ContentKind is 討滅戦, アライアンスレイドなど
	ContentKind enum.ContentKind `json:"contentKind"`

	// ContentDifficulty is 極, 絶などのグレード
	// アルテマウェポンなどは極もあるが、その上の絶もあるので、それをフィルタリングするために使う
	ContentDifficulty enum.ContentDifficulty `json:"contentDifficulty"`

	// エオルゼアデータベースURL JP
	EorzeaDatabaseJPURL string `json:"eorzeaDatabaseJPURL"`

	// Persons is 参加人数
	// アライアンスレイドは24人なのか、8人なのか・・・
	Persons int `json:"persons"`

	// MinItemLevel 最低平均アイテムレベル
	MinItemLevel int `json:"minItemLevel"`

	// SearchAlias is 検索に利用するコンテンツ名の別名
	// e.g. 極ガルーダ討滅戦 -> ガルーダ
	SearchAlias []string `json:"searchAlias"`

	// OpenedUsers is 開放済みのUserList
	OpenedUsers []string `json:"openedUsers"`

	// CompletedUsers is LevelSyncでクリアしているUserList
	CompletedUsers []string `json:"completedUsers"`

	CreatedAt time.Time `json:"createdAt"`
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

func (s *ContentStore) Get(ctx context.Context, contentID string) (*Content, error) {
	v, err := s.CollectionRef().Doc(contentID).Get(ctx)
	if err != nil {
		return nil, err
	}
	var content *Content
	if err := v.DataTo(&content); err != nil {
		return nil, err
	}
	content.ContentID = contentID
	return content, nil
}

func (s *ContentStore) GetBySearchAlias(ctx context.Context, contentName string) (*Content, error) {
	iter := s.CollectionRef().Where("SearchAlias", "array-contains", contentName).Documents(ctx)
	defer iter.Stop()
	for {
		v, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var content *Content
		if err := v.DataTo(&content); err != nil {
			return nil, err
		}
		content.ContentID = v.Ref.ID
		return content, nil
	}

	return nil, errpkg.ErrNotFound
}
