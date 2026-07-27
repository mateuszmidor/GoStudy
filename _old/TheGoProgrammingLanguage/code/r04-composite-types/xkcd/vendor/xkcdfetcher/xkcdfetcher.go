package xkcdfetcher

import (
	"encoding/json"
	"fetcher"
	"fmt"
	"net/http"
	"sharedkernel"
)

// xkcdComicBook is the JSON format xkcd.com API uses for providing comic books
type xkcdComicBook struct {
	Month      string `json:"month"`           // "2"
	Day        string `json:"day"`             // "5"
	Num        int    `json:"num"`             // 2264
	Year       string `json:"year"`            // "2020"
	SafeTitle  string `json:"titlesafe_title"` // "Satellite"
	Transcript string `json:"transcript"`      // ""
	Alt        string `json:"alt"`             // "If you're going to let it burn up, make sure it happens over the deep end of the bathtub and not any populated parts of the house!"
	ImgURL     string `json:"img"`             // "https://imgs.xkcd.com/comics/satellite.png"
	Title      string `json:"title"`           // "Satellite"
}

// SingleComicBookFetcher is a function fetching single ComicBook stored at "url"
type singleComicBookFetcher = func(url string) (*sharedkernel.ComicBook, error)

// XkcdFetcher implements Fetcher interface
type XkcdFetcher struct {
	fetchSingleComicBook singleComicBookFetcher // func
}

// NewXkcdFetcher is XkcdFetcher constructor
func NewXkcdFetcher(singleComicBookFetcher singleComicBookFetcher) *XkcdFetcher {
	return &XkcdFetcher{singleComicBookFetcher}
}

// StartFetchingComicBooks begins the process of fetching comic books from xkcd.com
func (f *XkcdFetcher) StartFetchingComicBooks(c fetcher.ComicBookChannel) (int, error) {
	latestComicBook, err := f.fetchSingleComicBook(getURLForLatestComicBook())
	if err != nil {
		return 0, err
	}

	numOfLatestComicBook := latestComicBook.Number

	// limit num of books
	if numOfLatestComicBook > 15 {
		numOfLatestComicBook = 15
		fmt.Print("Number of comic books to fetch was limited to 15")
	}

	// to be made even more async
	go func() {
		for i := 1; i <= numOfLatestComicBook; i++ {
			url := getURLForComicBookNo(i)
			comicBook, err := f.fetchSingleComicBook(url)
			if err != nil {
				fmt.Printf("Errorf fetching comic book %s: %s", url, err)
			}
			c <- comicBook // nil if error
		}
	}()

	return numOfLatestComicBook, nil
}

// DefaultSingleComicBookFetcher is handy implementation of SingleComicBookFetcher
func DefaultSingleComicBookFetcher(url string) (*sharedkernel.ComicBook, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Query %q failed: %s", url, resp.Status)
	}

	var xkcd xkcdComicBook
	if err := json.NewDecoder(resp.Body).Decode(&xkcd); err != nil {
		return nil, err
	}

	return xkcdToComicBook(xkcd), nil
}

func getURLForLatestComicBook() string {
	return "https://xkcd.com/info.0.json"
}

func getURLForComicBookNo(number int) string {
	return fmt.Sprintf("https://xkcd.com/%d/info.0.json", number)
}

func xkcdToComicBook(in xkcdComicBook) (result *sharedkernel.ComicBook) {
	return &sharedkernel.ComicBook{
		Number:        in.Num,
		Title:         in.Title,
		Transcription: in.Transcript,
		URL:           in.ImgURL,
	}
}
