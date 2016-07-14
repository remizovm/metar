package metar

import (
	"log"
	"regexp"
)

const (
	skyConditionPattern          = `([A-Z]{3})((\d{3})|\/\/\/)`
	skyVerticalVisibilityPattern = `VV(\d{3})`
	skyClearPattern              = `(SKC|CLR)`
)

var (
	skyConditionRe          *regexp.Regexp
	skyVerticalVisibilityRe *regexp.Regexp
	skyClearRe              *regexp.Regexp
)

func init() {
	skyConditionRe = regexp.MustCompile(skyConditionPattern)
	skyVerticalVisibilityRe = regexp.MustCompile(skyVerticalVisibilityPattern)
	skyClearRe = regexp.MustCompile(skyClearPattern)
}

func ParseSkyConditionGroup(m string) {
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
