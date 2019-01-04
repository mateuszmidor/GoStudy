package thesaurus

import (
	"encoding/json"
	"errors"
	"net/http"
)

// synonym generator based on http://words.bighugelabs.com/
// API key: 111b117c0a4125f6c05193420828a33d
// API example request: http://words.bighugelabs.com/api/2/111b117c0a4125f6c05193420828a33d/handsome/json

type BigHuge struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	request := "http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json"
	response, err := http.Get(request)
	if err != nil {
		return syns, errors.New("bighuge: Couldnt fetch synonyms of '" + term + "'" + err.Error())
	}
	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}
	if data.Noun != nil {
		syns = append(syns, data.Noun.Syn...)
	}
	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}
	return syns, nil
}
