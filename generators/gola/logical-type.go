package gola

type PsCaseEntry struct { // this needs a better name, not Ps
	AssertFn string
}

type TestCaseEntry struct {
}

type BhTest struct { // binder helper
}

type BhTestCollection map[string]*BhTest

type LogicalType struct {
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
	QuoteExpect        string
	Equate             string
	Validatable        bool
	ForeignValidatorFn bool
	GenerateSlice      bool
	SliceFlagName      string
	SliceShort         string
	DefSliceVal        string
	ExpectSlice        string
	SliceValue         string
	OptionValue        string
	TcEntry            PsCaseEntry
	BindDoc            string
	BindValidatedDoc   string
	Containable        bool
	BhTests            BhTestCollection
}
