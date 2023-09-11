package gola

type PsCaseEntry struct { // this needs a better name, not Ps
	AssertFn string
}

type TestCaseEntry struct {
	DoesContain    []string
	DoesNotContain []string
	Below          []string
	EqualLo        []string
	Inside         []string
	EqualHi        []string
	Above          []string
}

type BhTest struct { // binder helper
	First  string
	Second string
	Assign string
	Entry  TestCaseEntry
}

type BhTestCollection map[string]*BhTest

type TypeNameID = string

type TypeSpec struct {
	TypeName           string
	GoType             string
	DisplayType        string
	UnderlyingTypeName string
	FlagName           string
	Short              string
	Def                any
	Assign             string
	Setup              string
	BindTo             string
	Assert             string
	QuoteExpect        bool
	Equate             string
	Validatable        bool
	ForeignValidatorFn bool
	GenerateSlice      bool
	SliceFlagName      string
	SliceShort         string
	DefSliceVal        string
	ExpectSlice        string
	SliceValue         string
	IsOptionLess       bool
	OptionValue        string
	QuoteOptionValue   bool
	CommandLineValue   string
	TcEntry            *PsCaseEntry
	Comparable         bool
	BindDoc            string
	BindValidatedDoc   string
	Containable        bool
	BhParent           string
	CastLiteralsAs     string
	BhTests            BhTestCollection
}

func buildTypes() *typeCollection {
	return &typeCollection{
		"Enum": &TypeSpec{
			TypeName:           "Enum",
			GoType:             "string",
			DisplayType:        "enum",
			UnderlyingTypeName: "String",
			FlagName:           "Format",
			Short:              "f",
			Def:                "xml",
			Assign:             "outputFormatEnum := outputFormatEnumInfo.NewValue()",
			Setup:              "paramSet.Native.Format = XMLFormatEn",
			BindTo:             "&outputFormatEnum.Source",
			Assert:             `Expect(value).To(Equal("xml"))`,
			QuoteExpect:        true,
			Equate:             "Equal",
			Validatable:        true,
			ForeignValidatorFn: true,
			GenerateSlice:      false,
			SliceFlagName:      "Formats",
			SliceShort:         "F",
			DefSliceVal:        "[]string{}",
			ExpectSlice:        `[]string{"xml", "json", "text"}`,
			SliceValue:         "xml,json,text",
			OptionValue:        "json",
			TcEntry: &PsCaseEntry{
				AssertFn: `func() { Expect(outputFormatEnum.Source).To(Equal("json")) }`,
			},
			BindDoc: `

// Note that normally the client would bind to a member of the native parameter
// set. However, since there is a discrepancy between the type of the native int
// based pseudo enum member and the equivalent acceptable string value typed by
// the user on the command line (idiomatically stored on the enum info), the
// client needs to extract the enum value from the enum info, something like this:
//
// paramSet.Native.Format = OutputFormatEnumInfo.Value()
//
// The best place to put this would be inside the PreRun/PreRunE function, assuming the
// param set and the enum info are both in scope. Actually, every int based enum
// flag, would need to have this assignment performed.
			`,
			BindValidatedDoc: `

// Custom enum types created via the generic 'EnumInfo'/'EnumValue' come with a 'IsValid' method.
// The client can utilise this method inside a custom function passed into 'BindValidatedEnum'.
// The implementation would simply call this method, either on the EnumInfo or the EnumValue.
// Please see the readme for more details.
			`,
			Containable: true,
			BhTests: BhTestCollection{
				"Contains": &BhTest{
					First:  `[]string{"json", "text", "xml"}`,
					Second: "\"null\"",
					Assign: "outputFormatEnum.Source = value",
					Entry: TestCaseEntry{
						DoesContain:    []string{`"xml"`, "true"},
						DoesNotContain: []string{`"scr"`, "false"},
					},
				},
			},
		},

		"String": &TypeSpec{
			TypeName:      "String",
			GoType:        "string",
			FlagName:      "Pattern",
			Short:         "P",
			Def:           "default-pattern",
			Setup:         `paramSet.Native.Pattern = \"{{OPTION-VALUE}}\"`,
			Assert:        `Assert        = "Expect(value).To(Equal(\"{{OPTION-VALUE}}\"))"`,
			QuoteExpect:   true,
			Equate:        "Equal",
			Validatable:   true,
			GenerateSlice: true,
			SliceFlagName: "Directories",
			SliceShort:    "C",
			DefSliceVal:   "[]string{}",
			ExpectSlice:   `[]string{"alpha", "beta", "delta"}`,
			SliceValue:    "alpha, beta, delta",
			OptionValue:   "music.infex",
			//
			TcEntry: &PsCaseEntry{
				// AssertFn function is optional, but is the first item checked for
				// Next we assume Expect(value).To(something),
				// where something can be
				// - Equal[type](value) => type is optional template variable
				// - BeTrue()
				// or any other matcher
			},
			Comparable:  true,
			Containable: true,

			BhTests: BhTestCollection{
				"Within": &BhTest{
					First:  `"c"`,
					Second: `"e"`,
					Entry: TestCaseEntry{
						Below:   []string{`"b"`, "false"},
						EqualLo: []string{`"c"`, "true"},
						Inside:  []string{`"d"`, "true"},
						EqualHi: []string{`"e"`, "true"},
						Above:   []string{`"f"`, "false"},
					},
				},

				// TODO: more operators to come ...
			},
		},

		"Int": &TypeSpec{
			TypeName:      "Int",
			GoType:        "int",
			FlagName:      "Offset",
			Short:         "o",
			Def:           -1,
			Setup:         "paramSet.Native.Offset = {{OPTION-VALUE}}",
			Assert:        "Expect(value).To(Equal({{OPTION-VALUE}}))",
			Equate:        "Equal",
			Validatable:   true,
			GenerateSlice: true,
			SliceFlagName: "Offsets",
			SliceShort:    "D",
			DefSliceVal:   "[]int{}",
			ExpectSlice:   `[]int{2, 4, 6, 8}`,
			SliceValue:    "2,4,6,8",
			OptionValue:   "-9",
			TcEntry:       &PsCaseEntry{},
			//
			Comparable:  true,
			Containable: true,
			//
		},

		"Int8": &TypeSpec{
			TypeName:    "Int8",
			GoType:      "int8",
			FlagName:    "Offset8",
			Short:       "o",
			Def:         "int8(-1)",
			Setup:       "paramSet.Native.Offset8 = {{OPTION-VALUE}}",
			Assert:      "Expect(value).To(Equal({{OPTION-VALUE}}))",
			Equate:      "Equal",
			Validatable: true,
			OptionValue: "-99",
			//
			TcEntry: &PsCaseEntry{},
			//
			Comparable:  true,
			Containable: true,
			//
			BhParent:       "Int",
			CastLiteralsAs: "int8",
		},

		"Int16": &TypeSpec{
			TypeName:    "Int16",
			GoType:      "int16",
			FlagName:    "Offset16",
			Short:       "o",
			Def:         "int16(-1)",
			Setup:       "paramSet.Native.Offset16 = {{OPTION-VALUE}}",
			Assert:      "Expect(value).To(Equal({{OPTION-VALUE}}))",
			Equate:      "Equal",
			Validatable: true,
			OptionValue: "-999",
			//
			TcEntry: &PsCaseEntry{}, // why no generate slice?
			//
			Comparable:  true,
			Containable: true,
			//
			BhParent:       "Int",
			CastLiteralsAs: "int16",
		},

		"Int32": &TypeSpec{
			TypeName:    "Int32",
			GoType:      "int32",
			FlagName:    "Offset32",
			Short:       "o",
			Def:         "int32(-1)",
			Setup:       "paramSet.Native.Offset32 = {{OPTION-VALUE}}",
			Assert:      "Expect(value).To(Equal({{OPTION-VALUE}}))",
			Equate:      "Equal",
			Validatable: true,
			OptionValue: "-9999",
			//
			TcEntry: &PsCaseEntry{},
			//
			GenerateSlice: true,
			SliceFlagName: "Offsets32",
			SliceShort:    "O",
			DefSliceVal:   "[]int32{}",
			ExpectSlice:   "[]int32{2, 4, 6, 8}",
			SliceValue:    "2, 4, 6, 8",
			//
			Comparable:  true,
			Containable: true,
			//
			BhParent:       "Int",
			CastLiteralsAs: "int32",
		},

		"Int64": &TypeSpec{
			TypeName:    "Int64",
			GoType:      "int64",
			FlagName:    "Offset64",
			Short:       "o",
			Def:         "int64(-1)",
			Setup:       "paramSet.Native.Offset64 = {{OPTION-VALUE}}",
			Assert:      "Expect(value).To(Equal({{OPTION-VALUE}}))",
			Equate:      "Equal",
			Validatable: true,
			OptionValue: "-99999",
			//
			TcEntry: &PsCaseEntry{},
			//
			GenerateSlice: true,
			SliceFlagName: "Offsets64",
			SliceShort:    "O",
			DefSliceVal:   "[]int64{}",
			ExpectSlice:   "[]int64{2, 4, 6, 8}",
			SliceValue:    "2, 4, 6, 8",
			//
			Comparable:  true,
			Containable: true,
			//
			BhParent:       "Int",
			CastLiteralsAs: "int64",
		},

		"Uint": &TypeSpec{
			// ☢️☢️☢️ In the PowerShell version, misspelt as Unit, this means that
			// whats published, is probably broken see issue:
			// - https://github.com/snivilised/cobrass/issues/192
			//
			TypeName:    "Uint",
			GoType:      "uint",
			FlagName:    "Count",
			Short:       "c",
			Def:         "uint(0)",
			Setup:       "paramSet.Native.Count = {{OPTION-VALUE}}",
			Assert:      "Expect(value).To(Equal({{OPTION-VALUE}}))",
			Equate:      "Equal",
			Validatable: true,
			OptionValue: "99999",
			//
			TcEntry: &PsCaseEntry{},
			//
			GenerateSlice: true,
			SliceFlagName: "Counts",
			SliceShort:    "P",
			DefSliceVal:   "[]uint{}",
			ExpectSlice:   "[]uint{2, 4, 6, 8}",
			SliceValue:    "2, 4, 6, 8",
			//
			Comparable:  true,
			Containable: true,
			//
			BhParent:       "Int",
			CastLiteralsAs: "uint",
		},

		"Uint8": &TypeSpec{
			TypeName:    "Uint8",
			GoType:      "uint8",
			FlagName:    "Count8",
			Short:       "c",
			Def:         "uint8(0)",
			Setup:       "paramSet.Native.Count8 = {{OPTION-VALUE}}",
			Assert:      "Expect(value).To(Equal({{OPTION-VALUE}}))",
			Equate:      "Equal",
			Validatable: true,
			OptionValue: "33",
			//
			TcEntry: &PsCaseEntry{},
			//
			Comparable:  true,
			Containable: true,
			//
			BhParent:       "Int",
			CastLiteralsAs: "uint8",
		},

		"Uint16": &TypeSpec{
			TypeName:    "Uint16",
			GoType:      "uint16",
			FlagName:    "Count16",
			Short:       "c",
			Def:         "uint16(0)",
			Setup:       "paramSet.Native.Count16 = {{OPTION-VALUE}}",
			Assert:      "Expect(value).To(Equal({{OPTION-VALUE}}))",
			Equate:      "Equal",
			Validatable: true,
			OptionValue: "333",
			//
			TcEntry: &PsCaseEntry{},
			//
			Comparable:  true,
			Containable: true,
			//
			BhParent:       "Int",
			CastLiteralsAs: "uint16",
		},

		"Uint32": &TypeSpec{
			TypeName:    "Uint32",
			GoType:      "uint32",
			FlagName:    "Count32",
			Short:       "c",
			Def:         "uint32(0)",
			Setup:       "paramSet.Native.Count32 = {{OPTION-VALUE}}",
			Assert:      "Expect(value).To(Equal({{OPTION-VALUE}}))",
			Equate:      "Equal",
			Validatable: true,
			OptionValue: "3333",
			//
			TcEntry: &PsCaseEntry{},
			//
			Comparable:  true,
			Containable: true,
			//
			BhParent:       "Int",
			CastLiteralsAs: "uint32",
		},

		"Uint64": &TypeSpec{
			TypeName:    "Uint64",
			GoType:      "uint64",
			FlagName:    "Count64",
			Short:       "c",
			Def:         "uint64(0)",
			Setup:       "paramSet.Native.Count64 = {{OPTION-VALUE}}",
			Assert:      "Expect(value).To(Equal({{OPTION-VALUE}}))",
			Equate:      "Equal",
			Validatable: true,
			OptionValue: "33333",
			//
			TcEntry: &PsCaseEntry{},
			//
			Comparable:  true,
			Containable: true,
			//
			BhParent:       "Int",
			CastLiteralsAs: "uint32",
		},

		"Float32": &TypeSpec{
			TypeName:    "Float32",
			GoType:      "float32",
			FlagName:    "Gradientf32",
			Short:       "t",
			Def:         "float32(0)",
			Setup:       "paramSet.Native.Gradientf32 = {{OPTION-VALUE}}",
			Assert:      "Expect(value).To(Equal({{OPTION-VALUE}}))",
			Equate:      "Equal",
			Validatable: true,
			OptionValue: "32.0",
			//
			TcEntry: &PsCaseEntry{},
			//
			GenerateSlice: true,
			SliceFlagName: "Gradientsf32",
			SliceShort:    "G",
			DefSliceVal:   "[]float32{}",
			ExpectSlice:   "[]float32{3.0, 5.0, 7.0, 9.0}",
			SliceValue:    "3.0, 5.0, 7.0, 9.0",
			//
			Comparable:     true,
			Containable:    true,
			BhParent:       "Int",
			CastLiteralsAs: "float32",
		},

		"Float64": &TypeSpec{
			TypeName:    "Float64",
			GoType:      "float64",
			FlagName:    "Gradientf64",
			Short:       "t",
			Def:         "float64(0)",
			Setup:       "paramSet.Native.Gradientf64 = {{OPTION-VALUE}}",
			Assert:      "Expect(value).To(Equal({{OPTION-VALUE}}))",
			Equate:      "Equal",
			Validatable: true,
			OptionValue: "64.1234",
			//
			TcEntry: &PsCaseEntry{},
			//
			GenerateSlice: true,
			SliceFlagName: "Gradientsf64",
			SliceShort:    "G",
			DefSliceVal:   "[]float64{}",
			ExpectSlice:   "[]float64{4.0, 6.0, 8.0, 10.0}",
			SliceValue:    "4.0, 6.0, 8.0, 10.0",
			//
			Comparable:     true,
			Containable:    true,
			BhParent:       "Int",
			CastLiteralsAs: "float64",
		},

		"Bool": &TypeSpec{
			TypeName:    "Bool",
			GoType:      "bool",
			FlagName:    "Concise",
			Short:       "c",
			Def:         "false",
			Setup:       "paramSet.Native.Concise = {{OPTION-VALUE}}",
			Assert:      "Expect(value).To(BeTrue())",
			QuoteExpect: true,
			Equate:      "Equal",
			// bool is not Validatable, because there's not much to validate,
			// can only be true or false
			Validatable:   true,
			GenerateSlice: true,
			SliceFlagName: "Switches",
			SliceShort:    "S",
			DefSliceVal:   "[]bool{}",
			ExpectSlice:   "[]bool{true, false, true, false}",
			SliceValue:    "true, false, true, false",
			IsOptionLess:  true,
			OptionValue:   "true",
			TcEntry:       &PsCaseEntry{},
		},

		"Duration": &TypeSpec{
			TypeName: "Duration",
			GoType:   "time.Duration",
			FlagName: "Latency",
			Short:    "l",
			Def:      `duration("0ms")`,
			Setup:    "paramSet.Native.Latency = {{OPTION-VALUE}}",
			Assert: `
			expect := {{OPTION-VALUE}}
			Expect(value).To(BeEquivalentTo(expect))
			`,
			Equate:           "BeEquivalentTo",
			Validatable:      true,
			OptionValue:      "300ms",
			QuoteOptionValue: true,
			TcEntry:          &PsCaseEntry{},
			//
			GenerateSlice: true,
			SliceFlagName: "Latencies",
			SliceShort:    "L",
			DefSliceVal:   "[]time.Duration{}",
			ExpectSlice:   `[]time.Duration{duration("1s"), duration("2s"), duration("3s")}`,
			SliceValue:    "1s, 2s, 3s",
			Comparable:    true,
			//
			// 'duration' is a function defined in the test suite, that is syntactically the
			// same as a type cast.
			//
			CastLiteralsAs: "duration",
			BhTests:        BhTestCollection{
				// tbd ...
			},
		},

		"IPNet": &TypeSpec{
			TypeName:         "IPNet",
			GoType:           "net.IPNet",
			FlagName:         "IPAddress",
			Short:            "i",
			Def:              `ipnet("default")`,
			Setup:            "paramSet.Native.IPAddress = {{OPTION-VALUE}}",
			Assert:           "Expect(value).To(BeEquivalentTo({{OPTION-VALUE}}))",
			Equate:           "BeEquivalentTo",
			Validatable:      true,
			OptionValue:      `ipnet("orion.net")`,
			CommandLineValue: "172.16.0.0",
			TcEntry:          &PsCaseEntry{},
		},

		"IPMask": &TypeSpec{
			TypeName:         "IPMask",
			GoType:           "net.IPMask",
			FlagName:         "IPMask",
			Short:            "m",
			Def:              `ipmask("default")`,
			Setup:            "paramSet.Native.IPMask = {{OPTION-VALUE}}",
			Assert:           "Expect(value).To(BeEquivalentTo({{OPTION-VALUE}}))",
			Equate:           "BeEquivalentTo",
			Validatable:      true,
			OptionValue:      `ipmask("orion.net")`,
			CommandLineValue: "255.255.255.0",
			TcEntry:          &PsCaseEntry{},
		},
	}
}
