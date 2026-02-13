package abnfer

type AlternativesRule struct {
	Repetition
	Rules []Rule
}

func NewAlternatives(name string, min, max int, rs []Rule) Rule {
	return AlternativesRule{
		Repetition: Repetition{
			Name: name,
			Min:  min,
			Max:  max,
		},
		Rules: rs,
	}
}

func (r AlternativesRule) GetName() string {
	return r.Name
}

func (r AlternativesRule) New(min, max int) Rule {
	return NewAlternatives(r.Name, min, max, r.Rules)
}

func (r AlternativesRule) Matches(s string) MatchResult {
	for _, v := range r.Rules {
		if result := v.Matches(s); result.Matched {
			return result
		}
	}

	return MatchResult{}
}
