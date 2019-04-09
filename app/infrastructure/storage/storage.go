package storage

import (
	"os"
)

type FileStorage interface {
	Init() error
	Save(text string) error
}

type fileStorage struct {
	path string
}

func NewFileStorage(path string) FileStorage {
	return &fileStorage{
		path,
	}
}

func (f *fileStorage) Init() error {
	return initFile(f.path)
}

func (f *fileStorage) Save(text string) error {
	var file, err = os.OpenFile(f.path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text + "\n")
	if err != nil {
		return err
	}

	// Save file changes.
	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}

func initFile(path string) error {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		return createFile(path)
	}

	if err := removeFile(path); err != nil {
		return err
	}

	return createFile(path)

}

func createFile(path string) error {
	var file, err = os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

func removeFile(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}
