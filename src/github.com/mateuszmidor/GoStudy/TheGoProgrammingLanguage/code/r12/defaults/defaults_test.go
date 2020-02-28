package defaults

import "testing"

type address struct {
	City   string `default:"Krakow"` // `default` is custom tag that holds default value used by SetEmptyFieldsToDefault()
	Street string
	Numbe  uint `default:"25"`
}

func TestNoFieldsSet(t *testing.T) {
	// input
	actualParams := address{}
	expectedParams := address{"Krakow", "", 25}

	// work
	SetEmptyFieldsToDefault(&actualParams)

	// check
	if actualParams != expectedParams {
		t.Errorf("actual RequestParams != expected: %v != %v", actualParams, expectedParams)
	}
}

func TestAllFieldsSet(t *testing.T) {
	// input
	actualParams := address{"Gdynia", "Targ Solny", 100}
	expectedParams := address{"Gdynia", "Targ Solny", 100}

	// work
	SetEmptyFieldsToDefault(&actualParams)

	// check
	if actualParams != expectedParams {
		t.Errorf("actual RequestParams != expected: %v != %v", actualParams, expectedParams)
	}
}
