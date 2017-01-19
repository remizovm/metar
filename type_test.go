package metar

import "testing"

const (
	testMetarReport = `KHQI 182115Z AUTO 14006G12KT 4SM BR CLR 24/22 A2997
                         RMK A01`
	testSPECIReport = `SPECI VMMC 242341Z 10010KT 3500N VCSH FEW010 SCT018 BKN070
                     28/26 Q1004 TEMPO 5000 SHRA=`
)

func TestTypeMETAR(t *testing.T) {
	result, err := Parse(testMetarReport)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Fail()
	}
	if result.Type != "METAR" {
		t.Fail()
	}
}

func TestTypeSPECI(t *testing.T) {
	result, err := Parse(testSPECIReport)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Fail()
	}
	if result.Type != "SPECI" {
		t.Fail()
	}
}
