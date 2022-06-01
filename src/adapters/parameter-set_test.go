package adapters

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OutputFormatEnum int

const (
	XmlFormatEn OutputFormatEnum = iota + 1
	JsonFormatEn
	TextFormatEn
	ScribbleFormatEn
)

type InfexionShapeEnum int

const (
	FullNameShapeEn InfexionShapeEnum = iota + 1
	NameShapeEn
	SubPathShapeEn
)

type TravseralStratgeyEnum int

const (
	TraverseAllEn TravseralStratgeyEnum = iota + 1
	TraverseLeafEn
)

type ParameterSetSuite struct {
	suite.Suite

	Native           FooParameterSet
	Expected         GenericParameterSet
	MissingDirectory GenericParameterSet
}

func (suite *ParameterSetSuite) SetupTest() {
	suite.Native = FooParameterSet{
		Directory:           "/once/upon/a/time",
		Output:              "foo-bar",
		Format:              XmlFormatEn,
		Shape:               SubPathShapeEn,
		IsConcise:           true,
		Strategy:            TraverseLeafEn,
		IsOverwrite:         false,
		SegmentsFilePattern: "*infex*",
	}

	// This is the item that is generated via visiting the Cobra command's flags set
	// We don't want to force the client to use a map (generic parameter set),
	// instead we want to use a native abstraction, so turn it into a parameter set, ie:
	// --> FooParameterSet
	//
	suite.Expected = GenericParameterSet{
		"Directory":           "/once/upon/a/time",
		"Output":              "foo-bar",
		"Format":              XmlFormatEn,
		"Shape":               SubPathShapeEn,
		"IsConcise":           true,
		"Strategy":            TraverseLeafEn,
		"IsOverwrite":         false,
		"SegmentsFilePattern": "*infex*",
	}

	suite.MissingDirectory = GenericParameterSet{
		"Output":              "foo-bar",
		"Format":              XmlFormatEn,
		"Shape":               SubPathShapeEn,
		"IsConcise":           true,
		"Strategy":            TraverseLeafEn,
		"IsOverwrite":         false,
		"SegmentsFilePattern": "*infex*",
	}
}

func TestParameterSetSuite(t *testing.T) {
	suite.Run(t, new(ParameterSetSuite))
}

type FooParameterSet struct {
	Directory           string
	Output              string
	Format              OutputFormatEnum
	Shape               InfexionShapeEnum
	IsConcise           bool
	Strategy            TravseralStratgeyEnum
	IsOverwrite         bool
	SegmentsFilePattern string
}

// func (suite *ParameterSetSuite) TestNativeParameterCreationFromGenericSetWithReflection() {
// 	//
// 	// This test case is an excercise in using reflection and thus would not
// 	// normally be tested this way. See the other test TestNativeParameterCreationFromGenericSet.
// 	//
// 	// based on the laws of reflection (https://go.dev/blog/laws-of-reflection)

// 	target := NewFooParameterSet(suite.Expected)

// 	refElemStruct := reflect.ValueOf(target).Elem()
// 	refTypeOfStruct := refElemStruct.Type()

// 	if reflect.TypeOf(*target).Kind() == reflect.Struct {
// 		for i, n := 0, refElemStruct.NumField(); i < n; i++ {
// 			name := refTypeOfStruct.Field(i).Name
// 			value := suite.Expected[name]
// 			assert.Equal(suite.T(), suite.Expected[name], value)
// 		}
// 	}
// }

// func (suite *ParameterSetSuite) TestNativeParameterCreationFromGenericSet() {
// 	actual := NewFooParameterSet(suite.Expected)

// 	assert.Equal(suite.T(), suite.Expected["Directory"], actual.Directory)
// 	assert.Equal(suite.T(), suite.Expected["Output"], actual.Output)
// 	assert.Equal(suite.T(), suite.Expected["Format"], actual.Format)
// 	assert.Equal(suite.T(), suite.Expected["Shape"], actual.Shape)
// 	assert.Equal(suite.T(), suite.Expected["IsConcise"], actual.IsConcise)
// 	assert.Equal(suite.T(), suite.Expected["Strategy"], actual.Strategy)
// 	assert.Equal(suite.T(), suite.Expected["IsOverwrite"], actual.IsOverwrite)
// 	assert.Equal(suite.T(), suite.Expected["SegmentsFilePattern"], actual.SegmentsFilePattern)
// }

// CreateParameterSet
func CreateFooParameterSet(params GenericParameterSet) *FooParameterSet {
	return CreateParameterSet[FooParameterSet](params)
}
func (suite *ParameterSetSuite) TestCreateParameterSetWithDefaultOptions() {
	actual := CreateFooParameterSet(suite.Expected)

	assert.Equal(suite.T(), suite.Expected["Directory"], actual.Directory)
	assert.Equal(suite.T(), suite.Expected["Output"], actual.Output)
	assert.Equal(suite.T(), suite.Expected["Format"], actual.Format)
	assert.Equal(suite.T(), suite.Expected["Shape"], actual.Shape)
	assert.Equal(suite.T(), suite.Expected["IsConcise"], actual.IsConcise)
	assert.Equal(suite.T(), suite.Expected["Strategy"], actual.Strategy)
	assert.Equal(suite.T(), suite.Expected["IsOverwrite"], actual.IsOverwrite)
	assert.Equal(suite.T(), suite.Expected["SegmentsFilePattern"], actual.SegmentsFilePattern)
}

func deferMissingGenericParam(t *testing.T) {
	if r := recover(); r != nil {
		recovery := fmt.Sprintf("%v", r)

		pattern := errorTypes[missingNativeMemberValueEn].Info[PatternEn]
		if matched, err := regexp.MatchString(pattern, recovery); err == nil {
			assert.True(t, matched, fmt.Sprintf("panic message: '%v' did not conform to format",
				recovery))
		} else {
			assert.Fail(t, "panic message regexp error")
		}
	}
}

func (suite *ParameterSetSuite) TestCreateParameterSetStrictMissingParam() {
	func() {
		defer deferMissingGenericParam(suite.T())

		CreateParameterSet[FooParameterSet](suite.MissingDirectory)
		assert.Fail(suite.T(), "should panic when generic param is missing")
	}()
}

func (suite *ParameterSetSuite) TestCreateParameterSetNotStrictMissingParam() {

	CreateParameterSetWith[FooParameterSet](suite.MissingDirectory, ParameterSetCreateOptions{
		Strict: false,
	})
}

func (suite *ParameterSetSuite) TestCreateParameterSetNotStrict() {
	CreateParameterSetWith[FooParameterSet](suite.Expected, ParameterSetCreateOptions{
		Strict: false,
	})
}

// func (suite *ParameterSetSuite) TestNativeObjectParamMismatchWithGenericEntry() {

// 	expected := GenericParameterSet{
// 		"directory":           "/once/upon/a/time",
// 		"Output":              "foo-bar",
// 		"Format":              XmlFormatEn,
// 		"Shape":               SubPathShapeEn,
// 		"IsConcise":           true,
// 		"Strategy":            TraverseLeafEn,
// 		"IsOverwrite":         false,
// 		"SegmentsFilePattern": "*infex*",
// 	}

// 	func() {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				recovery := fmt.Sprintf("%v", r)

// 				pattern := errorTypes[missingNativeMemberValueEn].Info[PatternEn]
// 				if matched, err := regexp.MatchString(pattern, recovery); err == nil {
// 					assert.True(suite.T(), matched, fmt.Sprintf("panic message: '%v' did not conform to format",
// 						recovery))
// 				} else {
// 					assert.Fail(suite.T(), "panic message regexp error")
// 				}
// 			}
// 		}()
// 		NewFooParameterSet(expected)
// 		assert.Fail(suite.T(), "should panic when no generic parameter for native member")
// 	}()
// }

// func (suite *ParameterSetSuite) TestNativeObjectMemberTypeMismatch() {
// 	expected := GenericParameterSet{
// 		"Directory":           42.42,
// 		"Output":              "foo-bar",
// 		"Format":              XmlFormatEn,
// 		"Shape":               SubPathShapeEn,
// 		"IsConcise":           true,
// 		"Strategy":            TraverseLeafEn,
// 		"IsOverwrite":         false,
// 		"SegmentsFilePattern": "*infex*",
// 	}

// 	func() {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				recovery := fmt.Sprintf("%v", r)

// 				pattern := errorTypes[mismatchNativeMemberValueTypeEn].Info[PatternEn]
// 				if matched, err := regexp.MatchString(pattern, recovery); err == nil {
// 					assert.True(suite.T(), matched, fmt.Sprintf("panic message: '%v' did not conform to format",
// 						recovery))
// 				} else {
// 					assert.Fail(suite.T(), "panic message regexp error")
// 				}
// 			}
// 		}()
// 		NewFooParameterSet(expected)
// 		assert.Fail(suite.T(), "should panic when member type mismatches generic param type")
// 	}()
// }
