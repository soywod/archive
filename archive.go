package archive

import (
	"os"
	"io/ioutil"
	"archive/zip"
	"path/filepath"
)

func CompressFolder(source string, dest string) error {
	zipFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	zipWriter := zip.NewWriter(zipFile)

	if filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// Registers file in archive
		fileWriter, err := zipWriter.Create(path)
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
	}) != nil {
		return err
	}

	err = zipWriter.Close()
	if err != nil {
		return err
	}

	return nil
}
