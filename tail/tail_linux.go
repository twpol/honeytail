package tail

import (
	"golang.org/x/sys/unix"
)

func getFileUniqueID(file string) (uint64, error) {
	stat := unix.Stat_t{}
	if err := unix.Stat(file, &stat); err != nil {
		return 0, err
	}
	return stat.Ino, nil
}
