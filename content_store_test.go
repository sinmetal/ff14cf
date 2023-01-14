package ff14cf

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"
	metadatabox "github.com/sinmetalcraft/gcpbox/metadata"
)

func TestContentStore_GetBySearchAlias(t *testing.T) {
	ctx := context.Background()

	contentStore := newTestContentStore(t)

	got, err := contentStore.GetBySearchAlias(ctx, "タイタン")
	if err != nil {
		t.Fatal(err)
	}
	if e, g := "RR-TR-EX-NA", got.ContentID; e != g {
		t.Errorf("ContentID want %s but got %s", e, g)
	}
}

func newTestContentStore(t *testing.T) *ContentStore {
	ctx := context.Background()

	pID, err := metadatabox.ProjectID()
	if err != nil {
		t.Fatal(err)
	}
	fs, err := firestore.NewClient(ctx, pID)
	if err != nil {
		t.Fatal(err)
	}

	store, err := NewContentStore(ctx, fs)
	if err != nil {
		t.Fatal(err)
	}
	return store
}
