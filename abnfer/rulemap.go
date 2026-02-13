package abnfer

type RuleMap struct {
	m map[string]Rule
}

func (t RuleMap) Exist(r Rule) bool {
	if r.GetName() == "" {
		return false
	}
	_, ok := t.m[r.GetName()]
	return ok
}

func (t *RuleMap) Append(rs ...Rule) {
	for _, r := range rs {
		if r.GetName() != "" && !t.Exist(r) {
			t.m[r.GetName()] = r
		}
	}
}

func (t RuleMap) Clone(r Rule) *RuleMap {
	clone := RuleMap{
		m: make(map[string]Rule),
	}

	for k, v := range t.m {
		clone.m[k] = v
	}

	return &clone
}
