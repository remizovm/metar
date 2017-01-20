package metar

import "testing"

const (
	testMetarRVRReport = `METAR KHQI 182115Z AUTO 14006G12KT 4SM BR CLR 24/22 A2997
                         RMK A01 R01L/0600FT`
	testMetarVaryingRVRReport = `METAR KHQI 182115Z AUTO 14006G12KT 4SM BR CLR 24/22 A2997
                         RMK A01 R01L/0600V0800FT`
)

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
	if result.RunwayDirection != "L" {
		t.Fail()
	}
	if result.RunwayValue != 600 {
		t.Fail()
	}
	if result.IsVaryingRVRValue {
		t.Fail()
	}
}

func TestParseRVRGroupVarying(t *testing.T) {
	result, err := Parse(testMetarVaryingRVRReport)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Fail()
	}
	if result.RunwayNumber != 1 {
		t.Fail()
	}
	if !result.IsVaryingRVRValue {
		t.Fail()
	}
	if result.RunwayValueFrom != 600 {
		t.Fail()
	}
	if result.RunwayValueTo != 800 {
		t.Fail()
	}
}
