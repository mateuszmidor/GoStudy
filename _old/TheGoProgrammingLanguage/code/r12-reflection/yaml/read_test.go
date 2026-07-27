package yaml_test

import (
	"strings"
	"testing"

	"github.com/mateuszmidor/GoStudy/TheGoProgrammingLanguage/code/r12-reflection/yaml"
)

func TestSimpleBool(t *testing.T) {
	// input
	input := `
---
flag: true
`
	input = strings.TrimSpace(input)

	actualOut := false
	expectedOut := true

	// work
	err := yaml.Unmarshall([]byte(input), &actualOut)

	// check errors
	if err != nil {
		t.Fatalf("Unmarshall error: %s\n", err)
	}

	// check output
	if actualOut != expectedOut {
		t.Errorf("Invalid value read,\nexpected:\n%v\nactual:\n%v\n", expectedOut, actualOut)
	}
}

func TestSimpleInt(t *testing.T) {
	// input
	input := `
---
val: 345
`
	input = strings.TrimSpace(input)

	actualOut := uint(1)
	expectedOut := uint(345)

	// work
	err := yaml.Unmarshall([]byte(input), &actualOut)

	// check errors
	if err != nil {
		t.Fatalf("Unmarshall error: %s\n", err)
	}

	// check output
	if actualOut != expectedOut {
		t.Errorf("Invalid value read,\nexpected:\n%v\nactual:\n%v\n", expectedOut, actualOut)
	}
}

func TestSimpleString(t *testing.T) {
	// input
	input := `
---
text: "good practice to quote all strings"
`
	input = strings.TrimSpace(input)

	actualOut := ""
	expectedOut := "good practice to quote all strings"

	// work
	err := yaml.Unmarshall([]byte(input), &actualOut)

	// check errors
	if err != nil {
		t.Fatalf("Unmarshall error: %s\n", err)
	}

	// check output
	if actualOut != expectedOut {
		t.Errorf("Invalid value read,\nexpected:\n%v\nactual:\n%v\n", expectedOut, actualOut)
	}
}

func TestSimpleStruct(t *testing.T) {
	// input
	input := `
---
person:
 Name: "Andrzej"
 Age: 33
 Male: true
`
	input = strings.TrimSpace(input)

	type person struct {
		Name string
		Age  uint
		Male bool
	}
	actualOut := person{"", 1, false}
	expectedOut := person{"Andrzej", 33, true}

	// work
	err := yaml.Unmarshall([]byte(input), &actualOut)

	// check errors
	if err != nil {
		t.Fatalf("Unmarshall error: %s\n", err)
	}

	// check output
	if actualOut != expectedOut {
		t.Errorf("Invalid value read,\nexpected:\n%v\nactual:\n%v\n", expectedOut, actualOut)
	}
}

func TestStructInStruct(t *testing.T) {
	// input
	input := `
---
person:
 Name: "Andrzej"
 Age: 33
 Male: true
 Brother:
  Name: "Mateusz"
  Age: 32
`
	input = strings.TrimSpace(input)

	type brother struct {
		Name string
		Age  uint
	}
	type person struct {
		Name    string
		Age     uint
		Male    bool
		Brother brother
	}
	actualOut := person{"", 1, false, brother{"", 2}}
	expectedOut := person{"Andrzej", 33, true, brother{"Mateusz", 32}}

	// work
	err := yaml.Unmarshall([]byte(input), &actualOut)

	// check errors
	if err != nil {
		t.Fatalf("Unmarshall error: %s\n", err)
	}

	// check output
	if actualOut != expectedOut {
		t.Errorf("Invalid value read,\nexpected:\n%v\nactual:\n%v\n", expectedOut, actualOut)
	}
}

func TestSimpleSlice(t *testing.T) {
	// input
	input := `
---
Colors:
- "RED"
- "GREEN"
- "BLUE"
`
	input = strings.TrimSpace(input)

	actualOut := []string{}
	expectedOut := []string{"RED", "GREEN", "BLUE"}

	// work
	err := yaml.Unmarshall([]byte(input), &actualOut)

	// check errors
	if err != nil {
		t.Fatalf("Unmarshall error: %s\n", err)
	}

	// check output
	if len(actualOut) != len(expectedOut) {
		t.Fatalf("actual slice len != expected: %d != %d\n", len(actualOut), len(expectedOut))
	}

	for i := 0; i < len(actualOut); i++ {
		if actualOut[i] != expectedOut[i] {
			t.Errorf("Invalid value read,\nexpected:\n%v\nactual:\n%v\n", expectedOut, actualOut)
		}
	}
}

func TestSliceInStruct(t *testing.T) {
	// input
	input := `
---
person:
 Name: "Andrzej"
 Age: 33
 Male: true
 FavouriteNumbers:
 - 2
 - 4
 - 8
`
	input = strings.TrimSpace(input)

	type person struct {
		Name             string
		Age              uint
		Male             bool
		FavouriteNumbers []uint
	}
	actualOut := person{"", 1, false, []uint{}}
	expectedOut := person{"Andrzej", 33, true, []uint{2, 4, 8}}

	// work
	err := yaml.Unmarshall([]byte(input), &actualOut)

	// check errors
	if err != nil {
		t.Fatalf("Unmarshall error: %s\n", err)
	}

	if actualOut.Name != expectedOut.Name || actualOut.Age != expectedOut.Age || actualOut.Male != expectedOut.Male {
		t.Errorf("Invalid value read,\nexpected:\n%v\nactual:\n%v\n", expectedOut, actualOut)
	}
	// check output
	if len(actualOut.FavouriteNumbers) != len(expectedOut.FavouriteNumbers) {
		t.Fatalf("actual slice len != expected: %d != %d\n", len(actualOut.FavouriteNumbers), len(expectedOut.FavouriteNumbers))
	}

	for i := 0; i < len(actualOut.FavouriteNumbers); i++ {
		if actualOut.FavouriteNumbers[i] != expectedOut.FavouriteNumbers[i] {
			t.Errorf("Invalid value read,\nexpected:\n%v\nactual:\n%v\n", expectedOut.FavouriteNumbers, actualOut.FavouriteNumbers)
		}
	}
}

// Structs in slice not supported yet
func UnsupportedTestStructInSlice(t *testing.T) {
	// input
	input := `
---
- Name: "Andrzej"
  Age: 33
- Name: "Jurek"
  Age: 43
`
	input = strings.TrimSpace(input)

	type person struct {
		Name string
		Age  uint
	}
	actualOut := []person{}
	expectedOut := []person{person{"Andrzej", 33}, person{"Jurek", 43}}

	// work
	err := yaml.Unmarshall([]byte(input), &actualOut)

	// check errors
	if err != nil {
		t.Fatalf("Unmarshall error: %s\n", err)
	}

	// check output
	if len(actualOut) != len(expectedOut) {
		t.Fatalf("actual slice len != expected: %d != %d\n", len(actualOut), len(expectedOut))
	}

	for i := 0; i < len(actualOut); i++ {
		if actualOut[i] != expectedOut[i] {
			t.Errorf("Invalid value read,\nexpected:\n%v\nactual:\n%v\n", expectedOut, actualOut)
		}
	}
}
