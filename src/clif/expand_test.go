package clif_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/cobrass/src/clif"
)

type baseTE struct {
	given        string
	shouldReturn string
	expected     []string
}

type expandTE struct {
	baseTE
	before clif.ThirdPartyPositionalArgs
	flags  clif.ThirdPartyCommandLine
	after  []clif.ThirdPartyFlagName
}

var _ = Describe("Expand", func() {
	DescribeTable("ThirdPartyFlags",
		func(entry *expandTE) {
			actual := clif.Expand(entry.before, entry.flags, entry.after...)
			Expect(actual).To(HaveExactElements(entry.expected))
		},
		func(entry *expandTE) string {
			return fmt.Sprintf("🧪 ===> given: '%v', should: return '%v'",
				entry.given, entry.shouldReturn,
			)
		},

		Entry(nil, &expandTE{
			baseTE: baseTE{
				given:        "none before/none after",
				shouldReturn: "the single before",
				expected:     []string{"--dry-run"},
			},
			before: clif.ThirdPartyPositionalArgs{},
			flags:  clif.ThirdPartyCommandLine{"--dry-run"},
		}),

		Entry(nil, &expandTE{
			baseTE: baseTE{
				given:        "1 before, none after",
				shouldReturn: "the single before with flags",
				expected:     []string{"file.jpg", "--dry-run"},
			},
			before: clif.ThirdPartyPositionalArgs{"file.jpg"},
			flags:  clif.ThirdPartyCommandLine{"--dry-run"},
		}),

		Entry(nil, &expandTE{
			baseTE: baseTE{
				given:        "none before/none after",
				shouldReturn: "after and single after",
				expected:     []string{"--dry-run", "result.jpg"},
			},
			before: clif.ThirdPartyPositionalArgs{},
			flags:  clif.ThirdPartyCommandLine{"--dry-run"},
			after:  []clif.ThirdPartyFlagName{"result.jpg"},
		}),

		Entry(nil, &expandTE{
			baseTE: baseTE{
				given:        "multi before, some flags and single after",
				shouldReturn: "before, flags and after",
				expected: []string{"first.jpg", "second.jpg",
					"--dry-run", "--interlace", "plane",
					"result-1.jpg", "result-2.jpg",
				},
			},
			before: clif.ThirdPartyPositionalArgs{"first.jpg", "second.jpg"},
			flags:  clif.ThirdPartyCommandLine{"--dry-run", "--interlace", "plane"},
			after:  []clif.ThirdPartyFlagName{"result-1.jpg", "result-2.jpg"},
		}),

		Entry(nil, &expandTE{
			baseTE: baseTE{
				given:        "multi before, no flags and single after",
				shouldReturn: "before and after",
				expected: []string{"first.jpg", "second.jpg",
					"result-1.jpg", "result-2.jpg",
				},
			},
			before: clif.ThirdPartyPositionalArgs{"first.jpg", "second.jpg"},
			flags:  clif.ThirdPartyCommandLine{},
			after:  []clif.ThirdPartyFlagName{"result-1.jpg", "result-2.jpg"},
		}),
	)
})
