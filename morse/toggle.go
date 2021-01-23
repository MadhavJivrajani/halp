package morse

import (
	"os/exec"
	"strconv"
)

// switch the capslock led off
func off(path string) error {
	cmd := "echo 0 | sudo tee " + path
	err := exec.Command("bash", "-c", cmd).Run()

	return err
}

// switch the capslock led on
func on(path string, brightness int) error {
	cmd := "echo " + strconv.Itoa(brightness) + " | sudo tee " + path
	err := exec.Command("bash", "-c", cmd).Run()

	return err
}

// UpdateState updates the state of the capslock LED to on/off
func UpdateState(path string, brightness int) error {
	path = path + "/brightness"
	if brightness > 0 {
		return on(path, brightness)
	}
	return off(path)
}
