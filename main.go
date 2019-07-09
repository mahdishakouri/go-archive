package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"log"
	"os"
)

func listFiles(file *zip.File) error {
	fileread, err := file.Open()
	if err != nil {
		msg := "Failed to open zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
	defer fileread.Close()

	fileText := fmt.Sprintf("%s:\n", file.Name) + "\n"

	fileNew, _ := os.OpenFile("fileContent.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	datawriter := bufio.NewWriter(fileNew)

	_, _ = datawriter.WriteString(fileText)

	datawriter.Flush()
	fileNew.Close()

	return nil
}

func main() {
	read, err := zip.OpenReader("archiveFile.zip")
	if err != nil {
		msg := "Failed to open: %s"
		log.Fatalf(msg, err)
	}
	defer read.Close()

	for _, file := range read.File {
		if err := listFiles(file); err != nil {
			log.Fatalf("Failed to read %s from zip: %s", file.Name, err)
		}
	}

}
