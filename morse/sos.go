package morse

import (
	"fmt"
	"os"
	"strings"
	"time"
	"strconv"
)

const (
	units = 200
)

func morseDisplay(path, word string, brightness int) error {
	for _, c := range word {
		err := UpdateState(path, brightness)
		if err != nil {
			return err
		}
		if c == '.' {
			time.Sleep(time.Millisecond * 1 * units)
		} else if c == '-' {
			time.Sleep(time.Millisecond * 3 * units)
		}
		err = UpdateState(path, 0)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 1 * units)
	}
	time.Sleep(time.Millisecond * 2 * units)
	return nil
}

func readFile(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	state := make([]byte, 10)
	count, err := file.Read(state)
	if err != nil {
		return 0, err
	}
	content, err := strconv.Atoi(string(state[:count - 1]))
			
	return content, nil
}

func getInitialState(path string) (int, error) {
	brightness, err := readFile(path + "/brightness")
	return brightness, err
}

func getMaxBrightness(path string) (int, error) {
	max, err := readFile(path + "/max_brightness")
	return max, err	
}

func parse(path, seq string) error {
	initState, err := getInitialState(path)
	if err != nil {
		return err
	}
	maxBrightness, err := getMaxBrightness(path)
	if err != nil {
		return err
	}
	words := strings.Split(seq, "/")
	err = UpdateState(path, 0) // switch off LED initially
	if err != nil {
		return err
	}
	for _, word := range words {
		err := morseDisplay(path, word, maxBrightness)
		if err != nil {
			return err
		}
		err = UpdateState(path, maxBrightness)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 7 * units)
		err = UpdateState(path, 0)
		if err != nil {
			return err
		}
	}
	err = UpdateState(path, initState) // restore state of LED to initial state
	if err != nil {
		return err
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
