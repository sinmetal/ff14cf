package ff14cf

import (
	"context"
	"errors"
	"testing"
)

func TestGetFF14CharacterNameFromLodeStone(t *testing.T) {
	ctx := context.Background()

	fn, ln, err := GetFF14CharacterNameFromLodeStone(ctx, "41165069")
	if err != nil {
		t.Fatal(err)
	}
	if e, g := "Sinmetal", fn; e != g {
		t.Errorf("firstName want %s but got %s\n", e, g)
	}
	if e, g := "Iron", ln; e != g {
		t.Errorf("lastName want %s but got %s\n", e, g)
	}

	// invalid characterID
	{
		_, _, err := GetFF14CharacterNameFromLodeStone(ctx, "9999999")
		if !errors.Is(err, ErrLodeStoneCharacterNotFound) {
			t.Errorf("want ErrLodeStoneInvalidFormat: %s", err)
		}
	}
}
