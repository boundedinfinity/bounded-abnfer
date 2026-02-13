package abnfer

import (
	"math"

	"github.com/boundedinfinity/abnfer/char"
)

var (
	COREMAP = &RuleMap{
		m: map[string]Rule{},
	}

	ALPHA Rule = NewTerminal("ALPHA", 1, 1,
		char.Concat(
			char.Range(char.UPPERCASE_A, char.UPPERCASE_Z),
			char.Range(char.LOWERCASE_A, char.LOWERCASE_Z),
		),
	)

	BIT Rule = NewTerminal("BIT", 1, 1,
		char.AnyOf(char.ZERO, char.ONE),
	)

	CHAR Rule = NewTerminal("CHAR", 1, 1,
		char.Range(char.START_OF_HEADER, char.DELETE),
	)

	CR Rule = NewTerminal("CR", 1, 1,
		char.AnyOf(char.CARRIAGE_RETURN),
	)

	CRLF Rule = NewConcatenation("CRLF", 1, 1, []Rule{CR, LF})

	CTL Rule = NewAlternatives("CTL", 1, 1,
		[]Rule{
			NewTerminal("", 1, 1, char.Range(char.NULL, char.UNIT_SEPARATOR)),
			NewTerminal("", 1, 1, char.AnyOf(char.DELETE)),
		},
	)

	DIGIT = NewTerminal("DIGIT", 1, 1, char.Range(char.ZERO, char.NINE))

	DQUOTE = NewTerminal("DQUOTE", 1, 1, char.AnyOf(char.DOUBLE_QUOTE))

	HEXDIG = NewAlternatives("HEXDIG", 1, 1,
		[]Rule{
			DIGIT,
			NewTerminal("", 1, 1, char.AnyOf(
				char.UPPERCASE_A,
				char.UPPERCASE_B,
				char.UPPERCASE_C,
				char.UPPERCASE_D,
				char.UPPERCASE_E,
				char.UPPERCASE_F,
			)),
		},
	)

	HTAB Rule = NewTerminal("HTAB", 1, 1, char.AnyOf(char.HORIZONTAL_TAB))

	LF Rule = NewTerminal("LF", 1, 1, char.AnyOf(char.LINE_FEED))

	LWSP Rule = NewAlternatives("LWSP", 0, math.MaxInt, []Rule{
		WSP,
		NewConcatenation("", 0, math.MaxInt, []Rule{CRLF, WSP}),
	},
	)

	// OCTET Rule = NewTerminal("OCTET", 1, 1, char.Range(char.NULL, char.Char(0xFF)))

	SP Rule = NewTerminal("SP", 1, 1, char.AnyOf(char.SPACE))

	VCHAR Rule = NewTerminal("VCHAR", 1, 1, char.Range(char.EXCLAMATION_MARK, char.TILDE))

	WSP Rule = NewAlternatives("WSP", 1, 1, []Rule{SP, HTAB})
)

func init() {
	COREMAP.Append(
		ALPHA, BIT, CHAR, CR, CRLF, CTL, DIGIT, DQUOTE,
		HEXDIG, HTAB, LF, LWSP, SP, VCHAR, WSP,
	)
}
