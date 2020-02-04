package stringgames

import "testing"

const getNumTripletsFailureString = "Expected %d triplets, got %d for %q"
const commaFailureString = "Expected %q, got %q"
const getTripletFailureString = "Expected getTriplet(%d) to return %q, got %q"

func TestGetNumTriplets0(t *testing.T) {
	in := ""
	expectedOut := 0
	actualOut := getNumTriplets(in)
	if actualOut != expectedOut {
		t.Errorf(getNumTripletsFailureString, expectedOut, actualOut, in)
	}
}

func TestGetNumTriplets1(t *testing.T) {
	in := "1"
	expectedOut := 1
	actualOut := getNumTriplets(in)
	if actualOut != expectedOut {
		t.Errorf(getNumTripletsFailureString, expectedOut, actualOut, in)
	}
}

func TestGetNumTriplets2(t *testing.T) {
	in := "12"
	expectedOut := 1
	actualOut := getNumTriplets(in)
	if actualOut != expectedOut {
		t.Errorf(getNumTripletsFailureString, expectedOut, actualOut, in)
	}
}
func TestGetNumTriplets3(t *testing.T) {
	in := "123"
	expectedOut := 1
	actualOut := getNumTriplets(in)
	if actualOut != expectedOut {
		t.Errorf(getNumTripletsFailureString, expectedOut, actualOut, in)
	}
}
func TestGetNumTriplets1234(t *testing.T) {
	in := "1234"
	expectedOut := 2
	actualOut := getNumTriplets(in)
	if actualOut != expectedOut {
		t.Errorf(getNumTripletsFailureString, expectedOut, actualOut, in)
	}
}
func TestGetTriplet0OutOf3OK(t *testing.T) {
	in := "123456789"
	expectedOut := "123"
	actualOut, _ := getTriplet(in, 0)
	if actualOut != expectedOut {
		t.Errorf(getTripletFailureString, 0, expectedOut, actualOut)
	}
}

func TestGetTriplet1OutOf3OK(t *testing.T) {
	in := "123456789"
	expectedOut := "456"
	actualOut, _ := getTriplet(in, 1)
	if actualOut != expectedOut {
		t.Errorf(getTripletFailureString, 1, expectedOut, actualOut)
	}
}

func TestGetTriplet2OutOf3OK(t *testing.T) {
	in := "123456789"
	expectedOut := "789"
	actualOut, _ := getTriplet(in, 2)
	if actualOut != expectedOut {
		t.Errorf(getTripletFailureString, 2, expectedOut, actualOut)
	}
}

func TestGetTriplet3OutOf3Error(t *testing.T) {
	in := "123456789"
	_, ok := getTriplet(in, 3)
	if ok {
		t.Error("Expected getTriplet(3) of 123456789 to return error")
	}
}
func TestGetTripletNegativeIndexError(t *testing.T) {
	in := "123456789"
	_, ok := getTriplet(in, -1)
	if ok {
		t.Error("Expected getTriplet(-1) of 123456789 to return error")
	}
}
func TestInsertCommaShortIntegerNoSign(t *testing.T) {
	in := "123"
	expectedOut := "123"
	actualOut := Comma(in)
	if actualOut != expectedOut {
		t.Errorf(commaFailureString, expectedOut, actualOut)
	}
}

func TestInsertCommaShortIntegerSign(t *testing.T) {
	in := "-123"
	expectedOut := "-123"
	actualOut := Comma(in)
	if actualOut != expectedOut {
		t.Errorf(commaFailureString, expectedOut, actualOut)
	}
}
func TestInsertCommaMediumIntegerNoSign(t *testing.T) {
	in := "123456"
	expectedOut := "123,456"
	actualOut := Comma(in)
	if actualOut != expectedOut {
		t.Errorf(commaFailureString, expectedOut, actualOut)
	}
}

func TestInsertCommaMediumIntegerSign(t *testing.T) {
	in := "-123456"
	expectedOut := "-123,456"
	actualOut := Comma(in)
	if actualOut != expectedOut {
		t.Errorf(commaFailureString, expectedOut, actualOut)
	}
}
func TestInsertCommaLongIntegerNoSign(t *testing.T) {
	in := "123456789"
	expectedOut := "123,456,789"
	actualOut := Comma(in)
	if actualOut != expectedOut {
		t.Errorf(commaFailureString, expectedOut, actualOut)
	}
}

func TestInsertCommaLongIntegerSign(t *testing.T) {
	in := "-123456789"
	expectedOut := "-123,456,789"
	actualOut := Comma(in)
	if actualOut != expectedOut {
		t.Errorf(commaFailureString, expectedOut, actualOut)
	}
}

func TestInsertCommaAnotherIntegerSign(t *testing.T) {
	in := "-12345"
	expectedOut := "-12,345"
	actualOut := Comma(in)
	if actualOut != expectedOut {
		t.Errorf(commaFailureString, expectedOut, actualOut)
	}
}
