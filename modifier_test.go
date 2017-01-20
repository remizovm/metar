package metar

import "testing"

const testMetarReportCOR = `METAR KHQI 182115Z COR 14006G12KT 4SM BR CLR 24/22 A2997
                         RMK A01`

func TestModifierAuto(t *testing.T) {
	result, err := Parse(testMetarReport)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Fail()
	}

	if result.Modifier != ReportModifierAUTO {
		t.Fail()
	}
}

func TestModifierCOR(t *testing.T) {
	result, err := Parse(testMetarReportCOR)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Fail()
	}

	if result.Modifier != ReportModifierCOR {
		t.Fail()
	}
}
