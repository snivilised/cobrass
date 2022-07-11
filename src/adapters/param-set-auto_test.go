package adapters_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/snivilised/cobrass/src/adapters"
	"github.com/snivilised/cobrass/src/testhelpers"
	"github.com/spf13/cobra"
)

// the auto version of param-set_test.go

var _ = Describe("ParamSet (auto)", func() {

	When("Binding a flag (auto)", func() {
		var rootCommand *cobra.Command
		var widgetCommand *cobra.Command
		var paramSet *adapters.ParamSet[WidgetParameterSet]
		var outputFormatEnumInfo *adapters.EnumInfo[OutputFormatEnum]
		var outputFormatEnum adapters.EnumValue[OutputFormatEnum]

		BeforeEach(func() {
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

				PreRun: func(command *cobra.Command, args []string) {
					GinkgoWriter.Printf("**** 🍉 PRE-RUN\n")
				},
				RunE: func(command *cobra.Command, args []string) error {
					GinkgoWriter.Printf("===> 🍓 EXECUTE (Directory: '%v')\n", args[0])

					paramSet.Native.Directory = args[0]
					return nil
				},
				PostRun: func(command *cobra.Command, args []string) {
					GinkgoWriter.Printf("**** 🥥 POST-RUN\n")
				},
			}
			rootCommand.AddCommand(widgetCommand)

			paramSet = adapters.NewParamSet[WidgetParameterSet](widgetCommand)

			outputFormatEnumInfo = adapters.NewEnumInfo(AcceptableOutputFormats)
			outputFormatEnum = outputFormatEnumInfo.NewValue()
		})

		DescribeTable("binder",
			func(entry TcEntry) {
				entry.Binder()

				testhelpers.ExecuteCommand(
					rootCommand, "widget", "/usr/fuse/home/music", entry.CommandLine,
				)
				entry.Assert()
			},

			func(entry TcEntry) string {
				return fmt.Sprintf("🧪 --> 🍒 given: flag is '%v'", entry.Message)
			},

			// ----> auto generated(Build-PsTestEntry/gen-ps-t)

			Entry(nil, TcEntry{
				Message: "bool type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindBool(
						adapters.NewFlagInfo("concise", "c", false),
						&paramSet.Native.Concise,
					)
				},
				CommandLine: "--concise",
				Assert:      func() { Expect(paramSet.Native.Concise).To(Equal(true)) },
			}),

			Entry(nil, TcEntry{
				Message: "[]bool slice type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindBoolSlice(
						adapters.NewFlagInfo("switches", "S", []bool{}),
						&paramSet.Native.Switches,
					)
				},
				CommandLine: "--switches=true,false,true,false",
				Assert:      func() { Expect(paramSet.Native.Switches).To(BeEquivalentTo([]bool{true, false, true, false})) },
			}),

			Entry(nil, TcEntry{
				Message: "bool type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindBool(
						adapters.NewFlagInfo("concise", "", false),
						&paramSet.Native.Concise,
					)
				},
				CommandLine: "--concise",
				Assert:      func() { Expect(paramSet.Native.Concise).To(Equal(true)) },
			}),

			Entry(nil, TcEntry{
				Message: "[]bool slice type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindBoolSlice(
						adapters.NewFlagInfo("switches", "", []bool{}),
						&paramSet.Native.Switches,
					)
				},
				CommandLine: "--switches=true,false,true,false",
				Assert:      func() { Expect(paramSet.Native.Switches).To(BeEquivalentTo([]bool{true, false, true, false})) },
			}),

			Entry(nil, TcEntry{
				Message: "time.Duration type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindDuration(
						adapters.NewFlagInfo("latency", "l", duration("0ms")),
						&paramSet.Native.Latency,
					)
				},
				CommandLine: "--latency=300ms",
				Assert:      func() { Expect(paramSet.Native.Latency).To(BeEquivalentTo(duration("300ms"))) },
			}),

			Entry(nil, TcEntry{
				Message: "[]time.Duration slice type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindDurationSlice(
						adapters.NewFlagInfo("latencies", "L", []time.Duration{}),
						&paramSet.Native.Latencies,
					)
				},
				CommandLine: "--latencies=1s,2s,3s",
				Assert: func() {
					Expect(paramSet.Native.Latencies).To(BeEquivalentTo([]time.Duration{duration("1s"), duration("2s"), duration("3s")}))
				},
			}),

			Entry(nil, TcEntry{
				Message: "time.Duration type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindDuration(
						adapters.NewFlagInfo("latency", "", duration("0ms")),
						&paramSet.Native.Latency,
					)
				},
				CommandLine: "--latency=300ms",
				Assert:      func() { Expect(paramSet.Native.Latency).To(BeEquivalentTo(duration("300ms"))) },
			}),

			Entry(nil, TcEntry{
				Message: "[]time.Duration slice type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindDurationSlice(
						adapters.NewFlagInfo("latencies", "", []time.Duration{}),
						&paramSet.Native.Latencies,
					)
				},
				CommandLine: "--latencies=1s,2s,3s",
				Assert: func() {
					Expect(paramSet.Native.Latencies).To(BeEquivalentTo([]time.Duration{duration("1s"), duration("2s"), duration("3s")}))
				},
			}),

			Entry(nil, TcEntry{
				Message: "enum type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindEnum(
						adapters.NewFlagInfo("format", "f", "xml"),
						&outputFormatEnum.Source,
					)
				},
				CommandLine: "--format=json",
				Assert:      func() { Expect(outputFormatEnum.Source).To(Equal("json")) },
			}),

			Entry(nil, TcEntry{
				Message: "enum type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindEnum(
						adapters.NewFlagInfo("format", "", "xml"),
						&outputFormatEnum.Source,
					)
				},
				CommandLine: "--format=json",
				Assert:      func() { Expect(outputFormatEnum.Source).To(Equal("json")) },
			}),

			Entry(nil, TcEntry{
				Message: "float32 type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindFloat32(
						adapters.NewFlagInfo("gradientf32", "t", float32(0)),
						&paramSet.Native.Gradientf32,
					)
				},
				CommandLine: "--gradientf32=32.0",
				Assert:      func() { Expect(paramSet.Native.Gradientf32).To(Equal(float32(32.0))) },
			}),

			Entry(nil, TcEntry{
				Message: "[]float32 slice type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindFloat32Slice(
						adapters.NewFlagInfo("gradientsf32", "G", []float32{}),
						&paramSet.Native.Gradientsf32,
					)
				},
				CommandLine: "--gradientsf32=3.0,5.0,7.0,9.0",
				Assert:      func() { Expect(paramSet.Native.Gradientsf32).To(BeEquivalentTo([]float32{3.0, 5.0, 7.0, 9.0})) },
			}),

			Entry(nil, TcEntry{
				Message: "float32 type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindFloat32(
						adapters.NewFlagInfo("gradientf32", "", float32(0)),
						&paramSet.Native.Gradientf32,
					)
				},
				CommandLine: "--gradientf32=32.0",
				Assert:      func() { Expect(paramSet.Native.Gradientf32).To(Equal(float32(32.0))) },
			}),

			Entry(nil, TcEntry{
				Message: "[]float32 slice type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindFloat32Slice(
						adapters.NewFlagInfo("gradientsf32", "", []float32{}),
						&paramSet.Native.Gradientsf32,
					)
				},
				CommandLine: "--gradientsf32=3.0,5.0,7.0,9.0",
				Assert:      func() { Expect(paramSet.Native.Gradientsf32).To(BeEquivalentTo([]float32{3.0, 5.0, 7.0, 9.0})) },
			}),

			Entry(nil, TcEntry{
				Message: "float64 type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindFloat64(
						adapters.NewFlagInfo("gradientf64", "t", float64(0)),
						&paramSet.Native.Gradientf64,
					)
				},
				CommandLine: "--gradientf64=64.1234",
				Assert:      func() { Expect(paramSet.Native.Gradientf64).To(Equal(float64(64.1234))) },
			}),

			Entry(nil, TcEntry{
				Message: "[]float64 slice type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindFloat64Slice(
						adapters.NewFlagInfo("gradientsf64", "G", []float64{}),
						&paramSet.Native.Gradientsf64,
					)
				},
				CommandLine: "--gradientsf64=4.0,6.0,8.0,10.0",
				Assert:      func() { Expect(paramSet.Native.Gradientsf64).To(BeEquivalentTo([]float64{4.0, 6.0, 8.0, 10.0})) },
			}),

			Entry(nil, TcEntry{
				Message: "float64 type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindFloat64(
						adapters.NewFlagInfo("gradientf64", "", float64(0)),
						&paramSet.Native.Gradientf64,
					)
				},
				CommandLine: "--gradientf64=64.1234",
				Assert:      func() { Expect(paramSet.Native.Gradientf64).To(Equal(float64(64.1234))) },
			}),

			Entry(nil, TcEntry{
				Message: "[]float64 slice type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindFloat64Slice(
						adapters.NewFlagInfo("gradientsf64", "", []float64{}),
						&paramSet.Native.Gradientsf64,
					)
				},
				CommandLine: "--gradientsf64=4.0,6.0,8.0,10.0",
				Assert:      func() { Expect(paramSet.Native.Gradientsf64).To(BeEquivalentTo([]float64{4.0, 6.0, 8.0, 10.0})) },
			}),

			Entry(nil, TcEntry{
				Message: "int type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindInt(
						adapters.NewFlagInfo("offset", "o", -1),
						&paramSet.Native.Offset,
					)
				},
				CommandLine: "--offset=-9",
				Assert:      func() { Expect(paramSet.Native.Offset).To(Equal(-9)) },
			}),

			Entry(nil, TcEntry{
				Message: "[]int slice type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindIntSlice(
						adapters.NewFlagInfo("offsets", "D", []int{}),
						&paramSet.Native.Offsets,
					)
				},
				CommandLine: "--offsets=2,4,6,8",
				Assert:      func() { Expect(paramSet.Native.Offsets).To(BeEquivalentTo([]int{2, 4, 6, 8})) },
			}),

			Entry(nil, TcEntry{
				Message: "int type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindInt(
						adapters.NewFlagInfo("offset", "", -1),
						&paramSet.Native.Offset,
					)
				},
				CommandLine: "--offset=-9",
				Assert:      func() { Expect(paramSet.Native.Offset).To(Equal(-9)) },
			}),

			Entry(nil, TcEntry{
				Message: "[]int slice type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindIntSlice(
						adapters.NewFlagInfo("offsets", "", []int{}),
						&paramSet.Native.Offsets,
					)
				},
				CommandLine: "--offsets=2,4,6,8",
				Assert:      func() { Expect(paramSet.Native.Offsets).To(BeEquivalentTo([]int{2, 4, 6, 8})) },
			}),

			Entry(nil, TcEntry{
				Message: "int16 type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindInt16(
						adapters.NewFlagInfo("offset16", "o", int16(-1)),
						&paramSet.Native.Offset16,
					)
				},
				CommandLine: "--offset16=-999",
				Assert:      func() { Expect(paramSet.Native.Offset16).To(Equal(int16(-999))) },
			}),

			Entry(nil, TcEntry{
				Message: "int16 type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindInt16(
						adapters.NewFlagInfo("offset16", "", int16(-1)),
						&paramSet.Native.Offset16,
					)
				},
				CommandLine: "--offset16=-999",
				Assert:      func() { Expect(paramSet.Native.Offset16).To(Equal(int16(-999))) },
			}),

			Entry(nil, TcEntry{
				Message: "int32 type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindInt32(
						adapters.NewFlagInfo("offset32", "o", int32(-1)),
						&paramSet.Native.Offset32,
					)
				},
				CommandLine: "--offset32=-9999",
				Assert:      func() { Expect(paramSet.Native.Offset32).To(Equal(int32(-9999))) },
			}),

			Entry(nil, TcEntry{
				Message: "[]int32 slice type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindInt32Slice(
						adapters.NewFlagInfo("offsets32", "O", []int32{}),
						&paramSet.Native.Offsets32,
					)
				},
				CommandLine: "--offsets32=2,4,6,8",
				Assert:      func() { Expect(paramSet.Native.Offsets32).To(BeEquivalentTo([]int32{2, 4, 6, 8})) },
			}),

			Entry(nil, TcEntry{
				Message: "int32 type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindInt32(
						adapters.NewFlagInfo("offset32", "", int32(-1)),
						&paramSet.Native.Offset32,
					)
				},
				CommandLine: "--offset32=-9999",
				Assert:      func() { Expect(paramSet.Native.Offset32).To(Equal(int32(-9999))) },
			}),

			Entry(nil, TcEntry{
				Message: "[]int32 slice type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindInt32Slice(
						adapters.NewFlagInfo("offsets32", "", []int32{}),
						&paramSet.Native.Offsets32,
					)
				},
				CommandLine: "--offsets32=2,4,6,8",
				Assert:      func() { Expect(paramSet.Native.Offsets32).To(BeEquivalentTo([]int32{2, 4, 6, 8})) },
			}),

			Entry(nil, TcEntry{
				Message: "int64 type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindInt64(
						adapters.NewFlagInfo("offset64", "o", int64(-1)),
						&paramSet.Native.Offset64,
					)
				},
				CommandLine: "--offset64=-99999",
				Assert:      func() { Expect(paramSet.Native.Offset64).To(Equal(int64(-99999))) },
			}),

			Entry(nil, TcEntry{
				Message: "[]int64 slice type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindInt64Slice(
						adapters.NewFlagInfo("offsets64", "O", []int64{}),
						&paramSet.Native.Offsets64,
					)
				},
				CommandLine: "--offsets64=2,4,6,8",
				Assert:      func() { Expect(paramSet.Native.Offsets64).To(BeEquivalentTo([]int64{2, 4, 6, 8})) },
			}),

			Entry(nil, TcEntry{
				Message: "int64 type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindInt64(
						adapters.NewFlagInfo("offset64", "", int64(-1)),
						&paramSet.Native.Offset64,
					)
				},
				CommandLine: "--offset64=-99999",
				Assert:      func() { Expect(paramSet.Native.Offset64).To(Equal(int64(-99999))) },
			}),

			Entry(nil, TcEntry{
				Message: "[]int64 slice type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindInt64Slice(
						adapters.NewFlagInfo("offsets64", "", []int64{}),
						&paramSet.Native.Offsets64,
					)
				},
				CommandLine: "--offsets64=2,4,6,8",
				Assert:      func() { Expect(paramSet.Native.Offsets64).To(BeEquivalentTo([]int64{2, 4, 6, 8})) },
			}),

			Entry(nil, TcEntry{
				Message: "int8 type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindInt8(
						adapters.NewFlagInfo("offset8", "o", int8(-1)),
						&paramSet.Native.Offset8,
					)
				},
				CommandLine: "--offset8=-99",
				Assert:      func() { Expect(paramSet.Native.Offset8).To(Equal(int8(-99))) },
			}),

			Entry(nil, TcEntry{
				Message: "int8 type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindInt8(
						adapters.NewFlagInfo("offset8", "", int8(-1)),
						&paramSet.Native.Offset8,
					)
				},
				CommandLine: "--offset8=-99",
				Assert:      func() { Expect(paramSet.Native.Offset8).To(Equal(int8(-99))) },
			}),

			Entry(nil, TcEntry{
				Message: "net.IPMask type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindIPMask(
						adapters.NewFlagInfo("ipmask", "m", ipmask("default")),
						&paramSet.Native.IpMask,
					)
				},
				CommandLine: "--ipmask=255.255.255.0",
				Assert:      func() { Expect(paramSet.Native.IpMask).To(BeEquivalentTo(ipmask("orion.net"))) },
			}),

			Entry(nil, TcEntry{
				Message: "net.IPMask type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindIPMask(
						adapters.NewFlagInfo("ipmask", "", ipmask("default")),
						&paramSet.Native.IpMask,
					)
				},
				CommandLine: "--ipmask=255.255.255.0",
				Assert:      func() { Expect(paramSet.Native.IpMask).To(BeEquivalentTo(ipmask("orion.net"))) },
			}),

			Entry(nil, TcEntry{
				Message: "net.IPNet type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindIPNet(
						adapters.NewFlagInfo("ipaddress", "i", ipnet("default")),
						&paramSet.Native.IpAddress,
					)
				},
				CommandLine: "--ipaddress=172.16.0.0",
				Assert:      func() { Expect(paramSet.Native.IpAddress).ToNot(BeNil()) },
			}),

			Entry(nil, TcEntry{
				Message: "net.IPNet type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindIPNet(
						adapters.NewFlagInfo("ipaddress", "", ipnet("default")),
						&paramSet.Native.IpAddress,
					)
				},
				CommandLine: "--ipaddress=172.16.0.0",
				Assert:      func() { Expect(paramSet.Native.IpAddress).ToNot(BeNil()) },
			}),

			Entry(nil, TcEntry{
				Message: "string type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindString(
						adapters.NewFlagInfo("pattern", "p", "default-pattern"),
						&paramSet.Native.Pattern,
					)
				},
				CommandLine: "--pattern=*music.infex*",
				Assert:      func() { Expect(paramSet.Native.Pattern).To(Equal("*music.infex*")) },
			}),

			Entry(nil, TcEntry{
				Message: "[]string slice type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindStringSlice(
						adapters.NewFlagInfo("directories", "C", []string{}),
						&paramSet.Native.Directories,
					)
				},
				CommandLine: "--directories=alpha,beta,delta",
				Assert:      func() { Expect(paramSet.Native.Directories).To(BeEquivalentTo([]string{"alpha", "beta", "delta"})) },
			}),

			Entry(nil, TcEntry{
				Message: "string type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindString(
						adapters.NewFlagInfo("pattern", "", "default-pattern"),
						&paramSet.Native.Pattern,
					)
				},
				CommandLine: "--pattern=*music.infex*",
				Assert:      func() { Expect(paramSet.Native.Pattern).To(Equal("*music.infex*")) },
			}),

			Entry(nil, TcEntry{
				Message: "[]string slice type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindStringSlice(
						adapters.NewFlagInfo("directories", "", []string{}),
						&paramSet.Native.Directories,
					)
				},
				CommandLine: "--directories=alpha,beta,delta",
				Assert:      func() { Expect(paramSet.Native.Directories).To(BeEquivalentTo([]string{"alpha", "beta", "delta"})) },
			}),

			Entry(nil, TcEntry{
				Message: "uint16 type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindUint16(
						adapters.NewFlagInfo("count16", "c", uint16(0)),
						&paramSet.Native.Count16,
					)
				},
				CommandLine: "--count16=333",
				Assert:      func() { Expect(paramSet.Native.Count16).To(Equal(uint16(333))) },
			}),

			Entry(nil, TcEntry{
				Message: "uint16 type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindUint16(
						adapters.NewFlagInfo("count16", "", uint16(0)),
						&paramSet.Native.Count16,
					)
				},
				CommandLine: "--count16=333",
				Assert:      func() { Expect(paramSet.Native.Count16).To(Equal(uint16(333))) },
			}),

			Entry(nil, TcEntry{
				Message: "uint32 type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindUint32(
						adapters.NewFlagInfo("count32", "c", uint32(0)),
						&paramSet.Native.Count32,
					)
				},
				CommandLine: "--count32=3333",
				Assert:      func() { Expect(paramSet.Native.Count32).To(Equal(uint32(3333))) },
			}),

			Entry(nil, TcEntry{
				Message: "uint32 type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindUint32(
						adapters.NewFlagInfo("count32", "", uint32(0)),
						&paramSet.Native.Count32,
					)
				},
				CommandLine: "--count32=3333",
				Assert:      func() { Expect(paramSet.Native.Count32).To(Equal(uint32(3333))) },
			}),

			Entry(nil, TcEntry{
				Message: "uint64 type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindUint64(
						adapters.NewFlagInfo("count64", "c", uint64(0)),
						&paramSet.Native.Count64,
					)
				},
				CommandLine: "--count64=33333",
				Assert:      func() { Expect(paramSet.Native.Count64).To(Equal(uint64(33333))) },
			}),

			Entry(nil, TcEntry{
				Message: "uint64 type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindUint64(
						adapters.NewFlagInfo("count64", "", uint64(0)),
						&paramSet.Native.Count64,
					)
				},
				CommandLine: "--count64=33333",
				Assert:      func() { Expect(paramSet.Native.Count64).To(Equal(uint64(33333))) },
			}),

			Entry(nil, TcEntry{
				Message: "uint8 type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindUint8(
						adapters.NewFlagInfo("count8", "c", uint8(0)),
						&paramSet.Native.Count8,
					)
				},
				CommandLine: "--count8=33",
				Assert:      func() { Expect(paramSet.Native.Count8).To(Equal(uint8(33))) },
			}),

			Entry(nil, TcEntry{
				Message: "uint8 type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindUint8(
						adapters.NewFlagInfo("count8", "", uint8(0)),
						&paramSet.Native.Count8,
					)
				},
				CommandLine: "--count8=33",
				Assert:      func() { Expect(paramSet.Native.Count8).To(Equal(uint8(33))) },
			}),

			Entry(nil, TcEntry{
				Message: "uint type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindUint(
						adapters.NewFlagInfo("count", "c", uint(0)),
						&paramSet.Native.Count,
					)
				},
				CommandLine: "--count=99999",
				Assert:      func() { Expect(paramSet.Native.Count).To(Equal(uint(99999))) },
			}),

			Entry(nil, TcEntry{
				Message: "[]uint slice type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindUintSlice(
						adapters.NewFlagInfo("counts", "P", []uint{}),
						&paramSet.Native.Counts,
					)
				},
				CommandLine: "--counts=2,4,6,8",
				Assert:      func() { Expect(paramSet.Native.Counts).To(BeEquivalentTo([]uint{2, 4, 6, 8})) },
			}),

			Entry(nil, TcEntry{
				Message: "uint type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindUint(
						adapters.NewFlagInfo("count", "", uint(0)),
						&paramSet.Native.Count,
					)
				},
				CommandLine: "--count=99999",
				Assert:      func() { Expect(paramSet.Native.Count).To(Equal(uint(99999))) },
			}),

			Entry(nil, TcEntry{
				Message: "[]uint slice type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindUintSlice(
						adapters.NewFlagInfo("counts", "", []uint{}),
						&paramSet.Native.Counts,
					)
				},
				CommandLine: "--counts=2,4,6,8",
				Assert:      func() { Expect(paramSet.Native.Counts).To(BeEquivalentTo([]uint{2, 4, 6, 8})) },
			}),

			// <---- auto generated(Build-PsTestEntry/gen-ps-t)
		)
	})
})
