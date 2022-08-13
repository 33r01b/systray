package modules

import (
	"os"
)

func Temp() (string, error) {
	stat, err := os.Open("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return "", err
	}
	defer stat.Close()

	temp := make([]byte, 2)
	if _, err := stat.Read(temp); err != nil {
		return "", err
	}

	return string(temp), nil
}
