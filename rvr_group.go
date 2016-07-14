package metar

import (
	"log"
	"regexp"
)

const RVRPattern = `R(\d\d[RLC]?)\/([MP]?\d{4})(V([MP]?\d{4}))?FT`

var RVRRe *regexp.Regexp

func init() {
	RVRRe = regexp.MustCompile(RVRPattern)
}

func ParseRVRGroup(m string) {
	if !RVRRe.MatchString(m) {
		return
	}
	matches := RVRRe.FindAllStringSubmatch(m, -1)
	runway := matches[0][1]
	runwayDirection := "undefined"
	runwayVariable := false
	if len(runway) == 3 {
		runwayDirection = string(runway[2])
		runway = runway[:2]
	}
	runwayValue := matches[0][2]
	runwayUnits := "FT" //FIXME
	runwayLow := matches[0][2]
	runwayHigh := matches[0][4]
	if len(runwayHigh) != 0 {
		runwayVariable = true
	}
	log.Printf("Runway: %s", runway)
	log.Printf("Runway Direction: %s", runwayDirection)
	if runwayVariable {
		log.Printf("Runway Low: %s", runwayLow)
		log.Printf("Runway High: %s", runwayHigh)
	} else {
		log.Printf("Runway Value: %s", runwayValue)
	}
	log.Printf("Runway Units: %s", runwayUnits)
}
