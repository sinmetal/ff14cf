package ff14cf

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sinmetal/ff14cf/errpkg"
)

type ContentsQuizHandler struct {
	contentStore *ContentStore
}

func NewContentsQuizHandler(ctx context.Context, contentStore *ContentStore) (*ContentsQuizHandler, error) {
	return &ContentsQuizHandler{
		contentStore: contentStore,
	}, nil
}

type ContentsQuizReq struct {
	ContentID       string `json:"contentID"`
	ContentQuizName string `json:"contentQuizName"`
}

func (h *ContentsQuizHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := &ContentsQuizReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v, err := h.contentStore.Get(ctx, req.ContentID)
	if err == errpkg.ErrNotFound {
		fmt.Printf("ContentID(%s) is not found\n", req.ContentID)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("failed ContentStore.Get : %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, name := range v.SearchAlias {
		if name == req.ContentQuizName {
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(v); err != nil {
				fmt.Printf("failed json.NewEncoder : %s\n", err)
			}
			return
		}
	}

	fmt.Printf("ContentID(%s), ContentQuizName(%s) is not found\n", req.ContentID, req.ContentQuizName)
	w.WriteHeader(http.StatusNotFound)
	return
}
