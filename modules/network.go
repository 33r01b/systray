package modules

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

var rxPrev int
var txPrev int

func Network() (down int, up int, err error) {
	rxCurrent, err := netTraffic("/sys/class/net/[ew]*/statistics/rx_bytes")
	if err != nil {
		if err != nil {
			return 0, 0, err
		}
	}

	down = (rxCurrent - rxPrev) / 1024
	rxPrev = rxCurrent

	txCurrent, err := netTraffic("/sys/class/net/[ew]*/statistics/tx_bytes")
	if err != nil {
		if err != nil {
			return 0, 0, err
		}
	}

	up = (txCurrent - txPrev) / 1024
	txPrev = txCurrent

	return down, up, nil
}

func netTraffic(statPathsPattern string) (int, error) {
	netStat := 0

	rxPaths, err := filepath.Glob(statPathsPattern)
	if err != nil {
		return 0, err
	}

	for _, rp := range rxPaths {
		rxc, err := readTraffic(rp)
		if err != nil {
			return 0, err
		}

		netStat += rxc
	}

	return netStat, nil
}

func readTraffic(path string) (int, error) {
	netStat, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("cannot open %s: %w", path, err)
	}
	defer netStat.Close()

	reader := bufio.NewReader(netStat)
	ln, _, err := reader.ReadLine()
	if err != nil {
		return 0, fmt.Errorf("cannot read %s: %w", path, err)
	}

	netBytes, err := strconv.Atoi(string(ln))
	if err != nil {
		return 0, fmt.Errorf("cannot parse %s: %w", path, err)
	}

	return netBytes, err
}
