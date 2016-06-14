package archive

import (
	"os"
	"testing"
)

func TestCompressFolder(t *testing.T) {
	var source = "./tests"
	var dest = "./test.zip"

	CompressFolder(source, dest)

	file, err := os.Open(dest)

	if err != nil {
		t.Errorf("Archive not created")
	}

	info, err := file.Stat()
	if err != nil {
		t.Errorf("Impossible to retrieve file info")
	}

	if info.Name() != "test.zip" {
		t.Errorf("Bad archive name, want %q, got %q", "test.zip", info.Name())
	}

	if info.Size() == 0 {
		t.Errorf("Archive size is null")
	}

	os.Remove(dest)
}