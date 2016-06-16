package archive

import (
	"os"
	"testing"
)

func TestCompressFolder(t *testing.T) {
	var source = "tests/data"
	var dest = "tests/compress/archive.zip"

	defer os.Remove(dest)

	if err := CompressFolder(source, dest); err != nil {
		t.Errorf(err.Error())
	}

	if err := fileExistsAndNotEmpty(dest); err != nil {
		t.Errorf(err.Error())
	}
}
