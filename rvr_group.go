package metar

import "regexp"

const RVRPattern = `R(\d\d[RLC]?)\/([MP]?\d{4})(V([MP]?\d{4}))?FT`

var RVRRe *regexp.Regexp

type RVR struct {
	Runway     string
	Direction  string
	IsVariable bool
	Value      string
	Units      string
	Low        string
	High       string
}

func init() {
	RVRRe = regexp.MustCompile(RVRPattern)
}

func (r *Report) ParseRVRGroup() {
	rvr := &RVR{}

	if !RVRRe.MatchString(r.raw) {
		return
	}

	matches := RVRRe.FindAllStringSubmatch(r.raw, -1)
	rvr.Runway = matches[0][1]
	if len(rvr.Runway) == 3 {
		rvr.Direction = string(rvr.Runway[2])
		rvr.Runway = rvr.Runway[:2]
	}

	rvr.Value = matches[0][2]
	rvr.Units = "FT" //FIXME
	runwayLow := matches[0][2]
	runwayHigh := matches[0][4]

	if len(runwayHigh) != 0 {
		rvr.IsVariable = true
		rvr.Value = ""
		rvr.High = runwayHigh
		rvr.Low = runwayLow
	}

	r.RVR = rvr
}
