package data

import (
	"fmt"
	"time"

	"github.com/sinmetal/ff14cf"
	"github.com/sinmetal/ff14cf/enum"
)

func RealmRebornTrials() []*ff14cf.Content {
	var raws []*ff14cf.Content
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "究極幻想 アルテマウェポン破壊作戦",
		NameID:              "MBUB",
		ReleaseVersion:      enum.RealmReborn,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/0de44d2eccf/",
		Persons:             8,
		SearchAlias:         []string{"究極幻想", "アルテマウェポン", "アルテマウエポン"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})

	rets, err := cook(raws)
	if err != nil {
		panic(fmt.Errorf("failed ARealmRebornTrials :%w", err))
	}

	return rets
}
