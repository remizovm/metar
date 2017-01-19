package metar

import (
	"errors"
	"regexp"
	"strings"
)

var ErrInvalidReport = errors.New("Invalid report")

type Report struct {
	Type string
}

var typeRe *regexp.Regexp

func init() {
	typeRe = regexp.MustCompile("SPECI\\s")
}

func Parse(raw string) (*Report, error) {
	chunks := strings.Split(raw, " ")
	if len(chunks) == 1 {
		return nil, ErrInvalidReport
	}
	m := &Report{
		Type: "METAR",
	}
	if typeRe.MatchString(raw) {
		m.Type = "SPECI"
	}
	return m, nil
}
