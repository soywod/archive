package archive

import (
	"testing"
)

func TestHelper(t *testing.T) {
	var tests = []struct {
		path string
		ok   bool
	}{
		{"tests/data/fileA.txt", true},
		{"tests/data/folder", true},
		{"tests/data/empty.txt", false},
		{"tests/data/noexists.txt", false},
	}

	for _, test := range tests {
		err := fileExistsAndNotEmpty(test.path)

		if err != nil && test.ok {
			t.Errorf(err.Error())
		}

		if err == nil && ! test.ok {
			t.Errorf("File %q expected %t got %t", test.path, test.ok, ! test.ok)
		}
	}
}
