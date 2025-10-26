package jsonstorage

import (
	"encoding/json"
	"io"
	"os"
	"sharedkernel"
)

type openForReadingFunc func(filename string) (io.Reader, error)
type openForWritingFunc func(filename string) (io.Writer, error)

// JSONStorage implements Storage interface
type JSONStorage struct {
	openForReading openForReadingFunc
	openForWriting openForWritingFunc
}

// NewJSONStorage is JSONStorage constructor
func NewJSONStorage(openForReading openForReadingFunc, openForWriting openForWritingFunc) *JSONStorage {
	return &JSONStorage{openForReading, openForWriting}
}

// LoadComicBooks loads ComicBooks from offline JSON file
func (j *JSONStorage) LoadComicBooks() (sharedkernel.ComicBooks, error) {
	var result sharedkernel.ComicBooks

	reader, err := j.openForReading(getStorageFileName())
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(reader).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// StoreComicBooks stores ComicBooks in offline JSON file
func (j *JSONStorage) StoreComicBooks(comicBooks sharedkernel.ComicBooks) error {
	writer, err := j.openForWriting(getStorageFileName())
	if err != nil {
		return err
	}

	if err := json.NewEncoder(writer).Encode(&comicBooks); err != nil {
		return err
	}

	return nil
}

// DefaultOpenForReading is handy implementation of openForReadingFunc
func DefaultOpenForReading(filename string) (io.Reader, error) {
	return os.Open(filename)
}

// DefaultOpenForWriting is handy implementation of openForWritingFunc
func DefaultOpenForWriting(filename string) (io.Writer, error) {
	return os.Create(filename)
}

func getStorageFileName() string {
	return "ComicBooks.json"
}
