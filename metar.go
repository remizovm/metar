package metar

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidReport     = errors.New("Invalid report")
	ErrTypeNotFound      = errors.New("Type of the report could not be determined")
	ErrStationIDNotFound = errors.New("Could not find station identifier")
)

type Report struct {
	Type      string
	StationID string

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

	return r, nil
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
	if matches[0][1] != "AUTO" {
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
