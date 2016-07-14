package metar

import "regexp"

const remarksPattern = ` RMK (.+)`

var remarksRe *regexp.Regexp

func init() {
	remarksRe = regexp.MustCompile(remarksPattern)
}

func parseRemarks(m string) string {
	if !remarksRe.MatchString(m) {
		return ""
	}
	return remarksRe.FindAllStringSubmatch(m, -1)[0][1]
}
