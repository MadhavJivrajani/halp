package morse

import (
	"os/exec"
)

// switch the capslock led off
func off(path string) error {
	cmd := "echo 0 | sudo tee " + path
	err := exec.Command("bash", "-c", cmd).Run()

	return err
}

// switch the capslock led on
func on(path string) error {
	cmd := "echo 1 | sudo tee " + path
	err := exec.Command("bash", "-c", cmd).Run()

	return err
}

// UpdateState updates the state of the capslock LED to on/off
func UpdateState(path string, isOn bool) error {
	if isOn {
		return on(path)
	}
	return off(path)
}
