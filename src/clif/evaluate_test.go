package clif_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/cobrass/src/clif"
)

type evaluateTE struct {
	baseTE
	present   clif.PresentFlagsCollection
	secondary clif.ThirdPartyCommandLine
}

var _ = Describe("Evaluate", Ordered, func() {
	var knownBy clif.KnownByCollection

	BeforeAll(func() {
		knownBy = clif.KnownByCollection{
			"dry-run":         "D",
			"gaussian-blur":   "b",
			"sampling-factor": "f",
			"interlace":       "i",
			"strip":           "s",
			"quality":         "q",
		}
	})

	DescribeTable("ThirdPartyCommandLine",
		func(entry *evaluateTE) {
			actual := clif.Evaluate(entry.present, knownBy, entry.secondary)
			Expect(actual).To(HaveExactElements(entry.expected))
		},
		func(entry *evaluateTE) string {
			return fmt.Sprintf("🧪 ===> given: '%v', should: return '%v'",
				entry.given, entry.shouldReturn,
			)
		},

		// secondary empty
		//
		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "present contains single switch; secondary is empty",
				shouldReturn: "present",
				expected:     []string{"--dry-run"},
			},
			present: clif.PresentFlagsCollection{
				"dry-run": "true",
			},
			secondary: clif.ThirdPartyCommandLine{},
		}),

		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "present contains single flag; secondary is empty",
				shouldReturn: "present",
				expected:     []string{"--sampling-factor", "4:2:0"},
			},
			present: clif.PresentFlagsCollection{
				"sampling-factor": "4:2:0",
			},
			secondary: clif.ThirdPartyCommandLine{},
		}),

		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "present contains single flag; secondary is empty",
				shouldReturn: "present",
				expected:     []string{"--sampling-factor", "4:2:0"},
			},
			present: clif.PresentFlagsCollection{
				"sampling-factor": "4:2:0",
			},
			secondary: clif.ThirdPartyCommandLine{},
		}),

		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "present contains flag and a switch; secondary is empty",
				shouldReturn: "all present",
				expected:     []string{"--dry-run", "--sampling-factor", "4:2:0"},
			},
			present: clif.PresentFlagsCollection{
				"dry-run":         "true",
				"sampling-factor": "4:2:0",
			},
			secondary: clif.ThirdPartyCommandLine{},
		}),
		//
		// end: secondary empty

		// single secondary token
		//
		// ---> secondary switch in present
		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "present contains single switch; single long secondary switch in present",
				shouldReturn: "present, ignore secondary",
				expected:     []string{"--dry-run"},
			},
			present: clif.PresentFlagsCollection{
				"dry-run": "true",
			},
			secondary: clif.ThirdPartyCommandLine{"--dry-run"},
		}),

		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "present contains single switch; single short secondary switch in present",
				shouldReturn: "present, ignore secondary",
				expected:     []string{"--dry-run"},
			},
			present: clif.PresentFlagsCollection{
				"dry-run": "true",
			},
			secondary: clif.ThirdPartyCommandLine{"-D"},
		}),
		// ---> secondary switch NOT in present
		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "present contains single switch; single long secondary switch NOT in present",
				shouldReturn: "present with secondary",
				expected:     []string{"--dry-run", "--strip"},
			},
			present: clif.PresentFlagsCollection{
				"dry-run": "true",
			},
			secondary: clif.ThirdPartyCommandLine{"--strip"},
		}),

		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "present contains single switch; single short secondary switch NOT in present",
				shouldReturn: "present with secondary",
				expected:     []string{"--dry-run", "-s"},
			},
			present: clif.PresentFlagsCollection{
				"dry-run": "true",
			},
			secondary: clif.ThirdPartyCommandLine{"-s"},
		}),
		//
		// end: single secondary token

		// single flag/option secondary tokens
		//
		// ---> secondary flag in present

		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "present contains single flag; long secondary flag/option in present",
				shouldReturn: "present, ignore secondary",
				expected:     []string{"--sampling-factor", "4:2:0"},
			},
			present: clif.PresentFlagsCollection{
				"sampling-factor": "4:2:0",
			},
			secondary: clif.ThirdPartyCommandLine{"--sampling-factor", "2x1"},
		}),

		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "present contains single flag; short secondary flag/option in present",
				shouldReturn: "present, ignore secondary",
				expected:     []string{"--sampling-factor", "4:2:0"},
			},
			present: clif.PresentFlagsCollection{
				"sampling-factor": "4:2:0",
			},
			secondary: clif.ThirdPartyCommandLine{"-f", "2x1"},
		}),
		// ---> secondary flag NOT in present
		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "present contains single flag; long secondary flag/option NOT in present",
				shouldReturn: "present with secondary",
				expected:     []string{"--sampling-factor", "4:2:0", "--gaussian-blur", "0.05"},
			},
			present: clif.PresentFlagsCollection{
				"sampling-factor": "4:2:0",
			},
			secondary: clif.ThirdPartyCommandLine{"--gaussian-blur", "0.05"},
		}),

		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "present contains single flag; short secondary flag/option NOT in present",
				shouldReturn: "present with secondary",
				expected:     []string{"--sampling-factor", "4:2:0", "-b", "0.05"},
			},
			present: clif.PresentFlagsCollection{
				"sampling-factor": "4:2:0",
			},
			secondary: clif.ThirdPartyCommandLine{"-b", "0.05"},
		}),
		//
		// end: single flag/option secondary tokens

		// secondary switch followed by a flag
		//
		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "secondary switch followed by a flag; both in present",
				shouldReturn: "present, ignore secondary",
				expected:     []string{"--dry-run", "--sampling-factor", "4:2:0"},
			},
			present: clif.PresentFlagsCollection{
				"dry-run":         "true",
				"sampling-factor": "4:2:0",
			},
			secondary: clif.ThirdPartyCommandLine{"--dry-run", "--sampling-factor", "2x1"},
		}),

		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "secondary switch followed by a flag; switch in present",
				shouldReturn: "present, with secondary flag",
				expected:     []string{"--dry-run", "--sampling-factor", "2x1"},
			},
			present: clif.PresentFlagsCollection{
				"dry-run": "true",
			},
			secondary: clif.ThirdPartyCommandLine{"--dry-run", "--sampling-factor", "2x1"},
		}),

		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "secondary switch followed by a flag; flag in present",
				shouldReturn: "present, with secondary switch",
				expected:     []string{"--sampling-factor", "4:2:0", "--dry-run"},
			},
			present: clif.PresentFlagsCollection{
				"sampling-factor": "4:2:0",
			},
			secondary: clif.ThirdPartyCommandLine{"--dry-run", "--sampling-factor", "2x1"},
		}),

		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "secondary switch followed by a flag; neither in present",
				shouldReturn: "present, secondary switch and flag",
				expected:     []string{"--gaussian-blur", "0.05", "--dry-run", "--sampling-factor", "2x1"},
			},
			present: clif.PresentFlagsCollection{
				"gaussian-blur": "0.05",
			},
			secondary: clif.ThirdPartyCommandLine{"--dry-run", "--sampling-factor", "2x1"},
		}),
		//
		// end: secondary switch followed by a flag

		Entry(nil, &evaluateTE{
			baseTE: baseTE{
				given:        "many in present; many in secondary",
				shouldReturn: "present flags/options overriding secondary flags/options",
				expected: []string{
					"--gaussian-blur", "0.05",
					"-i", "plane",
					"-D",
					"-f", "2x1",
					"--strip",
				},
			},
			present: clif.PresentFlagsCollection{
				"gaussian-blur": "0.05",
				"i":             "plane",
			},
			secondary: clif.ThirdPartyCommandLine{
				"-D", "-f", "2x1", "--strip", "--gaussian-blur", "0.15", "--interlace", "line",
			},
		}),
	)
})
