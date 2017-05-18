package storage

import (
	"log"
	"os"
	"fmt"
	"io/ioutil"
)

func (dao FileSystemDao) Save(id string, content []byte) (string, error) {
	directory := directory()
	filename :=  fmt.Sprintf("%s/email-%s.html", directory, id)
	err := ioutil.WriteFile(filename, content, 0666)
	if err != nil {
		return "", err
	}
	log.Printf("saved email to: %s ", filename)
	return filename, nil
}

func directory() string {
	directory := os.Getenv("ses.output.directory")
	if directory == "" {
		return os.TempDir()
	}
	return directory
}

type EmailDao interface {
	Save(id string, content []byte) (string, error)
}

type FileSystemDao struct {
}

type InMemoryDao struct {
}
