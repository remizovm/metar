package metar

import "regexp"

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

type SkyCondition struct {
	SkyCover           string
	SkyLayerHeight     string
	VerticalVisibility string
	IsClear            bool
}

func init() {
	skyConditionRe = regexp.MustCompile(skyConditionPattern)
	skyVerticalVisibilityRe = regexp.MustCompile(skyVerticalVisibilityPattern)
	skyClearRe = regexp.MustCompile(skyClearPattern)
}

func (sc *SkyCondition) condition(m string) {
	if !skyConditionRe.MatchString(m) {
		return
	}
	matches := skyConditionRe.FindAllStringSubmatch(m, -1)
	sc.SkyCover = matches[0][1]
	sc.SkyLayerHeight = matches[0][2]
}

func (sc *SkyCondition) verticalVisibility(m string) {
	if !skyVerticalVisibilityRe.MatchString(m) {
		return
	}
	sc.VerticalVisibility = skyVerticalVisibilityRe.FindAllStringSubmatch(m, -1)[0][1]
}

func (r *Report) ParseSkyConditionGroup() {
	sc := &SkyCondition{}

	sc.condition(r.raw)
	sc.verticalVisibility(r.raw)

	if skyClearRe.MatchString(r.raw) {
		sc.IsClear = true
	}

	r.SkyCondition = sc
}
