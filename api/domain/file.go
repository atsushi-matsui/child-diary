package domain

import (
	"log"
	"os"
)

type MyFile struct {
	Dir string
	FileName string
	Extension string
	File *os.File
}

func NewFile(dir string, fileName string, extension string) (*MyFile, error) {
	file, err := os.OpenFile(dir + fileName + "." + extension, os.O_RDWR|os.O_CREATE, 0666)
    if err != nil {
		log.Fatalf("Failed to read file: %v", err)
    }

	return &MyFile{
		Dir: dir,
		FileName: fileName,
		Extension: extension,
		File: file,
	}, nil
}


func (file *MyFile) Write(content string) {
	file.File.WriteString(content)
}


func (file *MyFile) Close() {
	file.File.Close()
}
