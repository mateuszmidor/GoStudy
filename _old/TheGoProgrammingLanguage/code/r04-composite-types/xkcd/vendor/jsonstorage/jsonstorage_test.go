package jsonstorage

import (
	"bytes"
	"encoding/json"
	"io"
	"sharedkernel"
	"strings"
	"testing"
)

func readFunc(filename string) (io.Reader, error) {
	const comicBookJSON = `[{"Title": "title", "Transcription": "transcription", "URL" : "url", "Number" : 1 }]`
	reader := io.Reader(bytes.NewBufferString(comicBookJSON))
	return reader, nil
}

func TestLoadComicBooks(t *testing.T) {

	s := NewJSONStorage(readFunc, nil)
	comicBooks, err := s.LoadComicBooks()

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if len(comicBooks) != 1 {
		t.Fatalf("Actual comic book count != expected: %d != %d", len(comicBooks), 1)
	}

	comic := comicBooks[0]

	if comic.Number != 1 {
		t.Errorf("Actual comic book number != expected: %d != %d", comic.Number, 1)
	}

	if comic.Title != "title" {
		t.Errorf("Actual title != expected: %s != %s", comic.Title, "title")
	}

	if comic.Transcription != "transcription" {
		t.Errorf("Actual transcription != expected: %s != %s", comic.Transcription, "transcription")
	}

	if comic.URL != "url" {
		t.Errorf("Actual url != expected: %s != %s", comic.URL, "url")
	}
}

func TestStoreComicBooks(t *testing.T) {
	var b bytes.Buffer
	writer := io.Writer(&b)
	writeFunc := func(filename string) (io.Writer, error) {
		return writer, nil
	}
	s := NewJSONStorage(nil, writeFunc)
	comicBooks := sharedkernel.ComicBooks{
		{Title: "title", Transcription: "transcription", URL: "url", Number: 1},
	}

	err := s.StoreComicBooks(comicBooks)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	// check json can be read again
	var comicBooks2 sharedkernel.ComicBooks
	jsonString := b.String()
	reader := strings.NewReader(jsonString)
	if err := json.NewDecoder(reader).Decode(&comicBooks2); err != nil {
		t.Fatalf("Enexpectd error: %s", err)
	}

	if len(comicBooks2) != 1 {
		t.Fatalf("Actual comic book count != expected: %d != %d", len(comicBooks2), 1)
	}

	actual := comicBooks2[0]
	expected := comicBooks[0]
	if actual != expected {
		t.Errorf("Actual comic book != expected comic book:\n %v\n!=\n%v", actual, expected)
	}
}
