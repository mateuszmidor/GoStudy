package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mateuszmidor/GoStudy/GoProgrammingBlueprints/thesaurus"
)

func main() {
	apiKey := "111b117c0a4125f6c05193420828a33d"
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Could not fetch synonyms for '"+word+"':", err)
		}
		if len(syns) == 0 {
			log.Fatalln("Could not fetch synonyms for '"+word+"':", err)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}

}
