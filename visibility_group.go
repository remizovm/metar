package metar

import "regexp"

const visibilityPattern = ` (.{1,5})SM `

var visibilityRe *regexp.Regexp

func init() {
	visibilityRe = regexp.MustCompile(visibilityPattern)
}

func (r *Report) ParseVisibilityGroup() {
	if !visibilityRe.MatchString(r.raw) {
		return
	}

	r.Visibility = visibilityRe.FindAllStringSubmatch(r.raw, -1)[0][1]

	if r.Visibility[0] == 'M' {
		r.IsVisibilityAutomated = true
		r.Visibility = r.Visibility[1:]
	}
}
