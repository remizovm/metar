package metar

import "regexp"

const altimeterPattern = `A(\d{4})`

var altimeterRe *regexp.Regexp

func init() {
	altimeterRe = regexp.MustCompile(altimeterPattern)
}

func (r *Report) ParseAltimeter() {
	if !altimeterRe.MatchString(r.raw) {
		return
	}

	r.Altimeter = altimeterRe.FindAllStringSubmatch(r.raw, -1)[0][1]
}
