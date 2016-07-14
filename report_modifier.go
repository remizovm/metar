package metar

import "regexp"

const reportModifierPattern = `(AUTO|COR)`

var reportModifierRe *regexp.Regexp

func init() {
	reportModifierRe = regexp.MustCompile(reportModifierPattern)
}

func (r *Report) ParseReportModifier() {
	if !reportModifierRe.MatchString(r.raw) {
		return
	}

	r.ReportModifier = reportModifierRe.FindAllStringSubmatch(r.raw, -1)[0][1]
}
