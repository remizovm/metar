package metar

type Report struct {
	ReportType            string
	StationID             string
	DateTime              string
	ReportModifier        string
	Wind                  *Wind
	Visibility            string
	IsVisibilityAutomated bool
	RVR                   *RVR
	RawWeatherChunks      []string
	SkyCondition          *SkyCondition
	Temperature           string
	DewPoint              string
	Altimeter             string
	Remarks               string
	raw                   string
}

func GetReport(raw string) *Report {
	r := &Report{raw: raw}

	r.ParseReportType()
	r.ParseStationID()
	r.ParseDateTime()

	return r
}
