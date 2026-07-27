package repository

import (
	"fetcher"
	"sharedkernel"
	"storage"
	"sync"
)

// OnComicBookFetched is callback type to be called when every comicBook fetched
type OnComicBookFetched func(comicBookNo, total int)

// OnOfflineComicBooksUpdated is callback type to be called when online comic books are saved to offline storage
type OnOfflineComicBooksUpdated = func()

// Repository is ComicBook offline source with online updates capability
type Repository struct {
	fetcher      fetcher.Fetcher
	storage      storage.Storage
	storageMutex sync.Mutex // storage is accessed concurrently: load/update
}

// NewRepository is Repository constructo
func NewRepository(f fetcher.Fetcher, s storage.Storage) *Repository {
	return &Repository{f, s, sync.Mutex{}}
}

// LoadOfflineComicBooks loads comic books from offline storage
// Notice: async access to r.storage from user/StartUpdateFromOnlineService
func (r *Repository) LoadOfflineComicBooks() (sharedkernel.ComicBooks, error) {
	r.storageMutex.Lock()
	comicBooks, err := r.storage.LoadComicBooks()
	r.storageMutex.Unlock()

	return comicBooks, err
}

// saveOfflineComicBooks saves comic books  got from StartUpdateFromOnlineService to offline storage
// Notice: async access to r.storage from user/StartUpdateFromOnlineService
func (r *Repository) saveOfflineComicBooks(comicBooks sharedkernel.ComicBooks) error {
	r.storageMutex.Lock()
	err := r.storage.StoreComicBooks(comicBooks)
	r.storageMutex.Unlock()

	return err
}

// StartUpdateFromOnlineService begins the process of updating comic book repository from online service
func (r *Repository) StartUpdateFromOnlineService(onComicBookFetched OnComicBookFetched, onOfflineComicBooksUpdated OnOfflineComicBooksUpdated) (int, error) {
	// start fetchin online into channel
	c := make(fetcher.ComicBookChannel)
	count, err := r.fetcher.StartFetchingComicBooks(c)
	if count == 0 || err != nil {
		return 0, err
	}

	// start processing comic books that show up in the channel
	go func() {
		tmpComicBooks := make(sharedkernel.ComicBooks, count)
		for i := 0; i < count; i++ {
			comicBookPtr := <-c
			if onComicBookFetched != nil {
				onComicBookFetched(i+1, count) // i + 1 so finally i == count
			}
			if comicBookPtr != nil {
				tmpComicBooks[i] = *comicBookPtr
			}
		}
		r.saveOfflineComicBooks(tmpComicBooks)
		if onOfflineComicBooksUpdated != nil {
			onOfflineComicBooksUpdated()
		}
	}()

	// return the number of comic books to be finally fetched
	return count, nil
}
