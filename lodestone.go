package ff14cf

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dyatlov/go-opengraph/opengraph"
)

var ErrLodeStoneCharacterNotFound = fmt.Errorf("loadstone character notfound")
var ErrLodeStoneInvalidFormat = fmt.Errorf("loadstone response invalid format")

func GetFF14CharacterNameFromLodeStone(ctx context.Context, characterID string) (firstName string, lastName string, err error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", fmt.Sprintf("https://jp.finalfantasyxiv.com/lodestone/character/%s/", characterID), nil)
	if err != nil {
		return "", "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("failed LodeStone Get Character Name %s", err)
		}
	}()
	switch resp.StatusCode {
	case http.StatusOK:
		// noop
	case http.StatusNotFound:
		return "", "", ErrLodeStoneCharacterNotFound
	default:
		return "", "", fmt.Errorf("failed Get LodeStone")
	}

	og := opengraph.NewOpenGraph()
	if err := og.ProcessHTML(resp.Body); err != nil {
		return "", "", err
	}
	v := strings.ReplaceAll(og.Description, "のページです。", "")
	list := strings.Split(v, " ")
	if len(list) != 2 {
		return "", "", ErrLodeStoneInvalidFormat
	}
	return list[0], list[1], nil
}
