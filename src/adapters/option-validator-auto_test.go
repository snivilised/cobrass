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
			Message: "time.Duration type (auto)",
			Setup: func() {
				paramSet.Native.Latency = duration("300ms")
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedDuration(
					adapters.NewFlagInfo("latency", "l", duration("0ms")),
					&paramSet.Native.Latency,
					func(value time.Duration) error {
						expect := duration("300ms")
						Expect(value).To(BeEquivalentTo(expect))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "[]time.Duration type (auto)",
			Setup: func() {
				paramSet.Native.Latencies = []time.Duration{duration("1s"), duration("2s"), duration("3s")}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedDurationSlice(
					adapters.NewFlagInfo("Latencies", "L", []time.Duration{}),
					&paramSet.Native.Latencies,
					func(value []time.Duration) error {
						Expect(value).To(BeEquivalentTo([]time.Duration{duration("1s"), duration("2s"), duration("3s")}))
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
				paramSet.Native.Gradientf32 = float32(32.1234)
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedFloat32(
					adapters.NewFlagInfo("gradientf32", "t", float32(0)),
					&paramSet.Native.Gradientf32,
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
				paramSet.Native.Gradientsf32 = []float32{2.99, 4.99, 6.99, 8.99}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedFloat32Slice(
					adapters.NewFlagInfo("Gradientsf32", "G", []float32{}),
					&paramSet.Native.Gradientsf32,
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
				paramSet.Native.Gradientf64 = float64(64.1234)
			},
			Validator: func() adapters.OptionValidator {

				return paramSet.BindValidatedFloat64(
					adapters.NewFlagInfo("gradientf64", "t", float64(0)),
					&paramSet.Native.Gradientf64,
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
				paramSet.Native.Gradientsf64 = []float64{3.99, 5.99, 7.99, 9.99}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedFloat64Slice(
					adapters.NewFlagInfo("Gradientsf64", "G", []float64{}),
					&paramSet.Native.Gradientsf64,
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
				paramSet.Native.Offsets = []int{2, 4, 6, 8}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedIntSlice(
					adapters.NewFlagInfo("Offsets", "D", []int{}),
					&paramSet.Native.Offsets,
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
			Message: "[]int32 type (auto)",
			Setup: func() {
				paramSet.Native.Offsets32 = []int32{2, 4, 6, 8}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedInt32Slice(
					adapters.NewFlagInfo("Offsets32", "O", []int32{}),
					&paramSet.Native.Offsets32,
					func(value []int32) error {
						Expect(value).To(Equal([]int32{2, 4, 6, 8}))
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
			Message: "[]int64 type (auto)",
			Setup: func() {
				paramSet.Native.Offsets64 = []int64{2, 4, 6, 8}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedInt64Slice(
					adapters.NewFlagInfo("Offsets64", "O", []int64{}),
					&paramSet.Native.Offsets64,
					func(value []int64) error {
						Expect(value).To(Equal([]int64{2, 4, 6, 8}))
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
				paramSet.Native.Directories = []string{"alpha", "beta", "delta"}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedStringSlice(
					adapters.NewFlagInfo("Directories", "C", []string{}),
					&paramSet.Native.Directories,
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
					adapters.NewFlagInfo("count16", "c", uint16(0)),
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
					adapters.NewFlagInfo("count32", "c", uint32(0)),
					&paramSet.Native.Count32,
					func(value uint32) error {
						Expect(value).To(Equal(uint32(3333)))
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
					adapters.NewFlagInfo("count64", "c", uint64(0)),
					&paramSet.Native.Count64,
					func(value uint64) error {
						Expect(value).To(Equal(uint64(33333)))
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
					adapters.NewFlagInfo("count8", "c", uint8(0)),
					&paramSet.Native.Count8,
					func(value uint8) error {
						Expect(value).To(Equal(uint8(33)))
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
				paramSet.Native.Counts = []uint{2, 4, 6, 8}
			},
			Validator: func() adapters.OptionValidator {
				return paramSet.BindValidatedUintSlice(
					adapters.NewFlagInfo("Counts", "P", []uint{}),
					&paramSet.Native.Counts,
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
