package metar

import "regexp"

const presentWeatherGroupPattern = `(\+|-|VC)?((MI|PR|BC|DR|BL|SH|TS|FZ)|(DZ|RA|SN|SG|IC|PL|GR|GS|UP)|(BR|FG|FU|VA|DU|SA|HZ|PY)|(PO|SQ|FC|SS|DS))`

var presentWeatherGroupRe *regexp.Regexp

func init() {
	presentWeatherGroupRe = regexp.MustCompile(presentWeatherGroupPattern)
}

func ParsePresentWeatherGroup(m string) []string {
	var rawWeatherChunks []string
	matches := presentWeatherGroupRe.FindAllStringSubmatch(m, -1)
	for _, match := range matches {
		rawWeatherChunks = append(rawWeatherChunks, match[0])

	}
	return rawWeatherChunks
}
