package morse

import (
	"os/exec"
)

// switch the capslock led off
func off() error {
	cmd := "echo 0 | sudo tee /sys/class/leds/input3::capslock/brightness"
	err := exec.Command("bash", "-c", cmd).Run()

	return err
}

// switch the capslock led on
func on() error {
	cmd := "echo 1 | sudo tee /sys/class/leds/input3::capslock/brightness"
	err := exec.Command("bash", "-c", cmd).Run()

	return err
}

// UpdateState updates the state of the capslock LED to on/off
func UpdateState(isOn bool) error {
	if isOn {
		return on()
	}
	return off()
}
