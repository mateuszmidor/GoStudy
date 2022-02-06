package postprocessing

import (
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/nations"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure/csv/loading"
)

// FilterNations turns stream of CSVNation into Nations list
func FilterNations(rawnations <-chan loading.CSVNation) nations.Nations {
	nb := nations.Builder{}

	for n := range rawnations {
		nb.Append(n.Code, n.Iso, n.Currency, n.Name)
	}

	return nb.Build()
}
