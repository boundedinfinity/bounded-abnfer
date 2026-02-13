package abnfer

func NewMap() *RuleMap {
	return &RuleMap{
		m: map[string]Rule{},
	}
}

func CoreMap() *RuleMap {
	rm := NewMap()

	rm.Append(ALPHA, CHAR)

	return rm
}
