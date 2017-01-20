package metar

import "testing"

func TestParsePresentWeather(t *testing.T) {
	result, err := Parse(testMetarReport)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Fail()
	}
	if result.PresentWeatherRaw != "BR" {
		t.Fail()
	}
}
