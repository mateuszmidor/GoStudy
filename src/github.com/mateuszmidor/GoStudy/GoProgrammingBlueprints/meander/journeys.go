package meander

import "strings"

type j struct {
	Name      string
	PlaceType []string
}

var Journeys = []interface{}{
	j{Name: "Romantic day", PlaceType: []string{"park", "bar", "movie_theatre", "restaurant", "florist", "taxi_stand"}},
	j{Name: "Shopping frenzy", PlaceType: []string{"department_store", "cafe", "clothing_store", "jewelry_store", "shoe_store"}},
	j{Name: "Night escapade", PlaceType: []string{"bar", "casino", "food", "bar", "night_club", "bar", "bar", "hostpital"}},
	j{Name: "Full culutre", PlaceType: []string{"museum", "cafe", "cemetery", "library", "art_galery"}},
	j{Name: "Moment for yourself", PlaceType: []string{"hair_care", "beauty_salon", "cafe", "spa"}},
}

func (journey j) Public() interface{} {
	return map[string]interface{}{
		"name":    journey.Name,
		"journey": strings.Join(journey.PlaceType, "|"),
	}
}
