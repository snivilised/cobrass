package adapters_test

import (
	"fmt"
	"net"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	"github.com/snivilised/cobrass/src/adapters"
)

type OvEntry struct {
	Message   string
	Validator func() adapters.OptionValidator
	Setup     func()
}

var _ = Describe("OptionValidator", func() {

	var rootCommand *cobra.Command
	var widgetCommand *cobra.Command
	var paramSet *adapters.ParamSet[WidgetParameterSet]
	var outputFormatEnumInfo *adapters.EnumInfo[OutputFormatEnum]

	BeforeEach(func() {
		outputFormatEnumInfo = adapters.NewEnumInfo(adapters.AcceptableEnumValues[OutputFormatEnum]{
			XmlFormatEn:      []string{"xml", "x"},
			JsonFormatEn:     []string{"json", "j"},
			TextFormatEn:     []string{"text", "tx"},
			ScribbleFormatEn: []string{"scribble", "scribbler", "scr"},
		})

		rootCommand = &cobra.Command{
			Use:   "poke",
			Short: "A brief description of your application",
			Long:  "A long description of the root poke command",
		}

		widgetCommand = &cobra.Command{
			Version: "1.0.1",
			Use:     "widget",
			Short:   "Create widget",
			Long:    "Index file system at root: '/'",
			Args:    cobra.ExactArgs(1),

			RunE: func(command *cobra.Command, args []string) error {
				GinkgoWriter.Printf("===> 🍓 EXECUTE (Directory: '%v')\n", args[0])

				paramSet.Native.Directory = args[0]
				return nil
			},
		}
		rootCommand.AddCommand(widgetCommand)

		paramSet = adapters.NewParamSet[WidgetParameterSet](widgetCommand)
	})

	DescribeTable("ParamSet with validation",
		func(entry OvEntry) {
			validator := entry.Validator()
			entry.Setup()
			validator.Validate()
		},
		func(entry OvEntry) string {
			return fmt.Sprintf("🧪 --> 🍒 given: flag type is '%v'", entry.Message)
		},

		// ----> auto generated(Build-TestEntry/gen-ov-t)

		Entry(nil, OvEntry{
			Message: "bool type (auto)",
			Setup: func() {
				paramSet.Native.Concise = true
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedBool(
					adapters.NewFlagInfo("concise", "c", false),
					&paramSet.Native.Concise,
					func(value bool) error {
						Expect(value).To(BeTrue())
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "[]bool type (auto)",
			Setup: func() {
				paramSet.Native.Switches = []bool{true, false, true, false}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedBoolSlice(
					adapters.NewFlagInfo("Switches", "S", []bool{}),
					&paramSet.Native.Switches,
					func(value []bool) error {
						Expect(value).To(Equal([]bool{true, false, true, false}))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "time.Duration type (auto)",
			Setup: func() {
				paramSet.Native.Latency, _ = time.ParseDuration("300ms")
			},
			Validator: func() adapters.OptionValidator {
				temp, _ := time.ParseDuration("0ms")
				return paramSet.BindValidatedDuration(
					adapters.NewFlagInfo("latency", "l", temp),
					&paramSet.Native.Latency,
					func(value time.Duration) error {
						expect, _ := time.ParseDuration("300ms")
						Expect(value).To(BeEquivalentTo(expect))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "string type (auto)",
			Setup: func() {
				paramSet.Native.Format = XmlFormatEn
			},
			Validator: func() adapters.OptionValidator {
				outputFormatEnum := outputFormatEnumInfo.NewValue()
				return paramSet.BindValidatedEnum(
					adapters.NewFlagInfo("format", "f", "xml"),
					&outputFormatEnum.Source,
					func(value string) error {
						Expect(value).To(Equal("xml"))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "float32 type (auto)",
			Setup: func() {
				paramSet.Native.Gradient = float32(32.1234)
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedFloat32(
					adapters.NewFlagInfo("gradient", "t", float32(999.123)),
					&paramSet.Native.Gradient,
					func(value float32) error {
						Expect(value).To(Equal(float32(32.1234)))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "[]float32 type (auto)",
			Setup: func() {
				paramSet.Native.Temperatures = []float32{2.99, 4.99, 6.99, 8.99}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedFloat32Slice(
					adapters.NewFlagInfo("Temperatures", "T", []float32{}),
					&paramSet.Native.Temperatures,
					func(value []float32) error {
						Expect(value).To(Equal([]float32{2.99, 4.99, 6.99, 8.99}))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "float64 type (auto)",
			Setup: func() {
				paramSet.Native.Threshold = float64(64.1234)
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedFloat64(
					adapters.NewFlagInfo("threshold", "t", float64(999.123)),
					&paramSet.Native.Threshold,
					func(value float64) error {
						Expect(value).To(Equal(float64(64.1234)))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "[]float64 type (auto)",
			Setup: func() {
				paramSet.Native.Scales = []float64{3.99, 5.99, 7.99, 9.99}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedFloat64Slice(
					adapters.NewFlagInfo("Scales", "S", []float64{}),
					&paramSet.Native.Scales,
					func(value []float64) error {
						Expect(value).To(Equal([]float64{3.99, 5.99, 7.99, 9.99}))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "int type (auto)",
			Setup: func() {
				paramSet.Native.Offset = -9
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedInt(
					adapters.NewFlagInfo("offset", "o", -1),
					&paramSet.Native.Offset,
					func(value int) error {
						Expect(value).To(Equal(-9))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "[]int type (auto)",
			Setup: func() {
				paramSet.Native.Dimensions = []int{2, 4, 6, 8}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedIntSlice(
					adapters.NewFlagInfo("Dimensions", "D", []int{}),
					&paramSet.Native.Dimensions,
					func(value []int) error {
						Expect(value).To(Equal([]int{2, 4, 6, 8}))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "int16 type (auto)",
			Setup: func() {
				paramSet.Native.Offset16 = int16(-999)
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedInt16(
					adapters.NewFlagInfo("offset16", "o", int16(-1)),
					&paramSet.Native.Offset16,
					func(value int16) error {
						Expect(value).To(Equal(int16(-999)))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "int32 type (auto)",
			Setup: func() {
				paramSet.Native.Offset32 = int32(-9999)
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedInt32(
					adapters.NewFlagInfo("offset32", "o", int32(-1)),
					&paramSet.Native.Offset32,
					func(value int32) error {
						Expect(value).To(Equal(int32(-9999)))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "int64 type (auto)",
			Setup: func() {
				paramSet.Native.Offset64 = int64(-99999)
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedInt64(
					adapters.NewFlagInfo("offset64", "o", int64(-1)),
					&paramSet.Native.Offset64,
					func(value int64) error {
						Expect(value).To(Equal(int64(-99999)))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "int8 type (auto)",
			Setup: func() {
				paramSet.Native.Offset8 = int8(-99)
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedInt8(
					adapters.NewFlagInfo("offset8", "o", int8(-1)),
					&paramSet.Native.Offset8,
					func(value int8) error {
						Expect(value).To(Equal(int8(-99)))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "net.IPMask type (auto)",
			Setup: func() {
				paramSet.Native.IpMask = net.IPMask([]byte{255, 255, 255, 0})
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedIPMask(
					adapters.NewFlagInfo("ipmask", "m", net.IPMask([]byte{0, 0, 0, 0})),
					&paramSet.Native.IpMask,
					func(value net.IPMask) error {
						Expect(value).To(BeEquivalentTo(net.IPMask([]byte{255, 255, 255, 0})))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "net.IPNet type (auto)",
			Setup: func() {
				paramSet.Native.IpAddress = net.IPNet{IP: net.IPv4(172, 16, 0, 0), Mask: net.IPMask([]byte{255, 255, 255, 0})}
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedIPNet(
					adapters.NewFlagInfo("ipaddress", "i", net.IPNet{IP: net.IPv4(0, 0, 0, 0), Mask: net.IPMask([]byte{0, 0, 0, 0})}),
					&paramSet.Native.IpAddress,
					func(value net.IPNet) error {
						Expect(value).To(BeEquivalentTo(net.IPNet{IP: net.IPv4(172, 16, 0, 0), Mask: net.IPMask([]byte{255, 255, 255, 0})}))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "string type (auto)",
			Setup: func() {
				paramSet.Native.Pattern = "*music.infex*"
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedString(
					adapters.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern,
					func(value string) error {
						Expect(value).To(Equal("*music.infex*"))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "[]string type (auto)",
			Setup: func() {
				paramSet.Native.Categories = []string{"alpha", "beta", "delta"}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedStringSlice(
					adapters.NewFlagInfo("Categories", "C", []string{}),
					&paramSet.Native.Categories,
					func(value []string) error {
						Expect(value).To(Equal([]string{"alpha", "beta", "delta"}))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "uint16 type (auto)",
			Setup: func() {
				paramSet.Native.Count16 = uint16(333)
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedUint16(
					adapters.NewFlagInfo("count16", "c", uint16(1)),
					&paramSet.Native.Count16,
					func(value uint16) error {
						Expect(value).To(Equal(uint16(333)))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "uint32 type (auto)",
			Setup: func() {
				paramSet.Native.Count32 = uint32(3333)
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedUint32(
					adapters.NewFlagInfo("count32", "c", uint32(1)),
					&paramSet.Native.Count32,
					func(value uint32) error {
						Expect(value).To(Equal(uint32(3333)))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "uint8 type (auto)",
			Setup: func() {
				paramSet.Native.Count8 = uint8(33)
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedUint8(
					adapters.NewFlagInfo("count8", "c", uint8(1)),
					&paramSet.Native.Count8,
					func(value uint8) error {
						Expect(value).To(Equal(uint8(33)))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "uint64 type (auto)",
			Setup: func() {
				paramSet.Native.Count64 = uint64(33333)
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedUint64(
					adapters.NewFlagInfo("count64", "c", uint64(1)),
					&paramSet.Native.Count64,
					func(value uint64) error {
						Expect(value).To(Equal(uint64(33333)))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "uint type (auto)",
			Setup: func() {
				paramSet.Native.Count = uint(99999)
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedUint(
					adapters.NewFlagInfo("count", "c", uint(0)),
					&paramSet.Native.Count,
					func(value uint) error {
						Expect(value).To(Equal(uint(99999)))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "[]uint type (auto)",
			Setup: func() {
				paramSet.Native.Points = []uint{2, 4, 6, 8}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedUintSlice(
					adapters.NewFlagInfo("Points", "P", []uint{}),
					&paramSet.Native.Points,
					func(value []uint) error {
						Expect(value).To(Equal([]uint{2, 4, 6, 8}))
						return nil
					},
				)
			},
		}),

		// <---- end of auto generated
	)
})
