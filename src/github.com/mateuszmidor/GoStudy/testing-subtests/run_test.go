package test

import "testing"

func TestRun(t *testing.T) {
	testCases :=
		[]struct {
			text string
			len  int
		}{
			{text: "boat", len: 4},
			{text: "rudder", len: 6},
			{text: "transom", len: 7},
		}

	for _, test := range testCases {
		t.Run(test.text, func(t *testing.T) {
			if len(test.text) != test.len {
				t.Errorf("Expected different lenght")
			}
		})
	}
}

func TestRunNicer(t *testing.T) {
	type TestCase struct {
		text string
		len  int
	}

	testCases :=
		[]TestCase{
			{text: "boat", len: 4},
			{text: "rudder", len: 6},
			{text: "transom", len: 7},
		}

	testFunc := func(test *TestCase) func(t *testing.T) {
		return func(t *testing.T) {
			if len(test.text) != test.len {
				t.Errorf("Expected different lenght")
			}
		}
	}

	for _, test := range testCases {
		t.Run(test.text, testFunc(&test))
	}
}
