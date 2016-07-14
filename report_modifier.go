package metar

import "regexp"

const reportModifierPattern = `(AUTO|COR)`

var reportModifierRe *regexp.Regexp

func init() {
	reportModifierRe = regexp.MustCompile(reportModifierPattern)
}

func ParseReportModifier(m string) string {
	if !reportModifierRe.MatchString(m) {
		return ""
	}
	return reportModifierRe.FindAllStringSubmatch(m, -1)[0][1]
}
