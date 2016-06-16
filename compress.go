package archive

import (
	"os"
	"io/ioutil"
	"archive/zip"
	"path/filepath"
)

var zipWriter *zip.Writer
var prefix int

func CompressFolder(source string, dest string) error {
	prefix = -1

	zipFile, err := os.Create(dest)
	defer zipFile.Close()
	if err != nil {
		return err
	}

	zipWriter = zip.NewWriter(zipFile)
	defer zipWriter.Close()

	if err := filepath.Walk(source, walk); err != nil {
		return err
	}

	return nil
}

func walk(path string, info os.FileInfo, err error) error {
	if prefix == -1 {
		prefix = len(path)
	}

	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	// Creates file in archive
	fileWriter, err := zipWriter.Create(path[prefix:])
	if err != nil {
		return err
	}

	// Reads file content
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Puts file content into archive
	_, err = fileWriter.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}
