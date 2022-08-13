package modules

import (
	"io/ioutil"
	"os"
)

func BatCapacity() (string, error) {
	capacity, err := ioutil.ReadFile("/sys/class/power_supply/BAT0/capacity")
	if err != nil {
		return "", err
	}

	return string(capacity[:len(capacity)-1]), nil
}

func BatCharging() (bool, error) {
	power, err := os.Open("/sys/class/power_supply/BAT0/power_now")
	if err != nil {
		return false, err
	}
	defer power.Close()

	powerNow := make([]byte, 1)
	if _, err := power.Read(powerNow); err != nil {
		return false, err
	}

	return string(powerNow) == "0", nil
}
