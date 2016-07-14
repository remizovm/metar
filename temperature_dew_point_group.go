package metar

import (
	"log"
	"regexp"
)

const temperatureDewPointPattern = `(M?\d\d)\/(M?\d\d)`

var temperatureDewPointRe *regexp.Regexp

func init() {
	temperatureDewPointRe = regexp.MustCompile(temperatureDewPointPattern)
}

func ParseTemperatureDewPointGroup(m string) {
	if !temperatureDewPointRe.MatchString(m) {
		return
	}
	matches := temperatureDewPointRe.FindAllStringSubmatch(m, -1)
	temperature := matches[0][1]
	dewPoint := matches[0][2]
	log.Printf("Temperature: %s", temperature)
	log.Printf("Dew Point: %s", dewPoint)
}
