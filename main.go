package main

import (
	"log"
	"regexp"
	"strings"
)

const (
	metarSample                      = "R01L/0800FT METAR CYYT 131000Z 33017G27KT 4123SM -DZ BR OVC006 08/07 A2991 RMK ST8 SLP135 RMK AO2"
	stationIdentifierPattern         = `\s(\D{4})\s`
	dateTimePattern                  = `(\d{6})Z`
	reportModifierPattern            = `(AUTO|COR)`
	windDirectionPattern             = `(\d{3})KT`
	windDirectionSpeedPattern        = `(\d{3})(\d{2,3})KT`
	windDirectionSpeedGustPattern    = `(\d{3})(\d{2,3})G(\d{2,3})KT`
	windDirectionVariablePattern     = `VRB(\d{2,3})KT`
	windDirectionVariableHighPattern = `((\d{3})(\d{2,3}))KT (\d{3})V(\d{3})`
	windDirectionFullPattern         = `((\d{3})(\d{2,3}))G(\d{2,3})KT (\d{3})V(\d{3})/gm`
	windCalmPattern                  = `00000KT`
	visibilityPattern                = ` (.{1,5})SM `
	presentWeatherGroupPattern       = `(\+|-|VC)?((MI|PR|BC|DR|BL|SH|TS|FZ)|(DZ|RA|SN|SG|IC|PL|GR|GS|UP)|(BR|FG|FU|VA|DU|SA|HZ|PY)|(PO|SQ|FC|SS|DS))`
	skyConditionPattern              = `([A-Z]{3})((\d{3})|\/\/\/)`
	skyVerticalVisibilityPattern     = `VV(\d{3})`
	skyClearPattern                  = `(SKC|CLR)`
	temperatureDewPointPattern       = `(M?\d\d)\/(M?\d\d)`
	altimeterPattern                 = `A(\d{4})`
	RVRPattern                       = `R(\d\d[RLC]?)\/([MP]?\d{4})(V([MP]?\d{4}))?FT`
	remarksPattern                   = ` RMK (.+)`
)

var stationIdentifierRe *regexp.Regexp
var dateTimeRe *regexp.Regexp
var reportModifierRe *regexp.Regexp
var windDirectionRe *regexp.Regexp
var windDirectionSpeedRe *regexp.Regexp
var windDirectionSpeedGustRe *regexp.Regexp
var windDirectionVariableRe *regexp.Regexp
var windDirectionVariableHighRe *regexp.Regexp
var windDirectionFullRe *regexp.Regexp
var visibilityRe *regexp.Regexp
var presentWeatherGroupRe *regexp.Regexp
var skyConditionRe *regexp.Regexp
var skyVerticalVisibilityRe *regexp.Regexp
var skyClearRe *regexp.Regexp
var temperatureDewPointRe *regexp.Regexp
var altimeterRe *regexp.Regexp
var RVRRe *regexp.Regexp
var remarksRe *regexp.Regexp

func init() {
	stationIdentifierRe = regexp.MustCompile(stationIdentifierPattern)
	dateTimeRe = regexp.MustCompile(dateTimePattern)
	reportModifierRe = regexp.MustCompile(reportModifierPattern)
	windDirectionRe = regexp.MustCompile(windDirectionPattern)
	windDirectionSpeedRe = regexp.MustCompile(windDirectionSpeedPattern)
	windDirectionSpeedGustRe = regexp.MustCompile(windDirectionSpeedGustPattern)
	windDirectionVariableRe = regexp.MustCompile(windDirectionVariablePattern)
	windDirectionVariableHighRe = regexp.MustCompile(windDirectionVariableHighPattern)
	windDirectionFullRe = regexp.MustCompile(windDirectionFullPattern)
	visibilityRe = regexp.MustCompile(visibilityPattern)
	presentWeatherGroupRe = regexp.MustCompile(presentWeatherGroupPattern)
	skyConditionRe = regexp.MustCompile(skyConditionPattern)
	skyVerticalVisibilityRe = regexp.MustCompile(skyVerticalVisibilityPattern)
	skyClearRe = regexp.MustCompile(skyClearPattern)
	temperatureDewPointRe = regexp.MustCompile(temperatureDewPointPattern)
	altimeterRe = regexp.MustCompile(altimeterPattern)
	RVRRe = regexp.MustCompile(RVRPattern)
	remarksRe = regexp.MustCompile(remarksPattern)
}

func main() {
	log.Printf("Raw Report: %s", metarSample)
	parseReportType(metarSample)
	parseStationIdentifier(metarSample)
	parseDateTime(metarSample)
	parseReportModifier(metarSample)
	parseWindGroup(metarSample)
	parseVisibilityGroup(metarSample)
	parseRVRGroup(metarSample)
	parsePresentWeatherGroup(metarSample)
	parseSkyConditionGroup(metarSample)
	parseTemperatureDewPointGroup(metarSample)
	parseAltimeter(metarSample)
	remarks := parseRemarks(metarSample)
	log.Printf("Remarks: %s", remarks)
}

func parseRemarks(m string) string {
	if !remarksRe.MatchString(m) {
		return ""
	}
	return remarksRe.FindAllStringSubmatch(m, -1)[0][1]
}

func parseRVRGroup(m string) {
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

func parseAltimeter(m string) {
	if !altimeterRe.MatchString(m) {
		return
	}
	altimeter := altimeterRe.FindAllStringSubmatch(m, -1)[0][1]
	log.Printf("Altimeter: %s", altimeter)
}

func parseTemperatureDewPointGroup(m string) {
	if !temperatureDewPointRe.MatchString(m) {
		return
	}
	matches := temperatureDewPointRe.FindAllStringSubmatch(m, -1)
	temperature := matches[0][1]
	dewPoint := matches[0][2]
	log.Printf("Temperature: %s", temperature)
	log.Printf("Dew Point: %s", dewPoint)
}

func parseSkyConditionGroup(m string) {
	cond := func(m string) {
		skyCover := ""
		skyLayerHeight := ""
		if !skyConditionRe.MatchString(m) {
			return
		}
		matches := skyConditionRe.FindAllStringSubmatch(m, -1)
		skyCover = matches[0][1]
		skyLayerHeight = matches[0][2]
		log.Printf("Sky Cover: %s", skyCover)
		log.Printf("Sky Layer Height: %s", skyLayerHeight)
	}

	vv := func(m string) {
		verticalVisibility := ""
		if !skyVerticalVisibilityRe.MatchString(m) {
			return
		}
		matches := skyVerticalVisibilityRe.FindAllStringSubmatch(m, -1)
		verticalVisibility = matches[0][1]
		log.Printf("Vertical Visibility: %s", verticalVisibility)
	}

	sc := func(m string) {
		skyClear := ""
		if !skyClearRe.MatchString(m) {
			return
		}
		matches := skyClearRe.FindAllStringSubmatch(m, -1)
		skyClear = matches[0][1]
		log.Printf("Sky Clear: %s", skyClear)
	}

	cond(m)
	vv(m)
	sc(m)
}

func parsePresentWeatherGroup(m string) {
	var rawWeatherChunks []string
	matches := presentWeatherGroupRe.FindAllStringSubmatch(m, -1)
	for _, match := range matches {
		if len(match) != 0 {
			rawWeatherChunks = append(rawWeatherChunks, match[0])
		}
	}
	log.Printf("Raw present weather chunks: %v", rawWeatherChunks)
}

func parseVisibilityGroup(m string) {
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

func parseWindGroup(m string) {
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

func parseReportModifier(m string) {
	var result string
	if !reportModifierRe.MatchString(m) {
		return
	}
	result = reportModifierRe.FindAllStringSubmatch(m, -1)[0][1]
	log.Printf("Report Modifier: %s", result)
}

func parseDateTime(m string) {
	var result string
	if !dateTimeRe.MatchString(m) {
		return
	}
	result = dateTimeRe.FindAllStringSubmatch(m, -1)[0][1]
	log.Printf("Date and Time of the Report: %s", result)
}

func parseStationIdentifier(m string) {
	var result string
	match := stationIdentifierRe.FindAllStringSubmatch(m, -1)
	if len(match) != 0 {
		result = match[0][1]
	}
	log.Printf("Station Identifier: %s", result)
}

func parseReportType(m string) {
	result := "METAR"
	if strings.Contains(m, "SPECI") {
		result = "SPECI"
	}
	log.Printf("Type of Report: %s", result)
}
