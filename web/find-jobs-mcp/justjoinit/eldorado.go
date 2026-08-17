package justjoinit

import (
	"strings"

	"findjobsmcp/api"
)

// skillNames extracts skill names from a skill slice.
func skillNames(skills []skill) []string {
	names := make([]string, len(skills))
	for i, s := range skills {
		names[i] = s.Name
	}
	return names
}

// offersToEldorado converts JustJoin offers into the Eldorado representation.
func offersToEldorado(offers []offer) api.EldoradoOffers {
	jobs := make([]api.EldoradoOffer, 0, len(offers))
	for _, o := range offers {
		jobs = append(jobs, offerToEldorado(o))
	}
	return api.EldoradoOffers{Jobs: jobs}
}

// offerToEldorado maps a single JustJoin offer to an api.EldoradoOffer.
func offerToEldorado(offer offer) api.EldoradoOffer {
	return api.EldoradoOffer{
		Title:           offer.Title,
		Keywords:        skillNames(append(append([]skill{}, offer.RequiredSkills...), offer.NiceToHaveSkills...)),
		Company:         offer.CompanyName,
		Cities:          eldoradoCities(offer),
		WorkModes:       []string{eldoradoWorkMode(offer.WorkplaceType)},
		ContractTypes:   eldoradoContractTypes(offer.EmploymentTypes),
		EmploymentTypes: []string{eldoradoEmploymentType(offer.WorkingTime)},
		Seniority:       offer.ExperienceLevel,
		SalaryFrom:      eldoradoSalaryBound(offer.EmploymentTypes, true),
		SalaryTo:        eldoradoSalaryBound(offer.EmploymentTypes, false),
		URL:             offer.OfferURL,
	}
}

// eldoradoWorkMode maps JustJoin workplace type to the Eldorado work mode.
func eldoradoWorkMode(workplaceType string) string {
	if workplaceType == "partly_remote" {
		return "hybrid"
	}
	return workplaceType
}

// eldoradoContractTypes maps JustJoin employment types to Eldorado contract
// types, deduplicated (the API repeats the same type once per currency).
func eldoradoContractTypes(types []employmentType) []string {
	contracts := make([]string, 0, len(types))
	seen := make(map[string]bool)
	for _, t := range types {
		var contract string
		switch strings.ToLower(t.Type) {
		case "b2b":
			contract = "b2b"
		case "uop", "permanent":
			contract = "employment_contract"
		case "uoz":
			contract = "mandate_contract"
		default:
			contract = strings.ToLower(t.Type)
		}
		if contract == "" || seen[contract] {
			continue
		}
		seen[contract] = true
		contracts = append(contracts, contract)
	}
	return contracts
}

// eldoradoEmploymentType maps JustJoin working time to the Eldorado employment type.
func eldoradoEmploymentType(workingTime string) string {
	if strings.Contains(workingTime, "part") {
		return "part_time"
	}
	return "full_time"
}

// eldoradoSalaryBound returns the lower (or upper, when from is false) salary
// bound of the first employment type that discloses both, or nil when
// no salary information is available.
func eldoradoSalaryBound(types []employmentType, from bool) *int {
	for _, t := range types {
		if t.From == nil || t.To == nil {
			continue
		}
		v := *t.To
		if from {
			v = *t.From
		}
		salary := int(v)
		return &salary
	}
	return nil
}

// eldoradoCities returns unique city names from the offer's structured
// locations, falling back to the plain City field when there are none.
func eldoradoCities(o offer) []string {
	var cities []string
	seen := make(map[string]bool)
	add := func(c string) {
		if c == "" || seen[c] {
			return
		}
		seen[c] = true
		cities = append(cities, c)
	}
	if len(o.Locations) == 0 {
		add(o.City)
	} else {
		for _, l := range o.Locations {
			add(l.City)
		}
	}
	return cities
}
