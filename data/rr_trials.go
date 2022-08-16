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
		MinItemLevel:        61,
		SearchAlias:         []string{"究極幻想", "アルテマウェポン", "アルテマウエポン"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "極ガルーダ討滅戦",
		NameID:              "HE",
		ReleaseVersion:      enum.RealmReborn,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/7c17ae70cc6/",
		Persons:             8,
		MinItemLevel:        65,
		SearchAlias:         []string{"ガルーダ"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "極タイタン討滅戦",
		NameID:              "NA",
		ReleaseVersion:      enum.RealmReborn,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/589c727302e/",
		Persons:             8,
		MinItemLevel:        67,
		SearchAlias:         []string{"タイタン", "タコタン"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "極イフリート討滅戦",
		NameID:              "BOE",
		ReleaseVersion:      enum.RealmReborn,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/6af1a94ccca/",
		Persons:             8,
		MinItemLevel:        70,
		SearchAlias:         []string{"イフリート"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "極王モグル・モグXII世討滅戦",
		NameID:              "TH",
		ReleaseVersion:      enum.RealmReborn,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/b7c47c44490/",
		Persons:             8,
		MinItemLevel:        80,
		SearchAlias:         []string{"モグル・モグXII世", "モグル・モグ", "極王", "モグ王", "モグ", "モグル"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "極リヴァイアサン討滅戦",
		NameID:              "WHO",
		ReleaseVersion:      enum.RealmReborn,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/0850a8627aa/",
		Persons:             8,
		MinItemLevel:        80,
		SearchAlias:         []string{"リヴァイアサン", "リバイアサン", "リヴアイアサン"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "極ラムウ討滅戦",
		NameID:              "ST",
		ReleaseVersion:      enum.RealmReborn,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/4d8cae741db/",
		Persons:             8,
		MinItemLevel:        85,
		SearchAlias:         []string{"ラムウ"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "極シヴァ討滅戦",
		NameID:              "AAA",
		ReleaseVersion:      enum.RealmReborn,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/5f786d57228/",
		Persons:             8,
		MinItemLevel:        95,
		SearchAlias:         []string{"シヴァ"},
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
