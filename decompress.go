package archive

import (
	"io"
	"os"
	"strings"
	"archive/zip"
	"path/filepath"
)

func Decompress(source string, dest string) error {
	// Creates a new zip reader
	zipReadCloser, err := zip.OpenReader(source)
	defer zipReadCloser.Close()
	if err != nil {
		return err
	}

	// Creates the dest folder
	if err := os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	// Decompresses each file
	for _, zipFile := range zipReadCloser.File {
		path := filepath.Join(dest, zipFile.Name)

		if err := decompressFile(zipFile, path); err != nil {
			return nil
		}
	}

	return nil
}

func decompressFile(z *zip.File, path string) error {
	zipFileInfo := z.FileInfo()

	// Open the current zip file
	zipFile, err := z.Open()
	defer zipFile.Close()
	if err != nil {
		return err
	}

	if zipFileInfo.IsDir() {
		if os.MkdirAll(path, zipFileInfo.Mode()) != nil {
			return err
		}
	} else {
		// Creates the last folder
		if lastSlash := strings.LastIndex(path, string(os.PathSeparator)); lastSlash > -1 {
			if err := os.MkdirAll(path[:lastSlash], 0755); err != nil {
				return err
			}
		}

		// Creates the file
		file, err := os.Create(path)
		defer file.Close()
		if err != nil {
			return err
		}

		// Copy data from zip to file
		if _, err := io.Copy(file, zipFile); err != nil {
			return err
		}
	}

	return nil
}