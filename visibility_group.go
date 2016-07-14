package metar

import (
	"log"
	"regexp"
)

const visibilityPattern = ` (.{1,5})SM `

var visibilityRe *regexp.Regexp

func init() {
	visibilityRe = regexp.MustCompile(visibilityPattern)
}

func ParseVisibilityGroup(m string) {
	if !visibilityRe.MatchString(m) {
		return
	}
	surfaceVisibility := ""
	isAutomated := false
	rawVisibilityValue := visibilityRe.FindAllStringSubmatch(m, -1)[0][1]
	if rawVisibilityValue[0] == 'M' {
		isAutomated = true
		surfaceVisibility = rawVisibilityValue[1:]
	} else {
		surfaceVisibility = rawVisibilityValue
	}

	if isAutomated {
		log.Printf("Surface visibility is less than: %s", surfaceVisibility)
	} else {
		log.Printf("Surface visibility: %s", surfaceVisibility)
	}
}
