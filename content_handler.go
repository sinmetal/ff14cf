package ff14cf

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// ContentsHandler is FF14 Contentsを表示するハンドラ
type ContentsHandler struct {
	contentStore          *ContentStore
	contentDetailTemplate *template.Template
}

func NewContentsHandler(ctx context.Context, contentStore *ContentStore) (*ContentsHandler, error) {
	contentDetailTemplate, err := template.ParseFiles("./static/content.html")
	if err != nil {
		return nil, fmt.Errorf("failed ContentDetailTemplate Parse: %w", err)
	}

	return &ContentsHandler{
		contentStore:          contentStore,
		contentDetailTemplate: contentDetailTemplate,
	}, nil
}

func (h *ContentsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := strings.TrimPrefix(r.RequestURI, "/content/")

	v, err := h.contentStore.Get(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Printf("failed ContentStore.Get :%s\n", err)
		return
	}

	d := map[string]interface{}{
		"ContentID":         id,
		"ContentName":       v.Name,
		"MinItemLevel":      v.MinItemLevel,
		"ContentKind":       v.ContentKind.NameJP(),
		"ContentDifficulty": v.ContentDifficulty.NameJP(),
		"ReleaseVersion":    v.ReleaseVersion.NameJP(),
	}
	if err := h.contentDetailTemplate.Execute(w, d); err != nil {
		fmt.Printf("failed ContentTemplate.Execute %s\n", err)
	}
}
