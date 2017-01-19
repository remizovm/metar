package metar

import (
	"errors"
	"strings"
)

var (
	ErrInvalidReport = errors.New("Invalid report")
	ErrTypeNotFound  = errors.New("Type of the report could not be determined")
)

type Report struct {
	Type   string
	chunks []string
}

func Parse(raw string) (*Report, error) {
	chunks := strings.Split(raw, " ")
	if len(chunks) == 1 {
		return nil, ErrInvalidReport
	}

	r := &Report{
		Type:   "METAR",
		chunks: chunks,
	}

	if err := r.parseStatus(); err != nil {
		return nil, err
	}

	return r, nil
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
