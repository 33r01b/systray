package modules

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Mem() (int, error) {
	stat, err := os.Open("/proc/meminfo")
	if err != nil {
		return 0, fmt.Errorf("cannot open /proc/meminfo: %w", err)
	}
	defer stat.Close()

	scanner := bufio.NewScanner(stat)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "MemAvailable") {
			cols := strings.Fields(scanner.Text())
			if len(cols) < 2 {
				return 0, errors.New("meminfo fields length mismatch")
			}

			available, err := strconv.Atoi(cols[1])
			if err != nil {
				return 0, fmt.Errorf("cannot parse mem usage: %w", err)
			}

			if available < 1024 {
				return 0, nil
			}

			return available / 1024, nil
		}
	}

	return 0, nil
}
