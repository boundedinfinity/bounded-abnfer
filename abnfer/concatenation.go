package abnfer

type ConcatenationRule struct {
	Repetition
	Rules []Rule
}

func NewConcatenation(name string, min, max int, rs []Rule) Rule {
	return ConcatenationRule{
		Repetition: Repetition{
			Name: name,
			Min:  min,
			Max:  max,
		},
		Rules: rs,
	}
}

func (r ConcatenationRule) GetName() string {
	return r.Name
}

func (r ConcatenationRule) New(min, max int) Rule {
	return NewConcatenation(r.Name, min, max, r.Rules)
}

func (r ConcatenationRule) Matches(s string) MatchResult {
	var result MatchResult
	var mr1 []MatchResult
	s2 := s

	for {
		var mr2 []MatchResult

		for _, v := range r.Rules {
			mr := v.Matches(s2)

			if !mr.Matched {
				break
			} else {
				mr2 = append(mr2, mr)
				s2 = s2[mr.Count:]
			}
		}

		if len(mr2) != len(r.Rules) {
			break
		} else {
			mr1 = append(mr1, mr2...)
		}

		if s2 == "" {
			break
		}
	}

	dividend := len(mr1) / len(r.Rules)
	remainder := len(mr1) % len(r.Rules)

	if remainder == 0 && dividend >= r.Min && dividend <= r.Max {
		result.Matched = true

		for _, v := range mr1 {
			result.Count += v.Count
		}
	}

	if result.Count == 0 {
		result.Matched = false
	}

	return result
}
