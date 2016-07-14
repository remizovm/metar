package metar

import "regexp"

const altimeterPattern = `A(\d{4})`

var altimeterRe *regexp.Regexp

func init() {
	altimeterRe = regexp.MustCompile(altimeterPattern)
}

func ParseAltimeter(m string) string {
	if !altimeterRe.MatchString(m) {
		return ""
	}
	return altimeterRe.FindAllStringSubmatch(m, -1)[0][1]
}
