package metar

import "regexp"

const stationIDPattern = `\s(\D{4})\s`

var stationIDRe *regexp.Regexp

func init() {
	stationIDRe = regexp.MustCompile(stationIDPattern)
}

func (r *Report) ParseStationID() {
	if !stationIDRe.MatchString(r.raw) {
		return
	}

	r.StationID = stationIDRe.FindAllStringSubmatch(r.raw, -1)[0][1]
}
