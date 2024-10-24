package assistant_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
	. "github.com/onsi/gomega"    //nolint:revive // gomega ok
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/internal/lab"
	"github.com/spf13/cobra"
)

// the auto version of param-set_test.go

var _ = Describe("ParamSet (auto)", func() {

	When("Binding a flag (auto)", func() {
		var rootCommand *cobra.Command
		var widgetCommand *cobra.Command
		var paramSet *assistant.ParamSet[WidgetParameterSet]
		var outputFormatEnumInfo *assistant.EnumInfo[OutputFormatEnum]
		var outputFormatEnum assistant.EnumValue[OutputFormatEnum]

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

				PreRun: func(_ *cobra.Command, _ []string) {
					GinkgoWriter.Printf("**** 🍉 PRE-RUN\n")
				},
				RunE: func(_ *cobra.Command, args []string) error {
					GinkgoWriter.Printf("===> 🍓 EXECUTE (Directory: '%v')\n", args[0])

					paramSet.Native.Directory = args[0]
					return nil
				},
				PostRun: func(_ *cobra.Command, _ []string) {
					GinkgoWriter.Printf("**** 🥥 POST-RUN\n")
				},
			}
			rootCommand.AddCommand(widgetCommand)

			paramSet = assistant.NewParamSet[WidgetParameterSet](widgetCommand)

			outputFormatEnumInfo = assistant.NewEnumInfo(AcceptableOutputFormats)
			outputFormatEnum = outputFormatEnumInfo.NewValue()
		})

		DescribeTable("binder",
			func(entry TcEntry) {
				entry.Binder()

				_, _ = lab.ExecuteCommand(
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
						assistant.NewFlagInfo("concise", "c", false),
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
						assistant.NewFlagInfo("switches", "S", []bool{}),
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
						assistant.NewFlagInfo("concise", "", false),
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
						assistant.NewFlagInfo("switches", "", []bool{}),
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
						assistant.NewFlagInfo("latency", "l", duration("0ms")),
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
						assistant.NewFlagInfo("latencies", "L", []time.Duration{}),
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
						assistant.NewFlagInfo("latency", "", duration("0ms")),
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
						assistant.NewFlagInfo("latencies", "", []time.Duration{}),
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
						assistant.NewFlagInfo("format", "f", "xml"),
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
						assistant.NewFlagInfo("format", "", "xml"),
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
						assistant.NewFlagInfo("gradientf32", "t", float32(0)),
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
						assistant.NewFlagInfo("gradientsf32", "G", []float32{}),
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
						assistant.NewFlagInfo("gradientf32", "", float32(0)),
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
						assistant.NewFlagInfo("gradientsf32", "", []float32{}),
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
						assistant.NewFlagInfo("gradientf64", "t", float64(0)),
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
						assistant.NewFlagInfo("gradientsf64", "G", []float64{}),
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
						assistant.NewFlagInfo("gradientf64", "", float64(0)),
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
						assistant.NewFlagInfo("gradientsf64", "", []float64{}),
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
						assistant.NewFlagInfo("offset", "o", -1),
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
						assistant.NewFlagInfo("offsets", "D", []int{}),
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
						assistant.NewFlagInfo("offset", "", -1),
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
						assistant.NewFlagInfo("offsets", "", []int{}),
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
						assistant.NewFlagInfo("offset16", "o", int16(-1)),
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
						assistant.NewFlagInfo("offset16", "", int16(-1)),
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
						assistant.NewFlagInfo("offset32", "o", int32(-1)),
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
						assistant.NewFlagInfo("offsets32", "O", []int32{}),
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
						assistant.NewFlagInfo("offset32", "", int32(-1)),
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
						assistant.NewFlagInfo("offsets32", "", []int32{}),
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
						assistant.NewFlagInfo("offset64", "o", int64(-1)),
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
						assistant.NewFlagInfo("offsets64", "O", []int64{}),
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
						assistant.NewFlagInfo("offset64", "", int64(-1)),
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
						assistant.NewFlagInfo("offsets64", "", []int64{}),
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
						assistant.NewFlagInfo("offset8", "o", int8(-1)),
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
						assistant.NewFlagInfo("offset8", "", int8(-1)),
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
						assistant.NewFlagInfo("ipmask", "m", ipmask("default")),
						&paramSet.Native.IPMask,
					)
				},
				CommandLine: "--ipmask=255.255.255.0",
				Assert:      func() { Expect(paramSet.Native.IPMask).To(BeEquivalentTo(ipmask("orion.net"))) },
			}),

			Entry(nil, TcEntry{
				Message: "net.IPMask type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindIPMask(
						assistant.NewFlagInfo("ipmask", "", ipmask("default")),
						&paramSet.Native.IPMask,
					)
				},
				CommandLine: "--ipmask=255.255.255.0",
				Assert:      func() { Expect(paramSet.Native.IPMask).To(BeEquivalentTo(ipmask("orion.net"))) },
			}),

			Entry(nil, TcEntry{
				Message: "net.IPNet type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindIPNet(
						assistant.NewFlagInfo("ipaddress", "i", ipnet("default")),
						&paramSet.Native.IPAddress,
					)
				},
				CommandLine: "--ipaddress=172.16.0.0",
				Assert:      func() { Expect(paramSet.Native.IPAddress).ToNot(BeNil()) },
			}),

			Entry(nil, TcEntry{
				Message: "net.IPNet type, (without-short) (auto)",
				Binder: func() {
					paramSet.BindIPNet(
						assistant.NewFlagInfo("ipaddress", "", ipnet("default")),
						&paramSet.Native.IPAddress,
					)
				},
				CommandLine: "--ipaddress=172.16.0.0",
				Assert:      func() { Expect(paramSet.Native.IPAddress).ToNot(BeNil()) },
			}),

			Entry(nil, TcEntry{
				Message: "string type, (with-short) (auto)",
				Binder: func() {
					paramSet.BindString(
						assistant.NewFlagInfo("pattern", "p", "default-pattern"),
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
						assistant.NewFlagInfo("directories", "C", []string{}),
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
						assistant.NewFlagInfo("pattern", "", "default-pattern"),
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
						assistant.NewFlagInfo("directories", "", []string{}),
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
						assistant.NewFlagInfo("count16", "c", uint16(0)),
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
						assistant.NewFlagInfo("count16", "", uint16(0)),
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
						assistant.NewFlagInfo("count32", "c", uint32(0)),
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
						assistant.NewFlagInfo("count32", "", uint32(0)),
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
						assistant.NewFlagInfo("count64", "c", uint64(0)),
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
						assistant.NewFlagInfo("count64", "", uint64(0)),
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
						assistant.NewFlagInfo("count8", "c", uint8(0)),
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
						assistant.NewFlagInfo("count8", "", uint8(0)),
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
						assistant.NewFlagInfo("count", "c", uint(0)),
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
						assistant.NewFlagInfo("counts", "P", []uint{}),
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
						assistant.NewFlagInfo("count", "", uint(0)),
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
						assistant.NewFlagInfo("counts", "", []uint{}),
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
