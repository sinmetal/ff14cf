package data

import (
	"fmt"

	"github.com/sinmetal/ff14cf"
)

func cook(raws []*ff14cf.Content) ([]*ff14cf.Content, error) {
	m := map[string]bool{}
	var rets []*ff14cf.Content
	for _, raw := range raws {
		raw.ContentID = raw.NewContentID()
		_, ok := m[raw.ContentID]
		if ok {
			return nil, fmt.Errorf("ContentIDが衝突・・・！！！ %s", raw.ContentID)
		}
		m[raw.ContentID] = true
		rets = append(rets, raw)
	}
	return rets, nil
}
