package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mateuszmidor/GoStudy/GoProgrammingBlueprints/backup"
	"github.com/matryer/filedb"
)

type path struct {
	Path string
	Hash string
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()

	var (
		interval = flag.Duration("interval", 10*time.Second, "filesystem modification check interval")
		archive  = flag.String("archive", "archive", "archive directory path")
		dbpath   = flag.String("db", "db", "database path")
	)

	flag.Parse()

	m := &backup.Monitor{
		Destination: *archive,
		Archiver:    backup.ZIP,
		Paths:       make(map[string]string),
	}

	db, err := filedb.Dial(*dbpath)
	if err != nil {
		fatalErr = err
		return
	}
	defer db.Close()
	col, err := db.C("paths")
	if err != nil {
		fatalErr = err
		return
	}

	var path path
	col.ForEach(func(_ int, data []byte) bool {
		if err := json.Unmarshal(data, &path); err != nil {
			fatalErr = err
			return true
		}
		m.Paths[path.Path] = path.Hash
		return false
	})
	if fatalErr != nil {
		return
	}
	if len(m.Paths) < 1 {
		fatalErr = errors.New("No paths - use backup app to add paths")
		return
	}

	check(m, col)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-time.After(*interval):
			check(m, col)
		case <-signalChan:
			fmt.Println()
			log.Printf("Stopping...")
			return
		}
	}
}

func check(m *backup.Monitor, col *filedb.C) {
	log.Println("Checking....")
	counter, err := m.Now()
	if err != nil {
		log.Fatalln("Could not archive directories: ", err)
	}
	if counter > 0 {
		log.Printf("  Num archived directories: %d\n", counter)
		var path path
		col.SelectEach(func(_ int, data []byte) (bool, []byte, bool) {
			if err := json.Unmarshal(data, &path); err != nil {
				log.Println("Json unmarshal error. Skipping: ", err)
				return true, data, false
			}
			path.Hash, _ = m.Paths[path.Path]
			newdata, err := json.Marshal(&path)
			if err != nil {
				log.Println("Json marshall error. Skipping: ", err)
				return true, data, false
			}
			return true, newdata, false
		})
	}
}
