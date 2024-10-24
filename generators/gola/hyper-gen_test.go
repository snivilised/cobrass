package gola_test

import (
	"fmt"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
)

type typeSpec struct {
	SubTypes []string
	//
	TypeName           string
	GoType             string
	DisplayType        string
	UnderlyingTypeName string
	FlagName           string
	Short              string
	Def                string
	// Assign             string
	// Setup              string
	// BindTo             string
	// Assert             string
	// QuoteExpect        string // bool
	Equate string
	// Validatable        string // bool
	// ForeignValidatorFn string // bool
	// GenerateSlice      string // bool
	SliceFlagName string
	SliceShort    string
	DefSliceVal   string
	ExpectSlice   string
	SliceValue    string
	OptionValue   string
	TcEntry       string // object
}

func generate(spec *typeSpec, subType string) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf(`    "%v%v": &TypeSpec{`+"\n", spec.TypeName, subType))
	builder.WriteString(fmt.Sprintf(`      TypeName:           "%v%v",`+"\n", spec.TypeName, subType))
	builder.WriteString(fmt.Sprintf(`      GoType:           "%v%v",`+"\n", spec.GoType, subType))

	if spec.DisplayType != "" {
		builder.WriteString(fmt.Sprintf(`      DisplayType:      "%v",`+"\n", spec.DisplayType))
	}

	builder.WriteString(fmt.Sprintf(`      UnderlyingTypeName: "%v",`+"\n", spec.UnderlyingTypeName))
	builder.WriteString(fmt.Sprintf(`      FlagName: "%v%v",`+"\n", spec.FlagName, subType))
	builder.WriteString(fmt.Sprintf(`      Short: "%v",`+"\n", spec.Short))

	if subType == "" {
		builder.WriteString(fmt.Sprintf(`      Def: "%v",`+"\n", spec.Def))
	} else {
		builder.WriteString(fmt.Sprintf(`      Def: "%v%v(-1)",`+"\n", spec.GoType, subType))
	}

	builder.WriteString(`      Assign: "...",` + "\n")
	builder.WriteString(`      Setup: "...",` + "\n")
	builder.WriteString(`      BindTo: "...",` + "\n")
	builder.WriteString(`      Assert: "...",` + "\n")
	builder.WriteString(`      QuoteExpect: true,` + "\n")
	builder.WriteString(`      Equate: "Equal",` + "\n")
	builder.WriteString(`      Validatable: true,` + "\n")
	builder.WriteString(`      ForeignValidatorFn: true,` + "\n")
	builder.WriteString(`      GenerateSlice: false,` + "\n")
	//----

	builder.WriteString(fmt.Sprintf(`      SliceFlagName: "%v",`+"\n", spec.SliceFlagName))
	builder.WriteString(fmt.Sprintf(`      SliceShort: "%v",`+"\n", spec.SliceShort))
	builder.WriteString(fmt.Sprintf(`      DefSliceVal: "%v",`+"\n", spec.DefSliceVal))
	builder.WriteString(fmt.Sprintf(`      ExpectSlice: "%v",`+"\n", spec.ExpectSlice))
	builder.WriteString(`      SliceValue: "...",` + "\n")
	builder.WriteString(`      OptionValue: "...",`)
	builder.WriteString(`
	TcEntry: &PsCaseEntry{
	  AssertFn: "func() { Expect(outputFormatEnum.Source).To(Equal(\"json\")) }",
	},`)
	builder.WriteString("\n      Containable: true," + "\n")
	builder.WriteString("    }," + "\n")

	GinkgoWriter.Println(builder.String())
}

var _ = Describe("HyperGen", Ordered, func() {
	var (
		repo, testPath, sourcePath string
		orderedTypes               []string
		types                      map[string]*typeSpec
	)

	BeforeAll(func() {
		repo = Repo("../..")
		testPath = filepath.Join("generators", "gola", "out", "assistant")
		sourcePath = filepath.Join("src", "assistant")
		_ = testPath
		_ = sourcePath
		_ = repo

		orderedTypes = []string{
			"Enum",
			"String",
			"Int",
			// "Int8",
			// "Int16",
			// "Int32",
			// "Int64",
			"UInt",
			// "UInt8",
			// "UInt16",
			// "UInt32",
			// "UInt64",
			"Float32",
			"Float64",
			"Bool",
			"Duration",
			"IPNet",
			"IPMask",
		}

		types = map[string]*typeSpec{
			"Enum": {
				TypeName:           "Enum",
				GoType:             "string",
				DisplayType:        "enum",
				UnderlyingTypeName: "String",
				FlagName:           "Format",
				Short:              "f",
				Def:                "xml",
				Equate:             "Equal",
				SliceFlagName:      "Formats",
				SliceShort:         "F",
				DefSliceVal:        "[]string{}",
				ExpectSlice:        `[]string{\"xml\", \"json\", \"text\"}`,
			},

			"String": {
				TypeName:           "String",
				GoType:             "string",
				UnderlyingTypeName: "String",
				FlagName:           "Format",
				Short:              "f",
				Def:                "xml",
				Equate:             "Equal",
				SliceFlagName:      "Formats",
				SliceShort:         "F",
				DefSliceVal:        "[]string{}",
				ExpectSlice:        `[]string{\"xml\", \"json\", \"text\"}`,
			},

			"Int": {
				SubTypes:      []string{"8", "16", "32", "64"},
				TypeName:      "Int",
				GoType:        "int",
				FlagName:      "Format",
				Short:         "f",
				Def:           "xml",
				Equate:        "Equal",
				SliceFlagName: "Formats",
				SliceShort:    "F",
				DefSliceVal:   "[]string{}",
				ExpectSlice:   `[]string{\"xml\", \"json\", \"text\"}`,
			},

			"Uint": {
				SubTypes:      []string{"8", "16", "32", "64"},
				TypeName:      "Uint",
				GoType:        "uint",
				FlagName:      "Count",
				Short:         "f",
				Def:           "xml",
				Equate:        "Equal",
				SliceFlagName: "Formats",
				SliceShort:    "F",
				DefSliceVal:   "[]string{}",
				ExpectSlice:   `[]string{\"xml\", \"json\", \"text\"}`,
			},

			"Float32": {
				TypeName:      "Float32",
				GoType:        "float32",
				FlagName:      "Gradientf32",
				Short:         "t",
				Def:           "float32(0)",
				Equate:        "Equal",
				SliceFlagName: "Gradientsf32",
				SliceShort:    "G",
				DefSliceVal:   "[]float32{}",
				ExpectSlice:   "[]float32{3.0, 5.0, 7.0, 9.0}",
				SliceValue:    "3.0, 5.0, 7.0, 9.0",
			},

			"Float64": {
				TypeName:      "Float64",
				GoType:        "float64",
				FlagName:      "Gradientf64",
				Short:         "t",
				Def:           "float64(0)",
				Equate:        "Equal",
				SliceFlagName: "Gradientsf64",
				SliceShort:    "G",
				DefSliceVal:   "[]float64{}",
				ExpectSlice:   "[]float64{4.0, 6.0, 8.0, 10.0}",
				SliceValue:    "4.0, 6.0, 8.0, 10.0",
			},

			"Bool": {
				TypeName:      "Bool",
				GoType:        "bool",
				FlagName:      "Concise",
				Short:         "c",
				Def:           "false",
				Equate:        "Equal",
				SliceFlagName: "Switches",
				SliceShort:    "S",
				DefSliceVal:   "[]bool{}",
				ExpectSlice:   "[]bool{true, false, true, false}",
				SliceValue:    "true, false, true, false",
			},

			"Duration": {
				TypeName:      "Duration",
				GoType:        "time.Duration",
				FlagName:      "Latency",
				Short:         "l",
				Def:           `duration(\"0ms\")`,
				Equate:        "BeEquivalentTo",
				SliceFlagName: "Latencies",
				SliceShort:    "L",
				DefSliceVal:   "[]time.Duration{}",
				ExpectSlice:   `[]time.Duration{duration(\"1s\"), duration(\"2s\"), duration(\"3s\")}`,
				SliceValue:    "1s, 2s, 3s",
			},

			"IPNet": {
				TypeName: "IPNet",
				GoType:   "net.IPNet",
				FlagName: "IPAddress",
				Short:    "i",
				Def:      `ipnet(\"default\")`,
				Equate:   "BeEquivalentTo",
			},

			"IPMask": {
				TypeName: "IPMask",
				GoType:   "net.IPMask",
				FlagName: "IPMask",
				Short:    "m",
				Def:      `ipmask(\"default\")`,
				Equate:   "BeEquivalentTo",
			},
		}
	})

	Context("generate code", func() {
		XIt("show types in console...", func() {
			for _, ot := range orderedTypes {
				if spec, ok := types[ot]; ok {
					generate(spec, "")

					if len(spec.SubTypes) > 0 {
						for _, st := range spec.SubTypes {
							generate(spec, st)
						}
					}
				}
			}
		})
	})
})
