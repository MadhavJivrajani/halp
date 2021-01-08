package morse

import (
	"strings"

	"github.com/alwindoss/morse"
)

// GenerateMorse generates an array of unsigned ints
// representing the morse code equivalent of `msg`
func GenerateMorse(msg string) ([]uint8, error) {
	h := morse.NewHacker()

	morseCode, err := h.Encode(strings.NewReader(msg))
	if err != nil {
		return nil, err
	}
	return morseCode, nil
}
