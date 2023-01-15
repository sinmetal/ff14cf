package data

import (
	"fmt"
	"time"

	"github.com/sinmetal/ff14cf"
	"github.com/sinmetal/ff14cf/enum"
)

func HeavenswardTrials() []*ff14cf.Content {
	var raws []*ff14cf.Content
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "極ビスマルク討滅戦",
		NameID:              "LB", // The Limitless Blue
		ReleaseVersion:      enum.Heavensward,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/e9bb63551a4/",
		Persons:             8,
		MinItemLevel:        165,
		SearchAlias:         []string{"ビスマルク"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "極ラーヴァナ討滅戦",
		NameID:              "TT", // Thok ast Thok
		ReleaseVersion:      enum.Heavensward,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/83f028575c8/",
		Persons:             8,
		MinItemLevel:        175,
		SearchAlias:         []string{"ラーヴァナ"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "蒼天幻想 ナイツ・オブ・ラウンド討滅戦",
		NameID:              "MBTR", // The Minstrel's Ballad: Thordan's Reign
		ReleaseVersion:      enum.Heavensward,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/a8a4860068c/",
		Persons:             8,
		MinItemLevel:        190,
		SearchAlias:         []string{"ナイツ・オブ・ラウンド", "ナイツ", "ナイツオブラウンド", "蒼天幻想", "騎士王", "円卓の騎士", "トールダン"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "極ニーズヘッグ征竜戦",
		NameID:              "MBNR", // The Minstrel's Ballad: Nidhogg's Rage
		ReleaseVersion:      enum.Heavensward,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/0e880006330/",
		Persons:             8,
		MinItemLevel:        220,
		SearchAlias:         []string{"ニーズヘッグ", "征竜戦", "ニーズ", "邪竜", "邪龍"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "極魔神セフィロト討滅戦",
		NameID:              "S1T7", // Containment Bay S1T7
		ReleaseVersion:      enum.Heavensward,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/e05c982993d/",
		Persons:             8,
		MinItemLevel:        205,
		SearchAlias:         []string{"セフィロト", "魔神", "魔神セフィロト"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "極女神ソフィア討滅戦",
		NameID:              "P1T6", // Containment Bay P1T6
		ReleaseVersion:      enum.Heavensward,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/212ceb19a34/",
		Persons:             8,
		MinItemLevel:        235,
		SearchAlias:         []string{"ソフィア", "女神", "女神ソフィア"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})
	raws = append(raws, &ff14cf.Content{
		ContentID:           "",
		Name:                "極鬼神ズルワーン討滅戦",
		NameID:              "Z1T9", // Containment Bay Z1T9
		ReleaseVersion:      enum.Heavensward,
		ContentKind:         enum.Trial,
		ContentDifficulty:   enum.Extreme,
		EorzeaDatabaseJPURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/duty/26a86785be9/",
		Persons:             8,
		MinItemLevel:        250,
		SearchAlias:         []string{"ズルワーン", "鬼神", "鬼神ズルワーン"},
		OpenedUsers:         []string{},
		CompletedUsers:      []string{},
		CreatedAt:           time.Now(),
	})

	rets, err := cook(raws)
	if err != nil {
		panic(fmt.Errorf("failed HeavenswardTrials :%w", err))
	}

	return rets
}
