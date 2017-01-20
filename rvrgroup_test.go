package metar

import "testing"

const testMetarRVRReport = `METAR KHQI 182115Z AUTO 14006G12KT 4SM BR CLR 24/22 A2997
                         RMK A01 R01L/0600FT`

func TestParseRVRGroup(t *testing.T) {
	result, err := Parse(testMetarRVRReport)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Fail()
	}
	if result.RunwayNumber != 1 {
		t.Fail()
	}
}
