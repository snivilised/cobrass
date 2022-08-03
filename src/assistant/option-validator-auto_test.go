package assistant_test

import (
	"fmt"
	"net"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type OvEntry struct {
	Message   string
	Validator func() assistant.OptionValidator
	Setup     func()
}

var _ = Describe("OptionValidator", func() {

	var rootCommand *cobra.Command
	var widgetCommand *cobra.Command
	var paramSet *assistant.ParamSet[WidgetParameterSet]
	var outputFormatEnumInfo *assistant.EnumInfo[OutputFormatEnum]

	BeforeEach(func() {
		outputFormatEnumInfo = assistant.NewEnumInfo(AcceptableOutputFormats)

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

		paramSet = assistant.NewParamSet[WidgetParameterSet](widgetCommand)
	})

	DescribeTable("ParamSet with validation",
		func(entry OvEntry) {
			validator := entry.Validator()
			entry.Setup()
			_ = validator.Validate()
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
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedDuration(
					assistant.NewFlagInfo("latency", "l", duration("0ms")),
					&paramSet.Native.Latency,
					func(value time.Duration, flag *pflag.Flag) error {
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
			Validator: func() assistant.OptionValidator {
				return paramSet.BindValidatedDurationSlice(
					assistant.NewFlagInfo("Latencies", "L", []time.Duration{}),
					&paramSet.Native.Latencies,
					func(value []time.Duration, flag *pflag.Flag) error {
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
			Validator: func() assistant.OptionValidator {
				outputFormatEnum := outputFormatEnumInfo.NewValue()
				return paramSet.BindValidatedEnum(
					assistant.NewFlagInfo("format", "f", "xml"),
					&outputFormatEnum.Source,
					func(value string, flag *pflag.Flag) error {
						Expect(value).To(Equal("xml"))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "float32 type (auto)",
			Setup: func() {
				paramSet.Native.Gradientf32 = float32(float32(32.0))
			},
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedFloat32(
					assistant.NewFlagInfo("gradientf32", "t", float32(0)),
					&paramSet.Native.Gradientf32,
					func(value float32, flag *pflag.Flag) error {
						Expect(value).To(Equal(float32(float32(32.0))))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "[]float32 type (auto)",
			Setup: func() {
				paramSet.Native.Gradientsf32 = []float32{3.0, 5.0, 7.0, 9.0}
			},
			Validator: func() assistant.OptionValidator {
				return paramSet.BindValidatedFloat32Slice(
					assistant.NewFlagInfo("Gradientsf32", "G", []float32{}),
					&paramSet.Native.Gradientsf32,
					func(value []float32, flag *pflag.Flag) error {
						Expect(value).To(Equal([]float32{3.0, 5.0, 7.0, 9.0}))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "float64 type (auto)",
			Setup: func() {
				paramSet.Native.Gradientf64 = float64(float64(64.1234))
			},
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedFloat64(
					assistant.NewFlagInfo("gradientf64", "t", float64(0)),
					&paramSet.Native.Gradientf64,
					func(value float64, flag *pflag.Flag) error {
						Expect(value).To(Equal(float64(float64(64.1234))))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "[]float64 type (auto)",
			Setup: func() {
				paramSet.Native.Gradientsf64 = []float64{4.0, 6.0, 8.0, 10.0}
			},
			Validator: func() assistant.OptionValidator {
				return paramSet.BindValidatedFloat64Slice(
					assistant.NewFlagInfo("Gradientsf64", "G", []float64{}),
					&paramSet.Native.Gradientsf64,
					func(value []float64, flag *pflag.Flag) error {
						Expect(value).To(Equal([]float64{4.0, 6.0, 8.0, 10.0}))
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
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedInt(
					assistant.NewFlagInfo("offset", "o", -1),
					&paramSet.Native.Offset,
					func(value int, flag *pflag.Flag) error {
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
			Validator: func() assistant.OptionValidator {
				return paramSet.BindValidatedIntSlice(
					assistant.NewFlagInfo("Offsets", "D", []int{}),
					&paramSet.Native.Offsets,
					func(value []int, flag *pflag.Flag) error {
						Expect(value).To(Equal([]int{2, 4, 6, 8}))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "int16 type (auto)",
			Setup: func() {
				paramSet.Native.Offset16 = int16(int16(-999))
			},
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedInt16(
					assistant.NewFlagInfo("offset16", "o", int16(-1)),
					&paramSet.Native.Offset16,
					func(value int16, flag *pflag.Flag) error {
						Expect(value).To(Equal(int16(int16(-999))))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "int32 type (auto)",
			Setup: func() {
				paramSet.Native.Offset32 = int32(int32(-9999))
			},
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedInt32(
					assistant.NewFlagInfo("offset32", "o", int32(-1)),
					&paramSet.Native.Offset32,
					func(value int32, flag *pflag.Flag) error {
						Expect(value).To(Equal(int32(int32(-9999))))
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
			Validator: func() assistant.OptionValidator {
				return paramSet.BindValidatedInt32Slice(
					assistant.NewFlagInfo("Offsets32", "O", []int32{}),
					&paramSet.Native.Offsets32,
					func(value []int32, flag *pflag.Flag) error {
						Expect(value).To(Equal([]int32{2, 4, 6, 8}))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "int64 type (auto)",
			Setup: func() {
				paramSet.Native.Offset64 = int64(int64(-99999))
			},
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedInt64(
					assistant.NewFlagInfo("offset64", "o", int64(-1)),
					&paramSet.Native.Offset64,
					func(value int64, flag *pflag.Flag) error {
						Expect(value).To(Equal(int64(int64(-99999))))
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
			Validator: func() assistant.OptionValidator {
				return paramSet.BindValidatedInt64Slice(
					assistant.NewFlagInfo("Offsets64", "O", []int64{}),
					&paramSet.Native.Offsets64,
					func(value []int64, flag *pflag.Flag) error {
						Expect(value).To(Equal([]int64{2, 4, 6, 8}))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "int8 type (auto)",
			Setup: func() {
				paramSet.Native.Offset8 = int8(int8(-99))
			},
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedInt8(
					assistant.NewFlagInfo("offset8", "o", int8(-1)),
					&paramSet.Native.Offset8,
					func(value int8, flag *pflag.Flag) error {
						Expect(value).To(Equal(int8(int8(-99))))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "net.IPMask type (auto)",
			Setup: func() {
				paramSet.Native.IpMask = ipmask("orion.net")
			},
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedIPMask(
					assistant.NewFlagInfo("ipmask", "m", ipmask("default")),
					&paramSet.Native.IpMask,
					func(value net.IPMask, flag *pflag.Flag) error {
						Expect(value).To(BeEquivalentTo(ipmask("orion.net")))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "net.IPNet type (auto)",
			Setup: func() {
				paramSet.Native.IpAddress = ipnet("orion.net")
			},
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedIPNet(
					assistant.NewFlagInfo("ipaddress", "i", ipnet("default")),
					&paramSet.Native.IpAddress,
					func(value net.IPNet, flag *pflag.Flag) error {
						Expect(value).To(BeEquivalentTo(ipnet("orion.net")))
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
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedString(
					assistant.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern,
					func(value string, flag *pflag.Flag) error {
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
			Validator: func() assistant.OptionValidator {
				return paramSet.BindValidatedStringSlice(
					assistant.NewFlagInfo("Directories", "C", []string{}),
					&paramSet.Native.Directories,
					func(value []string, flag *pflag.Flag) error {
						Expect(value).To(Equal([]string{"alpha", "beta", "delta"}))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "uint16 type (auto)",
			Setup: func() {
				paramSet.Native.Count16 = uint16(uint16(333))
			},
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedUint16(
					assistant.NewFlagInfo("count16", "c", uint16(0)),
					&paramSet.Native.Count16,
					func(value uint16, flag *pflag.Flag) error {
						Expect(value).To(Equal(uint16(uint16(333))))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "uint32 type (auto)",
			Setup: func() {
				paramSet.Native.Count32 = uint32(uint32(3333))
			},
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedUint32(
					assistant.NewFlagInfo("count32", "c", uint32(0)),
					&paramSet.Native.Count32,
					func(value uint32, flag *pflag.Flag) error {
						Expect(value).To(Equal(uint32(uint32(3333))))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "uint64 type (auto)",
			Setup: func() {
				paramSet.Native.Count64 = uint64(uint64(33333))
			},
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedUint64(
					assistant.NewFlagInfo("count64", "c", uint64(0)),
					&paramSet.Native.Count64,
					func(value uint64, flag *pflag.Flag) error {
						Expect(value).To(Equal(uint64(uint64(33333))))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "uint8 type (auto)",
			Setup: func() {
				paramSet.Native.Count8 = uint8(uint8(33))
			},
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedUint8(
					assistant.NewFlagInfo("count8", "c", uint8(0)),
					&paramSet.Native.Count8,
					func(value uint8, flag *pflag.Flag) error {
						Expect(value).To(Equal(uint8(uint8(33))))
						return nil
					},
				)
			},
		}),

		Entry(nil, OvEntry{
			Message: "uint type (auto)",
			Setup: func() {
				paramSet.Native.Count = uint(uint(99999))
			},
			Validator: func() assistant.OptionValidator {

				return paramSet.BindValidatedUint(
					assistant.NewFlagInfo("count", "c", uint(0)),
					&paramSet.Native.Count,
					func(value uint, flag *pflag.Flag) error {
						Expect(value).To(Equal(uint(uint(99999))))
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
			Validator: func() assistant.OptionValidator {
				return paramSet.BindValidatedUintSlice(
					assistant.NewFlagInfo("Counts", "P", []uint{}),
					&paramSet.Native.Counts,
					func(value []uint, flag *pflag.Flag) error {
						Expect(value).To(Equal([]uint{2, 4, 6, 8}))
						return nil
					},
				)
			},
		}),

		// <---- end of auto generated
	)
})
