package postprocessing

import (
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/carriers"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure/csv/loading"
)

func ExtractCarriers(segments <-chan loading.CSVSegment) carriers.Carriers {
	codes := make(map[string]bool)

	for s := range segments {
		codes[s.CarrierCode] = true
	}

	var ab carriers.Builder
	for code := range codes {
		ab.Append(code)
	}

	return ab.Build()
}
