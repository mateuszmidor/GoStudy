package dataloading

import "github.com/mateuszmidor/GoStudy/flight-finder/src/nations"

// FilterRawNations turns stream of RawNation into Nations list
func FilterRawNations(rawnations <-chan RawNation) nations.Nations {
	nb := nations.Builder{}

	for n := range rawnations {
		nb.Append(n.Code, n.Iso, n.Currency, n.Name)
	}

	return nb.Build()
}
