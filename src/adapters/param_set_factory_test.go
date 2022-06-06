package adapters_test

import (
	"fmt"
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	"github.com/snivilised/cobrass/src/adapters"
	"github.com/snivilised/cobrass/src/testhelpers"
)

// We could have a CommandAssistant to handle a commands flags and to replace the
// param set factory and returns a native param-set object.
//

// - New (flagSet pflag.FlagSet, options ExtractPsOptions) - with visit strategy
// - Create()
// - CreateWith(CreatePsOptions)

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

type Cradle struct {
	Name string
	Size int
	Temp float64
}

type FooParameterSet struct {
	Directory string
	Output    string
	Format    OutputFormatEnum
	Shape     InfexionShapeEnum
	Concise   bool
	Strategy  TravseralStratgeyEnum
	Overwrite bool
	Pattern   string
}

type PsEntry struct {
	Params  adapters.GenericParamSet
	Strict  bool
	Message string
	PanicFn func()
}

type PsExtractEntry struct {
	PsEntry
	Strategy adapters.VisitStrategyEnum
}

var _ = Describe("ParamSetFactory", Label("adapter"), func() {
	var factory *adapters.ParamSetFactory[FooParameterSet]
	var Native FooParameterSet
	var Expected adapters.GenericParamSet
	var MissingDirectory adapters.GenericParamSet

	BeforeEach(func() {

		Native = FooParameterSet{
			Directory: "/once/upon/a/time",
			Output:    "foo-bar",
			Format:    XmlFormatEn,
			Shape:     SubPathShapeEn,
			Concise:   true,
			Strategy:  TraverseLeafEn,
			Overwrite: false,
			Pattern:   "*infex*",
		}

		// This is the item that is generated via visiting the Cobra command's flags set
		// We don't want to force the client's cli app to use a map (generic parameter set),
		// instead we want to use a native (client defined) abstraction, so turn it into
		// a parameter set, ie:
		// --> FooParameterSet
		//
		Expected = adapters.GenericParamSet{
			"ID":        "expected",
			"Directory": "/once/upon/a/time",
			"Output":    "foo-bar",
			"Format":    XmlFormatEn,
			"Shape":     SubPathShapeEn,
			"Concise":   true,
			"Strategy":  TraverseLeafEn,
			"Overwrite": false,
			"Pattern":   "*infex*",
		}

		MissingDirectory = adapters.GenericParamSet{
			"ID":        "missing",
			"Output":    "foo-bar",
			"Format":    XmlFormatEn,
			"Shape":     SubPathShapeEn,
			"Concise":   true,
			"Strategy":  TraverseLeafEn,
			"Overwrite": false,
			"Pattern":   "*infex*",
		}
	})

	Describe("Extract", func() {
		var rootCommand *cobra.Command
		var fooCommand *cobra.Command

		local := new(Cradle)

		BeforeEach(func() {

			rootCommand = &cobra.Command{
				Use:   "poke",
				Short: "A brief description of your application",
				Long:  "A long description of the root poke command",
			}

			fooCommand = &cobra.Command{
				Version: "1.0.1",
				Use:     "infex",
				Short:   "Create infexion",
				Long:    "Index file system at root: '/'",
				Args:    cobra.ExactArgs(1), // this is for Directory
				RunE: func(command *cobra.Command, args []string) error {

					adapters.InsertSafePositional(MissingDirectory, "Directory", args[0])
					// factory = &adapters.ParamSetFactory[FooParameterSet]{
					// 	// Params: MissingDirectory,
					// }
					factory = new(adapters.ParamSetFactory[FooParameterSet])

					return nil
				},
			}

			fooCommand.Flags().StringP("output", "o", "./", "output location")
			fooCommand.Flags().StringP("format", "f", "xml", "write format")
			fooCommand.Flags().StringP("shape", "s", "sp", "path shape")
			fooCommand.Flags().BoolP("concise", "c", false, "minimise output")
			fooCommand.Flags().StringP("strategy", "y", "a", "traversal strategy")
			fooCommand.Flags().BoolP("overwrite", "x", false, "overwrite file if exists")
			fooCommand.Flags().StringP("pattern", "p", "*infex*", "pattern of files to load")

			// This is the way we should do it:
			// For int based enums, we need a generic specialisation that use a string valued flag
			//
			fooCommand.Flags().StringVarP(&local.Name, "name", "n", "default-value", "name usage")
			fooCommand.Flags().IntVarP(&local.Size, "size", "z", 1, "size usage")
			fooCommand.Flags().Float64VarP(&local.Temp, "temp", "t", 1.2, "temperature usage")

			rootCommand.AddCommand(fooCommand)
		})

		Context("given: Visit", func() {
			It("🧪 should: extract only flags that were supplied in command line", Label("spot"), func() {

				testhelpers.ExecuteCommand(
					rootCommand, "infex", "/usr/fuse/home/music", "--pattern=*music.infex*", "--name=jasmine", "--size=42", "--temp=3.14",
				)

				actual := factory.Extract(*fooCommand.Flags(), adapters.ExtractPsOptions{
					Strategy: adapters.ActiveFlagsEn,
				})
				fmt.Printf("*** 🤖🤖🤖 Name: '%v'\n", local.Name)
				fmt.Printf("*** 🤖🤖🤖 Size: '%v'\n", local.Size)
				fmt.Printf("*** 🤖🤖🤖 Temp: '%v'\n", local.Temp)

				Expect(actual).ToNot(BeNil())
				// Expect(actual.Pattern).To(Equal("*music.infex*"))
			})
		})

		// Context("given: VisitAll", func() {
		// 	It("🧪 should: ???", func() {
		// 		Skip("NOT-IMPLEMENTED-YET")

		// 		factory.Extract(*fooCommand.Flags(), adapters.ExtractPsOptions{
		// 			Strategy: adapters.AllFlagsEn,
		// 		})

		// 	})
		// })
	})

	Describe("Create", func() {
		Context("given: generic params created with standard types", func() {
			// This is a fake scenario because in reality, we will only
			// be working with values created by cobra, not faking them up
			// with 'standard' type, eg bool, int etc
			//
			When("no create options specified", func() {
				BeforeEach(func() {
					factory = &adapters.ParamSetFactory[FooParameterSet]{
						Params: Expected,
					}
				})

				It("🧪 should: create native param set", func() {
					actual := factory.Create()

					GinkgoWriter.Printf("=== RUNNING for directory '%v'\n", Expected["Directory"])
					Expect(reflect.DeepEqual(*actual, Native)).To(BeTrue())
				})

				Context("assert via reflection", func() {
					It("🧪 should: create native param set (assert via reflection)", func() {
						//
						// This test case is an excercise in using reflection and thus would not
						// normally be tested this way. See the other test above.
						//
						// based on the laws of reflection (https://go.dev/blog/laws-of-reflection)

						actual := factory.Create()

						refElemStruct := reflect.ValueOf(actual).Elem()
						refTypeOfStruct := refElemStruct.Type()

						if reflect.TypeOf(*actual).Kind() == reflect.Struct {
							for i, n := 0, refElemStruct.NumField(); i < n; i++ {
								name := refTypeOfStruct.Field(i).Name
								value := Expected[name]

								Expect(Expected[name]).To(Equal(value))
							}
						}
					})
				})
			})

			When("create options specified", func() {
				// Cobra represents values as strings
				//
				DescribeTable("given: not strict ???",
					func(entry PsEntry) {
						GinkgoWriter.Printf("=== 🍉 RUNNING for directory '%v'\n", Expected["Directory"])

						factory = &adapters.ParamSetFactory[FooParameterSet]{
							Params: entry.Params,
						}
						actual := factory.CreateWith(adapters.CreatePsOptions{Strict: entry.Strict})

						Expect(actual).ToNot(BeNil())
					},
					func(entry PsEntry) string {
						return fmt.Sprintf("🧪 --> 🍒 should: extract when '%v'", entry.Message)
					},

					Entry(nil, PsEntry{Params: MissingDirectory, Strict: false, Message: "missing entry"}),
					Entry(nil, PsEntry{Params: Expected, Strict: false, Message: "all present"}),
				)
			})
		})
	})

	Context("given: generic params created cobra values", func() {
		// Cobra represents values as strings
		//
		BeforeEach(func() {
			factory = &adapters.ParamSetFactory[FooParameterSet]{
				Params: Expected,
			}
			GinkgoWriter.Printf("=== RUNNING for directory '%v'\n", Expected["Directory"])
		})

		When("no create options specified", func() {
			It("🧪 should: create native param set", func() {
				actual := factory.Create()

				Expect(reflect.DeepEqual(*actual, Native)).To(BeTrue())
			})
		})

		When("create options specified", func() {
			It("should: create native param set", func() {
				actual := factory.CreateWith(adapters.CreatePsOptions{Strict: true})

				Expect(reflect.DeepEqual(*actual, Native)).To(BeTrue())
			})
		})
	})

	DescribeTable("given: invalid param set",
		Label("current"),
		func(entry PsEntry) {
			defer entry.PanicFn()
			GinkgoWriter.Printf("*** 👹 RUNNING for directory '%v'\n", entry.Params["Directory"])

			factory = &adapters.ParamSetFactory[FooParameterSet]{
				Params: entry.Params,
			}
			factory.CreateWith(adapters.CreatePsOptions{Strict: true})
			Fail("🔥 what no panic?")
		},
		func(entry PsEntry) string {
			return fmt.Sprintf("🧪 --> 🍄 should: panic when '%v'\n", entry.Message)
		},

		Entry(nil,
			PsEntry{Params: MissingDirectory, Message: "strict and missing entry",
				PanicFn: func() {
					if r := recover(); r != nil {
						recovery := fmt.Sprintf("%v\n", r)

						pattern := adapters.ErrorTypes[adapters.MissingNativeMemberValueEn].Info[adapters.PatternEn]
						Expect(recovery).To(MatchRegexp(pattern),
							fmt.Sprintf("panic message: '%v' did not conform to format", recovery))
					}
				},
			},
		),
	)
})
