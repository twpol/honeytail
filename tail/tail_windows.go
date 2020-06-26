package tail

import (
	"os"
	"syscall"
)

func getFileUniqueID(file string) (uint64, error) {
	stat, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	winStat := stat.Sys().(*syscall.Win32FileAttributeData)
	return uint64(winStat.CreationTime.Nanoseconds()), nil
}
