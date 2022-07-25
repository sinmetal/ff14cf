package ff14cf

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterHandler struct {
	userStore *UserStore
}

func NewRegisterHandler(ctx context.Context, userStore *UserStore) (*RegisterHandler, error) {
	return &RegisterHandler{
		userStore: userStore,
	}, nil
}

type RegisterReq struct {
	CharacterID string `json:"characterID"`
}

type RegisterErrorResponse struct {
	Message string `json:"message"`
}

func (h *RegisterHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := &RegisterReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 数字Onlyであるかどうかも見た方がいいが、とりあえずめんどうなので
	if len(req.CharacterID) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fn, ln, err := GetFF14CharacterNameFromLodeStone(ctx, req.CharacterID)
	if err != nil {
		if errors.Is(err, ErrLodeStoneCharacterNotFound) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)
			err := json.NewEncoder(w).Encode(&RegisterErrorResponse{
				Message: "",
			})
			if err != nil {
				// noop
			}
		}
		if errors.Is(err, ErrLodeStoneInvalidFormat) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(w).Encode(&RegisterErrorResponse{
				Message: "",
			})
			if err != nil {
				// noop
			}
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(&RegisterErrorResponse{
			Message: "",
		})
		if err != nil {
			// noop
		}
	}

	viewID, err := h.userStore.NewViewID(ctx, fn, ln)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(&RegisterErrorResponse{
			Message: "",
		})
		if err != nil {
			// noop
		}
	}
	created, err := h.userStore.Create(ctx, &User{
		UserID:            req.CharacterID,
		ViewID:            viewID,
		FirstName:         fn,
		LastName:          ln,
		HomeWorld:         "",
		OpenedContents:    []string{},
		CompletedContents: []string{},
	})
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			w.WriteHeader(http.StatusConflict)
			err := json.NewEncoder(w).Encode(&RegisterErrorResponse{
				Message: "",
			})
			if err != nil {
				// noop
			}
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(&RegisterErrorResponse{
			Message: "",
		})
		if err != nil {
			// noop
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(created); err != nil {
		fmt.Printf("failed response write. err=%s\n", err)
	}
}
