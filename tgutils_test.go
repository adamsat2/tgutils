package tgutils

import "testing"

func TestUtils_RangedRandom(t *testing.T) {
	var testUtils Utils

	for i := 0; i < 100; i++ {
		val := testUtils.RangedRandom(1, 5)
		if val < 1 || val >= 5 {
			t.Errorf("RangedRandom returned out of range value: %d", val)
		}
	}
}

var strInSliceTests = []struct {
	str            string
	sl             []string
	expectedResult int
}{
	{str: "test", sl: []string{"test", "string"}, expectedResult: 0},
	{str: "not in slice", sl: []string{"test", "string"}, expectedResult: -1},
}

func TestUtils_StrInSlice(t *testing.T) {
	var testUtils Utils
	for i, e := range strInSliceTests {
		testResult := testUtils.StrInSlice(e.sl, e.str)
		if testResult != e.expectedResult {
			t.Errorf("Error in StrInSlice test %d, testResult = %d, expectedResult = %d", i, testResult, e.expectedResult)
		}
	}
}

var hasDigitTests = []struct {
	str            string
	expectedResult int
}{
	{str: "test", expectedResult: -1},
	{str: "test12", expectedResult: 4},
}

func TestUtils_HasDigit(t *testing.T) {
	var testUtils Utils
	for i, e := range hasDigitTests {
		testResult := testUtils.HasDigit(e.str)
		if testResult != e.expectedResult {
			t.Errorf("Error in HasDigit test %d, testResult = %d, expectedResult = %d", i, testResult, e.expectedResult)
		}
	}
}
