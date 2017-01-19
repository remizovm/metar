package metar

import "testing"

const invalidReport = ""

func TestParse(t *testing.T) {
	result, err := Parse(testMetarReport)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Fail()
	}
}

func TestParseInvalidReport(t *testing.T) {
	result, err := Parse(invalidReport)
	if err != ErrInvalidReport {
		t.Error(err)
	}
	if result != nil {
		t.Fail()
	}
}
