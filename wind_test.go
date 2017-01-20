package metar

import "testing"

func TestParseWindGroup(t *testing.T) {
	result, err := Parse(testMetarReport)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Fail()
	}

	if result.WindDirection != 140 {
		t.Fail()
	}

	if result.WindSpeed != 6 {
		t.Fail()
	}

	if result.WindGust != 12 {
		t.Fail()
	}
}
