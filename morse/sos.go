package morse

import (
	"strings"
	"time"
)

const (
	ON    = true
	OFF   = false
	units = 200
)

func morseDisplay(word string) error {
	for _, c := range word {
		err := UpdateState(ON)
		if err != nil {
			return err
		}
		if c == '.' {
			time.Sleep(time.Millisecond * 1 * units)
		} else if c == '-' {
			time.Sleep(time.Millisecond * 3 * units)
		}
		err = UpdateState(OFF)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 1 * units)
	}
	time.Sleep(time.Millisecond * 2 * units)
	return nil
}

func parse(seq string) error {
	words := strings.Split(seq, "/")
	for _, word := range words {
		err := morseDisplay(word)
		if err != nil {
			return err
		}
		err = UpdateState(ON)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 7 * units)
		err = UpdateState(OFF)
		if err != nil {
			return err
		}
	}
	return nil
}

// SendSignal 'sends' the morse signal to be
// displayed on the capslock led
func SendSignal(msg string) error {
	morse, err := GenerateMorse(msg)
	if err != nil {
		return err
	}
	err = parse(string(morse))
	return err
}
