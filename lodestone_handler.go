package ff14cf

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type LoadStoneHandler struct {
}

func NewLoadStoneHandler(ctx context.Context) (*LoadStoneHandler, error) {
	return &LoadStoneHandler{}, nil
}

type LoadStoneGetResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (h *LoadStoneHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.FormValue("characterID")
	if len(id) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fn, ln, err := GetFF14CharacterNameFromLodeStone(ctx, id)
	if errors.Is(err, ErrLodeStoneCharacterNotFound) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Printf("ErrLodeStoneCharacterNotFound.characterID=%s :%s\n", id, err)
		return
	}
	if errors.Is(err, ErrLodeStoneInvalidFormat) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("ErrLodeStoneInvalidFormat.characterID=%s :%s\n", id, err)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("failed get lodestone.characterID=%s :%s\n", id, err)
		return
	}

	res := &LoadStoneGetResponse{
		FirstName: fn,
		LastName:  ln,
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		fmt.Printf("failed response write.err=%s", err)
	}
}
