package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/sinmetal/ff14cf"
	"github.com/sinmetal/ff14cf/data"
	metadatabox "github.com/sinmetalcraft/gcpbox/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	ctx := context.Background()
	pID, err := metadatabox.ProjectID()
	if err != nil {
		panic(err)
	}

	fs, err := firestore.NewClient(ctx, pID)
	if err != nil {
		panic(err)
	}
	defer fs.Close()

	contentStore, err := ff14cf.NewContentStore(ctx, fs)
	if err != nil {
		panic(err)
	}

	launcher := &Launcher{
		contentStore,
	}
	if err := launcher.RealmRebornTrials(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Finish")
}

type Launcher struct {
	ContentStore *ff14cf.ContentStore
}

func (l *Launcher) RealmRebornTrials(ctx context.Context) error {
	list := data.RealmRebornTrials()
	for i, v := range list {
		_, err := l.ContentStore.Create(ctx, v)
		if status.Code(err) == codes.AlreadyExists {
			fmt.Printf("AlreadyExists %d:%s:%s\n", i, v.ContentID, v.Name)
			continue
		} else if err != nil {
			return fmt.Errorf("failed create %d:%s:%s :%w", i, v.ContentID, v.Name, err)
		}
		fmt.Printf("Create %d:%s:%s\n", i, v.ContentID, v.Name)
	}
	return nil
}
