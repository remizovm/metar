package metar

import "regexp"

const dateTimePattern = `(\d{6})Z`

var dateTimeRe *regexp.Regexp

func init() {
	dateTimeRe = regexp.MustCompile(dateTimePattern)
}

func (r *Report) ParseDateTime() {
	if !dateTimeRe.MatchString(r.raw) {
		return
	}

	r.DateTime = dateTimeRe.FindAllStringSubmatch(r.raw, -1)[0][1]
}
