package main

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadFile(url string, filepath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	outFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, response.Body)
	return err
}

func unzip(source string, destination string) error {
	zipReader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		filePath := filepath.Join(destination, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, file.Mode())
			continue
		}

		outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		_, err = io.Copy(outFile, fileReader)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := downloadFile("https://github.com/your-repo/mods.zip", "mods.zip")
	if err != nil {
		panic(err)
	}
	err = unzip("mods.zip", "path/to/install")
	if err != nil {
		panic(err)
	}
}
