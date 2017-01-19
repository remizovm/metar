package metar

import (
	"testing"

	"github.com/Sirupsen/logrus"
)

const testMetarReportNoStationID = `METAR 182115Z AUTO 14006G12KT 4SM BR CLR
																		24/22 A2997 RMK A01`

func TestStationID(t *testing.T) {
	result, err := Parse(testMetarReport)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Fail()
	}

	if result.StationID != "KHQI" {
		logrus.Infof("%s %s", result.StationID, "KHQI")
		t.Fail()
	}
}

func TestNoStationID(t *testing.T) {
	result, err := Parse(testMetarReportNoStationID)
	if err != ErrStationIDNotFound {
		t.Error(err)
	}
	if result != nil {
		t.Fail()
	}
}
