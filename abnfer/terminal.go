package abnfer

import (
	"strings"

	"github.com/boundedinfinity/abnfer/char"
)

type TerminalRule struct {
	Repetition
	Terminals []char.Char
}

func NewTerminal(name string, min, max int, cs []char.Char) Rule {
	return TerminalRule{
		Repetition: Repetition{
			Name: name,
			Min:  min,
			Max:  max,
		},
		Terminals: cs,
	}
}

func (r TerminalRule) GetName() string {
	return r.Name
}

func (r TerminalRule) New(min, max int) Rule {
	return NewTerminal(r.Name, min, max, r.Terminals)
}

func (r TerminalRule) Matches(s string) MatchResult {
	var result MatchResult

	if len(s) < r.Min {
		return result
	}

	b := make([]bool, 0)

	for i := 0; i < len(s); i++ {
		var found bool

		for _, v := range r.Terminals {
			if strings.HasPrefix(s[i:], v.String()) {
				b = append(b, true)
				found = true
				break
			}
		}

		if !found {
			break
		}
	}

	bl := len(b)

	if bl >= r.Min && bl <= r.Max {
		result.Matched = true
		result.Count = bl
	}

	if result.Count == 0 {
		result.Matched = false
	}

	return result
}
