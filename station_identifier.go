package metar

import "regexp"

const stationIdentifierPattern = `\s(\D{4})\s`

var stationIdentifierRe *regexp.Regexp

func init() {
	stationIdentifierRe = regexp.MustCompile(stationIdentifierPattern)
}

func ParseStationIdentifier(m string) string {
	if !stationIdentifierRe.MatchString(m) {
		return ""
	}
	return stationIdentifierRe.FindAllStringSubmatch(m, -1)[0][1]
}
