package archive

import (
	"os"
	"testing"
)

func TestDecompress(t *testing.T) {
	var source = "tests/data/archive.zip"
	var dest = "tests/decompress/all/"
	var paths = []string{
		dest + "fileA.txt",
		dest + "fileB.txt",
		dest + "folder/fileC.txt",
	}

	defer os.RemoveAll(dest)

	if err := Decompress(source, dest); err != nil {
		t.Errorf(err.Error())
	}

	for _, path := range paths {
		if err := fileExistsAndNotEmpty(path); err != nil {
			t.Errorf(err.Error())
		}
	}
}
