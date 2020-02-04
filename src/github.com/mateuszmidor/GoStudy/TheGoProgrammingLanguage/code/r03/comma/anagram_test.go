package stringgames

import "testing"

const ranagramsFailureString = "Expected %q and %q to be anagrams"

func TestAnagramEmpty(t *testing.T) {
	s1 := ""
	s2 := ""
	anagrams := Anagrams(s1, s2)
	if anagrams == false {
		t.Errorf(ranagramsFailureString, s1, s2)
	}
}

func TestAnagramSingleChar(t *testing.T) {
	s1 := "e"
	s2 := "e"
	anagrams := Anagrams(s1, s2)
	if anagrams == false {
		t.Errorf(ranagramsFailureString, s1, s2)
	}
}

func TestAnagramManyChar1(t *testing.T) {
	s1 := "adam"
	s2 := "dama"
	anagrams := Anagrams(s1, s2)
	if anagrams == false {
		t.Errorf(ranagramsFailureString, s1, s2)
	}
}

func TestAnagramManyChar2(t *testing.T) {
	s1 := "KORBA"
	s2 := "ROBAK"
	anagrams := Anagrams(s1, s2)
	if anagrams == false {
		t.Errorf(ranagramsFailureString, s1, s2)
	}
}
