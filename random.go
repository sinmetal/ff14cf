package ff14cf

import (
	"errors"
	"math/rand"
)

func MakeRandomStr(digit uint32) (string, error) {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error")
	}

	result := make([]rune, digit)
	for i, v := range b {
		result[i] = letters[int(v)%len(letters)]
	}
	return string(result), nil
}
