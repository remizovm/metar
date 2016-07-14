package metar

import "strings"

const (
	speci = "SPECI"
	metar = "METAR"
)

func (r *Report) ParseReportType() {
	if strings.Contains(r.raw, speci) {
		r.ReportType = speci
	} else {
		r.ReportType = metar
	}
}
