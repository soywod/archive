package archive

import (
	"os"
	"errors"
)

func fileExistsAndNotEmpty(path string) error {
	if fileStat, err := os.Stat(path); os.IsNotExist(err) {
		return err
	} else {
		if fileStat.Size() == 0 {
			return errors.New("File is empty")
		}
	}

	return nil
}