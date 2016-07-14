package metar

import (
	"log"
	"regexp"
	"strings"
)

const (
	windDirectionPattern             = `(\d{3})KT`
	windDirectionSpeedPattern        = `(\d{3})(\d{2,3})KT`
	windDirectionSpeedGustPattern    = `(\d{3})(\d{2,3})G(\d{2,3})KT`
	windDirectionVariablePattern     = `VRB(\d{2,3})KT`
	windDirectionVariableHighPattern = `((\d{3})(\d{2,3}))KT (\d{3})V(\d{3})`
	windDirectionFullPattern         = `((\d{3})(\d{2,3}))G(\d{2,3})KT (\d{3})V(\d{3})/gm`
	windCalmPattern                  = `00000KT`
)

var (
	windDirectionRe             *regexp.Regexp
	windDirectionSpeedRe        *regexp.Regexp
	windDirectionSpeedGustRe    *regexp.Regexp
	windDirectionVariableRe     *regexp.Regexp
	windDirectionVariableHighRe *regexp.Regexp
	windDirectionFullRe         *regexp.Regexp
)

func init() {
	windDirectionRe = regexp.MustCompile(windDirectionPattern)
	windDirectionSpeedRe = regexp.MustCompile(windDirectionSpeedPattern)
	windDirectionSpeedGustRe = regexp.MustCompile(windDirectionSpeedGustPattern)
	windDirectionVariableRe = regexp.MustCompile(windDirectionVariablePattern)
	windDirectionVariableHighRe = regexp.MustCompile(windDirectionVariableHighPattern)
	windDirectionFullRe = regexp.MustCompile(windDirectionFullPattern)
}

func ParseWindGroup(m string) {
	calmWind := false
	windDirectionVariable := false
	windDirection := ""
	windSpeed := ""
	windGust := ""
	windDirectionFrom := ""
	windDirectionTo := ""

	if strings.Contains(m, "00000KT") {
		calmWind = true
	} else {
		if windDirectionFullRe.MatchString(m) {
			matches := windDirectionVariableHighRe.FindAllStringSubmatch(m, -1)
			windDirection = matches[0][1]
			windSpeed = matches[0][2]
			windGust = matches[0][3]
			windDirectionVariable = true
			windDirectionFrom = matches[0][4]
			windDirectionTo = matches[0][5]
		} else if windDirectionVariableHighRe.MatchString(m) {
			matches := windDirectionVariableHighRe.FindAllStringSubmatch(m, -1)
			windDirection = matches[0][1]
			windSpeed = matches[0][2]
			windDirectionVariable = true
			windDirectionFrom = matches[0][3]
			windDirectionTo = matches[0][4]
		} else if windDirectionVariableRe.MatchString(m) {
			matches := windDirectionVariableRe.FindAllStringSubmatch(m, -1)
			windDirectionVariable = true
			windSpeed = matches[0][1]
		} else if windDirectionSpeedGustRe.MatchString(m) {
			matches := windDirectionSpeedGustRe.FindAllStringSubmatch(m, -1)
			windDirection = matches[0][1]
			windSpeed = matches[0][2]
			windGust = matches[0][3]
		} else if windDirectionSpeedRe.MatchString(m) {
			matches := windDirectionSpeedRe.FindAllStringSubmatch(m, -1)
			windDirection = matches[0][1]
			windSpeed = matches[0][2]
		} else if windDirectionRe.MatchString(m) {
			windDirection = windDirectionRe.FindAllStringSubmatch(m, -1)[0][1]
		}
	}
	log.Printf("Calm Wind: %v", calmWind)
	log.Printf("Wind Direction: %s", windDirection)
	log.Printf("Wind Speed: %s", windSpeed)
	log.Printf("Wind Gust: %s", windGust)
	log.Printf("Variable wind directioin: %v", windDirectionVariable)
	log.Printf("Wind directioin from: %s", windDirectionFrom)
	log.Printf("Wind directioin to: %s", windDirectionTo)
}
