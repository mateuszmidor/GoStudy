package main

import "testing"

func TestEmptyFormula(t *testing.T) {
	_, err := NewEvalXY("")
	if err == nil {
		t.Error("There should be error returned for empty formula")
	}
}

func TestErronousFormula(t *testing.T) {
	formula := "2 + 3 *"
	_, err := NewEvalXY(formula)
	if err == nil {
		t.Errorf("There should be error returned for erronous formula: %q", formula)
	}
}

func TestNoArgsFormula(t *testing.T) {
	formula := "2 * 3 + 4"
	x := 0.0
	y := 0.0
	expected := 10.0

	eval, err := NewEvalXY(formula)

	if err != nil {
		t.Errorf("Unexpected error returned for formula: %q - %s", formula, err)
	}

	result, err := eval.Eval(x, y)
	if err != nil {
		t.Errorf("Unexpected error returned for formula %q [x=%f y=%f]: %s", formula, x, y, err)
	}

	if result != expected {
		t.Errorf("Result of formula %q [x=%f y=%f] should be %f but was: %f", formula, x, y, expected, result)
	}
}

func TestTwoAgsFormula(t *testing.T) {
	formula := "2 * x + y"
	x := 2.0
	y := 1.0
	expected := 5.0

	eval, err := NewEvalXY(formula)

	if err != nil {
		t.Errorf("Unexpected error returned for formula: %q - %s", formula, err)
	}

	result, err := eval.Eval(x, y)
	if err != nil {
		t.Errorf("Unexpected error returned for formula %q [x=%f y=%f]: %s", formula, x, y, err)
	}

	if result != expected {
		t.Errorf("Result of formula %q [x=%f y=%f] should be %f but was: %f", formula, x, y, expected, result)
	}
}

func TestSinFunc(t *testing.T) {
	formula := "sin(1.57079632679)" // PI/2
	x := 0.0
	y := 0.0
	expected := 1.0

	eval, err := NewEvalXY(formula)
	if err != nil {
		t.Errorf("Unexpected error returned for formula: %q - %s", formula, err)
	}

	result, err := eval.Eval(x, y)
	if err != nil {
		t.Errorf("Unexpected error returned for formula %q [x=%f y=%f]: %s", formula, x, y, err)
	}

	if result != expected {
		t.Errorf("Result of formula %q [x=%f y=%f] should be %f but was: %f", formula, x, y, expected, result)
	}
}

func TestCosFunc(t *testing.T) {
	formula := "cos(0.0)"
	x := 0.0
	y := 0.0
	expected := 1.0

	eval, err := NewEvalXY(formula)
	if err != nil {
		t.Errorf("Unexpected error returned for formula: %q - %s", formula, err)
	}

	result, err := eval.Eval(x, y)
	if err != nil {
		t.Errorf("Unexpected error returned for formula %q [x=%f y=%f]: %s", formula, x, y, err)
	}

	if result != expected {
		t.Errorf("Result of formula %q [x=%f y=%f] should be %f but was: %f", formula, x, y, expected, result)
	}
}

func TestPowFunc(t *testing.T) {
	formula := "pow(2.0, 3.0)"
	x := 0.0
	y := 0.0
	expected := 8.0

	eval, err := NewEvalXY(formula)
	if err != nil {
		t.Errorf("Unexpected error returned for formula: %q - %s", formula, err)
	}

	result, err := eval.Eval(x, y)
	if err != nil {
		t.Errorf("Unexpected error returned for formula %q [x=%f y=%f]: %s", formula, x, y, err)
	}

	if result != expected {
		t.Errorf("Result of formula %q [x=%f y=%f] should be %f but was: %f", formula, x, y, expected, result)
	}
}

func TestHypotFunc(t *testing.T) {
	formula := "hypot(3.0, 4.0)"
	x := 0.0
	y := 0.0
	expected := 5.0

	eval, err := NewEvalXY(formula)
	if err != nil {
		t.Errorf("Unexpected error returned for formula: %q - %s", formula, err)
	}

	result, err := eval.Eval(x, y)
	if err != nil {
		t.Errorf("Unexpected error returned for formula %q [x=%f y=%f]: %s", formula, x, y, err)
	}

	if result != expected {
		t.Errorf("Result of formula %q [x=%f y=%f] should be %f but was: %f", formula, x, y, expected, result)
	}
}
