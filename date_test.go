package metar

import "testing"

const testMetarReportNoDate = `METAR KHQI AUTO 14006G12KT 4SM BR CLR 24/22 A2997
                         RMK A01`

func TestParseDate(t *testing.T) {
	result, err := Parse(testMetarReport)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Fail()
	}
	if result.Day != 18 { //182115
		t.Fail()
	}
	if result.Minute != 21 {
		t.Fail()
	}
	if result.Second != 15 {
		t.Fail()
	}
}

func TestParseDateFail(t *testing.T) {
	result, err := Parse(testMetarReportNoDate)
	if err != ErrDateNotFound {
		t.Error(err)
	}
	if result != nil {
		t.Fail()
	}

}
