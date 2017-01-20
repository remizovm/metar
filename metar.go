package metar

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInvalidReport     = errors.New("Invalid report")
	ErrTypeNotFound      = errors.New("Type of the report could not be determined")
	ErrStationIDNotFound = errors.New("Could not find station identifier")
	ErrDateNotFound      = errors.New("Could not find timestamp")
)

const (
	ReportModifierAUTO = "AUTO"
	ReportModifierCOR  = "COR"
)

type Report struct {
	Type          string
	StationID     string
	Day           int
	Minute        int
	Second        int
	Modifier      string
	WindDirection int
	WindSpeed     int
	WindGust      int
	Visibility    int

	chunks []string
	raw    string
}

func Parse(raw string) (*Report, error) {
	chunks := strings.Split(raw, " ")
	if len(chunks) == 1 {
		return nil, ErrInvalidReport
	}

	r := &Report{
		Type:   "METAR",
		chunks: chunks,
		raw:    raw,
	}

	if err := r.parseStatus(); err != nil {
		return nil, err
	}

	if err := r.parseStationID(); err != nil {
		return nil, err
	}

	if err := r.parseDate(); err != nil {
		return nil, err
	}

	r.parseModifier()

	if err := r.parseWindGroup(); err != nil {
		return nil, err
	}

	if err := r.parseVisibilityGroup(); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Report) parseVisibilityGroup() error {
	re := regexp.MustCompile("M{0,1}([\\d\\s]{1,6})SM")
	if !re.MatchString(r.raw) {
		return nil
	}
	matches := re.FindStringSubmatch(r.raw)
	if len(matches) <= 1 {
		return nil
	}
	visibilityStr := strings.Trim(matches[1], " ")
	visibility, err := strconv.Atoi(visibilityStr)
	if err != nil {
		return err
	}
	r.Visibility = visibility
	return nil
}

func (r *Report) parseWindGroup() error {
	/*
		TODO:
		- Variable wind direction
		- Calm wind
	*/
	re := regexp.MustCompile("(\\d{3})(\\d{2,3})G(\\d{2,3})KT")
	if !re.MatchString(r.raw) {
		return nil
	}
	matches := re.FindStringSubmatch(r.raw)
	if len(matches) <= 1 {
		return nil
	}
	switch len(matches) - 1 {
	case 3:
		// Parsing wind direction
		windDirectionStr := matches[1]
		windDirection, err := strconv.Atoi(windDirectionStr)
		if err != nil {
			return err
		}
		r.WindDirection = windDirection
		// Parsing wind speed
		windSpeedStr := matches[2]
		windSpeed, err := strconv.Atoi(windSpeedStr)
		if err != nil {
			return err
		}
		r.WindSpeed = windSpeed
		// Parsing wind gust
		windGustStr := matches[3]
		windGust, err := strconv.Atoi(windGustStr)
		if err != nil {
			return err
		}
		r.WindGust = windGust
		return nil

	}
	return nil
}

func (r *Report) parseModifier() {
	for i := range r.chunks {
		if r.chunks[i] == ReportModifierAUTO {
			r.Modifier = ReportModifierAUTO
			r.chunks = removeIndex(r.chunks, i)
			return
		} else if r.chunks[i] == ReportModifierCOR {
			r.Modifier = ReportModifierCOR
			r.chunks = removeIndex(r.chunks, i)
			return
		}
	}
}

func (r *Report) parseDate() error {
	re := regexp.MustCompile("(\\d\\d)(\\d\\d)(\\d\\d)Z")
	if !re.MatchString(r.raw) {
		return ErrDateNotFound
	}

	matches := re.FindStringSubmatch(r.raw)
	if len(matches) != 4 {
		return ErrDateNotFound
	}

	day, err := strconv.Atoi(matches[1])
	if err != nil {
		return err
	}
	r.Day = day

	minute, err := strconv.Atoi(matches[2])
	if err != nil {
		return err
	}
	r.Minute = minute

	second, err := strconv.Atoi(matches[3])
	if err != nil {
		return err
	}
	r.Second = second

	return nil
}

func (r *Report) parseStationID() error {
	re := regexp.MustCompile("\\s([A-Z]{4})\\s")
	if !re.MatchString(r.raw) {
		return ErrStationIDNotFound
	}
	matches := re.FindAllStringSubmatch(r.raw, -1)
	if len(matches) == 0 {
		return ErrStationIDNotFound
	}
	if matches[0][1] != ReportModifierAUTO {
		r.StationID = matches[0][1]
		return nil
	}
	return ErrStationIDNotFound
}

func (r *Report) parseStatus() error {
	for i := range r.chunks {
		if r.chunks[i] == "METAR" {
			r.Type = "METAR"
			r.chunks = removeIndex(r.chunks, i)
			return nil
		} else if r.chunks[i] == "SPECI" {
			r.Type = "SPECI"
			r.chunks = removeIndex(r.chunks, i)
			return nil
		}
	}

	return ErrTypeNotFound
}

func removeIndex(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
