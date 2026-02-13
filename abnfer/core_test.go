package abnfer_test

import (
	"testing"

	"github.com/boundedinfinity/abnfer/abnfer"
	"github.com/boundedinfinity/abnfer/abnfer/char"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRegex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lang Suite")
}

var _ = Describe("Rule", func() {
	Context("CRLF", func() {
		It("should match \\r\\n", func() {
			expected := abnfer.MatchResult{2, true}
			actual := abnfer.CRLF.Matches("\r\n")
			Expect(actual).To(Equal(expected))
		})

		It("should match \\r\\n", func() {
			expected := abnfer.MatchResult{4, true}
			actual := abnfer.CRLF.New(1, 2).Matches("\r\n\r\n")
			Expect(actual).To(Equal(expected))
		})

		It("should not match 'BELL'", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.CRLF.Matches("\b")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("ALPHA", func() {
		It("should match A", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.ALPHA.Matches("A")
			Expect(actual).To(Equal(expected))
		})

		It("should match a", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.ALPHA.Matches("a")
			Expect(actual).To(Equal(expected))
		})

		It("should not match !", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.ALPHA.Matches("!")
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.ALPHA.New(2, 2).Matches("A")
			Expect(actual).To(Equal(expected))
		})

		It("should match AA", func() {
			expected := abnfer.MatchResult{2, true}
			actual := abnfer.ALPHA.New(2, 2).Matches("AA")
			Expect(actual).To(Equal(expected))
		})

		It("should match AA1", func() {
			expected := abnfer.MatchResult{2, true}
			actual := abnfer.ALPHA.New(2, 2).Matches("AA1")
			Expect(actual).To(Equal(expected))
		})

		It("should not match AAA1", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.ALPHA.New(2, 2).Matches("AAA1")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("BIT", func() {
		It("should match 0", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.BIT.Matches("0")
			Expect(actual).To(Equal(expected))
		})

		It("should not match !", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.BIT.Matches("!")
			Expect(actual).To(Equal(expected))
		})

		It("should match 01", func() {
			expected := abnfer.MatchResult{2, true}
			actual := abnfer.BIT.New(1, 2).Matches("01")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("CHAR", func() {
		It("should match 0", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.CHAR.Matches("0")
			Expect(actual).To(Equal(expected))
		})

		It("should match 'BELL'", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.CHAR.Matches("\b")
			Expect(actual).To(Equal(expected))
		})

		It("should not match 'NULL'", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.CHAR.Matches(string(byte(0x00)))
			Expect(actual).To(Equal(expected))
		})
	})

	Context("CR", func() {
		It("should match \\r", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.CR.Matches("\r")
			Expect(actual).To(Equal(expected))
		})

		It("should not match 'BELL'", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.CR.Matches("\b")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("LF", func() {
		It("should match \\n", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.LF.Matches("\n")
			Expect(actual).To(Equal(expected))
		})

		It("should not match 'BELL'", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.LF.Matches("\b")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("CTL", func() {
		It("should match 'BELL'", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.CTL.Matches("\b")
			Expect(actual).To(Equal(expected))
		})

		It("should match 'DELETE'", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.CTL.Matches(char.DELETE.String())
			Expect(actual).To(Equal(expected))
		})

		It("should not match ' '", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.CTL.Matches(" ")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("DIGIT", func() {
		It("should match 1", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.DIGIT.Matches("1")
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.CTL.Matches("A")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("DQUOTE", func() {
		It("should match \"", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.DQUOTE.Matches(`"`)
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.DQUOTE.Matches("A")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("HEXDIG", func() {
		It("should match 1", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.HEXDIG.Matches(`1`)
			Expect(actual).To(Equal(expected))
		})

		It("should match F", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.HEXDIG.Matches(`F`)
			Expect(actual).To(Equal(expected))
		})

		It("should not match G", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.HEXDIG.Matches("G")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("HTAB", func() {
		It("should match \\t", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.HTAB.Matches("\t")
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.HTAB.Matches("A")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("SP", func() {
		It("should match 'SPACE'", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.SP.Matches(" ")
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.SP.Matches("A")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("WSP", func() {
		It("should match 'SPACE'", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.WSP.Matches(" ")
			Expect(actual).To(Equal(expected))
		})

		It("should match \\t", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.HTAB.Matches("\t")
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.WSP.Matches("A")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("LWSP", func() {
		It("should match 'SPACE'", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.LWSP.Matches(" ")
			Expect(actual).To(Equal(expected))
		})

		It("should match \\t", func() {
			expected := abnfer.MatchResult{1, true}
			actual := abnfer.LWSP.Matches("\t")
			Expect(actual).To(Equal(expected))
		})

		It("should match \\r\\n\\t", func() {
			expected := abnfer.MatchResult{3, true}
			actual := abnfer.LWSP.Matches("\r\n\t")
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := abnfer.MatchResult{}
			actual := abnfer.LWSP.Matches("A")
			Expect(actual).To(Equal(expected))
		})
	})

	// Context("OCTET", func() {
	// 	It("should match 'SPACE'", func() {
	// 		expected := abnfer.MatchResult{1, true}
	// 		actual := abnfer.OCTET.Matches(" ")
	// 		Expect(actual).To(Equal(expected))
	// 	})

	// 	It("should match \\t", func() {
	// 		expected := abnfer.MatchResult{1, true}
	// 		actual := abnfer.OCTET.Matches("\t")
	// 		Expect(actual).To(Equal(expected))
	// 	})
	// })
})
