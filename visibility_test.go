package metar

import "testing"

func TestParseVisibility(t *testing.T) {
	result, err := Parse(testMetarReport)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Fail()
	}
	if result.Visibility != 4 {
		t.Fail()
	}
}
