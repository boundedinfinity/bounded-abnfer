package abnfer

type Rule interface {
	Matches(string) MatchResult
	New(int, int) Rule
	GetName() string
}

type Repetition struct {
	Name string
	Min  int
	Max  int
}

type MatchResult struct {
	Count   int
	Matched bool
}
