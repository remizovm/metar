package metar

import "regexp"

const dateTimePattern = `(\d{6})Z`

var dateTimeRe *regexp.Regexp

func init() {
	dateTimeRe = regexp.MustCompile(dateTimePattern)
}

func ParseDateTime(m string) string {
	if !dateTimeRe.MatchString(m) {
		return ""
	}
	return dateTimeRe.FindAllStringSubmatch(m, -1)[0][1]
}
