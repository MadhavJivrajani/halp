package morse

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	ON    = true
	OFF   = false
	units = 200
)

func morseDisplay(path, word string) error {
	for _, c := range word {
		err := UpdateState(path, ON)
		if err != nil {
			return err
		}
		if c == '.' {
			time.Sleep(time.Millisecond * 1 * units)
		} else if c == '-' {
			time.Sleep(time.Millisecond * 3 * units)
		}
		err = UpdateState(path, OFF)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 1 * units)
	}
	time.Sleep(time.Millisecond * 2 * units)
	return nil
}

func parse(path, seq string) error {
	words := strings.Split(seq, "/")
	for _, word := range words {
		err := morseDisplay(path, word)
		if err != nil {
			return err
		}
		err = UpdateState(path, ON)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 7 * units)
		err = UpdateState(path, OFF)
		if err != nil {
			return err
		}
	}
	return nil
}

// SendSignal 'sends' the morse signal to be
// displayed on the capslock led
func SendSignal(path, msg string) error {
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("invalid led path!")
	}
	morse, err := GenerateMorse(msg)
	if err != nil {
		return err
	}
	err = parse(path, string(morse))
	return err
}
