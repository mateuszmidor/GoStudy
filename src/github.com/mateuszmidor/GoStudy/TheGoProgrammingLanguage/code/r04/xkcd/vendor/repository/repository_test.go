package repository

import (
	"fetcher"
	"sharedkernel"
	"sync"
	"testing"
	"time"
)

var comicBook1 = sharedkernel.ComicBook{Title: "Dzień Śmiechały", Transcription: "Kajkoszowy dzień żartów", URL: "", Number: 1}
var comicBook2 = sharedkernel.ComicBook{Title: "Szkoła latania", Transcription: "Mirmił robi PPL", URL: "", Number: 2}
var comicBook3 = sharedkernel.ComicBook{Title: "Wielki Turniej", Transcription: "Zawody na najwaleczniejszego barana", URL: "", Number: 3}
var comicBooks = sharedkernel.ComicBooks{comicBook1, comicBook2, comicBook3}

// cannedFetcher implements Fetcher imterface
type cannedFetcher struct {
}

func (f *cannedFetcher) StartFetchingComicBooks(c fetcher.ComicBookChannel) (int, error) {
	// simulate online fetching
	go func() {
		time.Sleep(100 * time.Millisecond)
		c <- &comicBook1

		time.Sleep(100 * time.Millisecond)
		c <- &comicBook2

		time.Sleep(100 * time.Millisecond)
		c <- &comicBook3
	}()

	return 3, nil
}

// cannedStorage implements Storage inteface
type cannedStorage struct {
	internalStorage sharedkernel.ComicBooks
}

func (s *cannedStorage) LoadComicBooks() (sharedkernel.ComicBooks, error) {
	return s.internalStorage, nil
}
func (s *cannedStorage) StoreComicBooks(comicBooks sharedkernel.ComicBooks) error {
	s.internalStorage = comicBooks
	return nil
}
func TestLoadEmptyOfflineComicBooks(t *testing.T) {
	storage := cannedStorage{sharedkernel.ComicBooks{}} // empty storage
	repository := NewRepository(nil, &storage)
	comicBooks, err := repository.LoadOfflineComicBooks()

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if len(comicBooks) != 0 {
		t.Errorf("Actual num of comic books != expected: %d != %d", len(comicBooks), 0)
	}
}

func TestLoadHandfulOfflineComicBooks(t *testing.T) {
	storage := cannedStorage{comicBooks} // 3 items in storage
	repository := NewRepository(nil, &storage)
	actualComicBooks, err := repository.LoadOfflineComicBooks()

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if len(actualComicBooks) != 3 {
		t.Errorf("Actual num of comic books != expected: %d != %d", len(actualComicBooks), 3)
	}
}

func TestStartUpdateFromOnlineService(t *testing.T) {
	const numExpectedComicBooks = 3
	asyncfetchingFinishedWG := sync.WaitGroup{}
	asyncfetchingFinishedWG.Add(1)
	numFetched := 0

	onSingleFetched := func(comicBookNo, total int) {
		numFetched++
	}

	onAllFetched := func() {
		asyncfetchingFinishedWG.Done()
	}

	fetcher := cannedFetcher{}
	storage := cannedStorage{sharedkernel.ComicBooks{}} // empty storage
	repository := NewRepository(&fetcher, &storage)
	numToFetch, err := repository.StartUpdateFromOnlineService(onSingleFetched, onAllFetched)

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if numToFetch != numExpectedComicBooks {
		t.Fatalf("Actual num of comic books to fetch != expected: %d != %d", numToFetch, numExpectedComicBooks)
	}

	// wait until fetching finished
	asyncfetchingFinishedWG.Wait()

	if numFetched != numExpectedComicBooks {
		t.Errorf("Actual num of fetched comic books != expected: %d != %d", numFetched, numExpectedComicBooks)
	}

	// offline storage updated when fetching done. Check offline storage
	loadedComicBooks, err := repository.LoadOfflineComicBooks()

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if len(loadedComicBooks) != numExpectedComicBooks {
		t.Errorf("Actual num of stored comic books != expected: %d != %d", len(loadedComicBooks), numExpectedComicBooks)
	}
}
