package metar

import "regexp"

const temperatureDewPointPattern = `(M?\d\d)\/(M?\d\d)`

var temperatureDewPointRe *regexp.Regexp

func init() {
	temperatureDewPointRe = regexp.MustCompile(temperatureDewPointPattern)
}

func (r *Report) ParseTemperatureDewPointGroup() {
	if !temperatureDewPointRe.MatchString(r.raw) {
		return
	}

	matches := temperatureDewPointRe.FindAllStringSubmatch(r.raw, -1)

	r.Temperature = matches[0][1]
	r.DewPoint = matches[0][2]
}
