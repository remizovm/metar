package metar

import (
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

type Wind struct {
	IsCalm              bool
	IsVariableDirection bool
	Direction           string
	Speed               string
	Gust                string
	DirectionFrom       string
	DirectionTo         string
}

func (r *Report) ParseWindGroup() {
	w := &Wind{}

	if strings.Contains(r.raw, "00000KT") {
		w.IsCalm = true
	} else {
		if windDirectionFullRe.MatchString(r.raw) {
			matches := windDirectionVariableHighRe.FindAllStringSubmatch(r.raw, -1)
			w.Direction = matches[0][1]
			w.Speed = matches[0][2]
			w.Gust = matches[0][3]
			w.IsVariableDirection = true
			w.DirectionFrom = matches[0][4]
			w.DirectionTo = matches[0][5]
		} else if windDirectionVariableHighRe.MatchString(r.raw) {
			matches := windDirectionVariableHighRe.FindAllStringSubmatch(r.raw, -1)
			w.Direction = matches[0][1]
			w.Speed = matches[0][2]
			w.IsVariableDirection = true
			w.DirectionFrom = matches[0][3]
			w.DirectionTo = matches[0][4]
		} else if windDirectionVariableRe.MatchString(r.raw) {
			matches := windDirectionVariableRe.FindAllStringSubmatch(r.raw, -1)
			w.IsVariableDirection = true
			w.Speed = matches[0][1]
		} else if windDirectionSpeedGustRe.MatchString(r.raw) {
			matches := windDirectionSpeedGustRe.FindAllStringSubmatch(r.raw, -1)
			w.Direction = matches[0][1]
			w.Speed = matches[0][2]
			w.Gust = matches[0][3]
		} else if windDirectionSpeedRe.MatchString(r.raw) {
			matches := windDirectionSpeedRe.FindAllStringSubmatch(r.raw, -1)
			w.Direction = matches[0][1]
			w.Speed = matches[0][2]
		} else if windDirectionRe.MatchString(r.raw) {
			w.Direction = windDirectionRe.FindAllStringSubmatch(r.raw, -1)[0][1]
		}
	}

	r.Wind = w
}
