package metar

import "regexp"

const remarksPattern = ` RMK (.+)`

var remarksRe *regexp.Regexp

func init() {
	remarksRe = regexp.MustCompile(remarksPattern)
}

func (r *Report) ParseRemarks() {
	if !remarksRe.MatchString(r.raw) {
		return
	}

	r.Remarks = remarksRe.FindAllStringSubmatch(r.raw, -1)[0][1]
}
