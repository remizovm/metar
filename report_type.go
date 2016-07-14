package metar

import "strings"

func ParseReportType(m string) string {
	if strings.Contains(m, "SPECI") {
		return "SPECI"
	}
	return "METAR"
}
