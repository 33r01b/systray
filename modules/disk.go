package modules

import (
	"math"
	"syscall"
)

func DiskAvailable() (uint64, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err != nil {
		return 0, err
	}

	return fs.Bavail * uint64(fs.Bsize) / uint64(math.Pow(float64(1024), 3)), nil
}
