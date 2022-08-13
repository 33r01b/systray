package modules

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var prevTotal, prevIdle int

func Cpu() (int, error) {
	total, idle, err := stat()
	if err != nil {
		return 0, err
	}

	var cpuUsage int
	if prevTotal > 0 && prevIdle > 0 && (total-prevTotal) > 0 {
		cpuUsage = 100 * ((total - prevTotal) - (idle - prevIdle)) / (total - prevTotal)
	}

	prevTotal = total
	prevIdle = idle

	return cpuUsage, nil
}

func stat() (total, idle int, err error) {
	stat, err := os.Open("/proc/stat")
	if err != nil {
		return 0, 0, fmt.Errorf("cannot open /proc/stat: %w", err)
	}
	defer stat.Close()

	reader := bufio.NewReader(stat)
	ln, _, err := reader.ReadLine()
	if err != nil {
		return 0, 0, fmt.Errorf("cannot read /proc/stat: %w", err)
	}

	cols := strings.Fields(string(ln))
	if len(cols) < 5 {
		return 0, 0, errors.New("stat fields length mismatch")
	}

	user, err := strconv.Atoi(cols[1])
	if err != nil {
		return 0, 0, fmt.Errorf("cannot parse cpu 'user' column: %w", err)
	}

	nice, err := strconv.Atoi(cols[2])
	if err != nil {
		return 0, 0, fmt.Errorf("cannot parse cpu 'nice' column: %w", err)
	}

	system, err := strconv.Atoi(cols[3])
	if err != nil {
		return 0, 0, fmt.Errorf("cannot parse cpu 'system' column: %w", err)
	}

	idle, err = strconv.Atoi(cols[4])
	if err != nil {
		return 0, 0, fmt.Errorf("cannot parse cpu 'idle' column: %w", err)
	}

	total = user + nice + system + idle

	return total, idle, err
}
