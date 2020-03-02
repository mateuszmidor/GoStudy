package storage

import (
	"sharedkernel"
)

// Storage is interface for offline comic books storage
type Storage interface {
	LoadComicBooks() (sharedkernel.ComicBooks, error)
	StoreComicBooks(comicBooks sharedkernel.ComicBooks) error
}
