package adapters_test

import (
	"fmt"
	"net"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	"github.com/snivilised/cobrass/src/adapters"
	"github.com/snivilised/cobrass/src/testhelpers"
)

type OutputFormatEnum int

const (
	XmlFormatEn OutputFormatEnum = iota + 1
	JsonFormatEn
	TextFormatEn
	ScribbleFormatEn
)

type WidgetParameterSet struct {
	Directory string
	Format    OutputFormatEnum
	Concise   bool
	Pattern   string
	//
	Offset   int
	Offset8  int8
	Offset16 int16
	Offset32 int32
	Offset64 int64
	//
	Count   uint
	Count8  uint8
	Count16 uint16
	Count32 uint32
	Count64 uint64
	//
	Threshold float64
	Gradient  float32
	Latency   time.Duration
	IpAddress net.IP
	IpMask    net.IPMask
	//
	Categories   []string
	Dimensions   []int
	Points       []uint
	Switches     []bool
	Scales       []float64
	Temperatures []float32
	Hosts        []net.IP
}

type TcEntry struct {
	Message     string
	Binder      func()
	CommandLine string
	Assert      func()
}

var _ = Describe("ParamSet", func() {

	When("Binding a flag", func() {
		var rootCommand *cobra.Command
		var widgetCommand *cobra.Command
		var paramSet *adapters.ParamSet[WidgetParameterSet]

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
					GinkgoWriter.Printf("===> 🍓 EXECUTE\n")
					return nil
				},
				PostRun: func(command *cobra.Command, args []string) {
					GinkgoWriter.Printf("**** 🍒 POST-RUN\n")
				},
			}
			rootCommand.AddCommand(widgetCommand)

			paramSet = adapters.NewParamSet[WidgetParameterSet](widgetCommand)
		})

		DescribeTable("binder",
			func(entry TcEntry) {
				entry.Binder()

				GinkgoWriter.Printf("🍧🍧🍧 ABOUT TO RUN ...\n")
				testhelpers.ExecuteCommand(
					rootCommand, "widget", "/usr/fuse/home/music", entry.CommandLine,
				)
				GinkgoWriter.Printf("🍧🍧🍧 AFTER RUN\n")

				entry.Assert()
			},

			func(entry TcEntry) string {
				return fmt.Sprintf("🧪 --> 🍒 given: flag is '%v'", entry.Message)
			},

			Entry(nil, TcEntry{
				Message: "string type",
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
				Message: "int type",
				Binder: func() {
					paramSet.BindInt(
						adapters.NewFlagInfo("offset", "o", -1),
						&paramSet.Native.Offset,
					)
				},
				CommandLine: "--offset=-99",
				Assert:      func() { Expect(paramSet.Native.Offset).To(Equal(-99)) },
			}),

			// ints ...

			Entry(nil, TcEntry{
				Message: "int8 type",
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
				Message: "int16 type",
				Binder: func() {
					paramSet.BindInt16(
						adapters.NewFlagInfo("offset16", "o", int16(-1)),
						&paramSet.Native.Offset16,
					)
				},
				CommandLine: "--offset16=-99",
				Assert:      func() { Expect(paramSet.Native.Offset16).To(Equal(int16(-99))) },
			}),

			Entry(nil, TcEntry{
				Message: "int32 type",
				Binder: func() {
					paramSet.BindInt32(
						adapters.NewFlagInfo("offset32", "o", int32(-1)),
						&paramSet.Native.Offset32,
					)
				},
				CommandLine: "--offset32=-99",
				Assert:      func() { Expect(paramSet.Native.Offset32).To(Equal(int32(-99))) },
			}),

			Entry(nil, TcEntry{
				Message: "int64 type",
				Binder: func() {
					paramSet.BindInt64(
						adapters.NewFlagInfo("offset64", "o", int64(-1)),
						&paramSet.Native.Offset64,
					)
				},
				CommandLine: "--offset64=-99",
				Assert:      func() { Expect(paramSet.Native.Offset64).To(Equal(int64(-99))) },
			}),

			Entry(nil, TcEntry{
				Message: "uint type",
				Binder: func() {
					paramSet.BindUint(
						adapters.NewFlagInfo("count", "c", uint(0)),
						&paramSet.Native.Count,
					)
				},
				CommandLine: "--count=33",
				Assert:      func() { Expect(paramSet.Native.Count).To(Equal(uint(33))) },
			}),

			Entry(nil, TcEntry{
				Message: "uint8 type",
				Binder: func() {
					paramSet.BindUint8(
						adapters.NewFlagInfo("count8", "c", uint8(1)),
						&paramSet.Native.Count8,
					)
				},
				CommandLine: "--count8=99",
				Assert:      func() { Expect(paramSet.Native.Count8).To(Equal(uint8(99))) },
			}),

			Entry(nil, TcEntry{
				Message: "uint16 type",
				Binder: func() {
					paramSet.BindUint16(
						adapters.NewFlagInfo("count16", "c", uint16(1)),
						&paramSet.Native.Count16,
					)
				},
				CommandLine: "--count16=99",
				Assert:      func() { Expect(paramSet.Native.Count16).To(Equal(uint16(99))) },
			}),

			Entry(nil, TcEntry{
				Message: "uint32 type",
				Binder: func() {
					paramSet.BindUint32(
						adapters.NewFlagInfo("count32", "c", uint32(1)),
						&paramSet.Native.Count32,
					)
				},
				CommandLine: "--count32=99",
				Assert:      func() { Expect(paramSet.Native.Count32).To(Equal(uint32(99))) },
			}),

			Entry(nil, TcEntry{
				Message: "uint64 type",
				Binder: func() {
					paramSet.BindUint64(
						adapters.NewFlagInfo("count64", "c", uint64(1)),
						&paramSet.Native.Count64,
					)
				},
				CommandLine: "--count64=99",
				Assert:      func() { Expect(paramSet.Native.Count64).To(Equal(uint64(99))) },
			}),

			Entry(nil, TcEntry{
				Message: "float64 type",
				Binder: func() {
					paramSet.BindFloat64(
						adapters.NewFlagInfo("threshold", "t", 999.9),
						&paramSet.Native.Threshold,
					)
				},
				CommandLine: "--threshold=99.1234",
				Assert:      func() { Expect(paramSet.Native.Threshold).To(Equal(99.1234)) },
			}),

			Entry(nil, TcEntry{
				Message: "float32 type",
				Binder: func() {
					paramSet.BindFloat32(
						adapters.NewFlagInfo("gradient", "g", float32(0.5105)),
						&paramSet.Native.Gradient,
					)
				},
				CommandLine: "--gradient=0.12345",
				Assert:      func() { Expect(paramSet.Native.Gradient).To(Equal(float32(0.12345))) },
			}),

			Entry(nil, TcEntry{
				Message: "bool type flag is present",
				Binder: func() {
					paramSet.BindBool(
						adapters.NewFlagInfo("concise", "c", false),
						&paramSet.Native.Concise,
					)
				},
				CommandLine: "--concise",
				Assert:      func() { Expect(paramSet.Native.Concise).To(BeTrue()) },
			}),

			Entry(nil, TcEntry{
				Message: "bool type flag is NOT present",
				Binder: func() {
					paramSet.BindBool(
						adapters.NewFlagInfo("concise", "c", false),
						&paramSet.Native.Concise,
					)
				},
				CommandLine: "",
				Assert:      func() { Expect(paramSet.Native.Concise).To(BeFalse()) },
			}),

			Entry(nil, TcEntry{
				Message: "duration type",
				Binder: func() {
					duration, _ := time.ParseDuration("0.5s")
					paramSet.BindDuration(
						adapters.NewFlagInfo("latency", "l", duration),
						&paramSet.Native.Latency,
					)
				},
				CommandLine: "--latency=300ms",
				Assert: func() {
					expected, _ := time.ParseDuration("300ms")
					Expect(paramSet.Native.Latency).To(BeEquivalentTo(expected))
				},
			}),

			Entry(nil, TcEntry{
				Message: "net.IP type",
				Binder: func() {
					address := net.IPv4(172, 16, 0, 0)
					paramSet.BindIp(
						adapters.NewFlagInfo("ip", "i", address),
						&paramSet.Native.IpAddress,
					)
				},
				CommandLine: "--ip=192.168.0.0",
				Assert: func() {
					expected := net.IPv4(192, 168, 0, 0)
					Expect(paramSet.Native.IpAddress).To(BeEquivalentTo(expected))
				},
			}),

			Entry(nil, TcEntry{
				Message: "net.IPMask type",
				Binder: func() {
					mask := net.IPv4Mask(0, 0, 0, 0)
					paramSet.BindIpMask(
						adapters.NewFlagInfo("ipmask", "m", mask),
						&paramSet.Native.IpMask,
					)
				},
				CommandLine: "--ipmask=255.255.255.0",
				Assert: func() {
					expected := net.IPv4Mask(255, 255, 255, 0)
					Expect(paramSet.Native.IpMask).To(BeEquivalentTo(expected))
				},
			}),

			Entry(nil, TcEntry{
				Message: "string slice",
				Binder: func() {
					paramSet.BindStringSlice(
						adapters.NewFlagInfo("categories", "c", []string{""}),
						&paramSet.Native.Categories,
					)
				},
				CommandLine: "--categories=a,b,c",
				Assert: func() {
					expected := []string{"a", "b", "c"}
					Expect(paramSet.Native.Categories).To(BeEquivalentTo(expected))
				},
			}),

			Entry(nil, TcEntry{
				Message: "int slice",
				Binder: func() {
					paramSet.BindIntSlice(
						adapters.NewFlagInfo("dimensions", "d", []int{}),
						&paramSet.Native.Dimensions,
					)
				},
				CommandLine: "--dimensions=1,2,3,4",
				Assert: func() {
					expected := []int{1, 2, 3, 4}
					Expect(paramSet.Native.Dimensions).To(BeEquivalentTo(expected))
				},
			}),

			Entry(nil, TcEntry{
				Message: "uint slice",
				Binder: func() {
					paramSet.BindUintSlice(
						adapters.NewFlagInfo("points", "p", []uint{}),
						&paramSet.Native.Points,
					)
				},
				CommandLine: "--points=0,15,30,45",
				Assert: func() {
					expected := []uint{0, 15, 30, 45}
					Expect(paramSet.Native.Points).To(BeEquivalentTo(expected))
				},
			}),

			Entry(nil, TcEntry{
				Message: "bool slice",
				Binder: func() {
					paramSet.BindBoolSlice(
						adapters.NewFlagInfo("switches", "s", []bool{}),
						&paramSet.Native.Switches,
					)
				},
				CommandLine: "--switches=true,false,true,false",
				Assert: func() {
					expected := []bool{true, false, true, false}
					Expect(paramSet.Native.Switches).To(BeEquivalentTo(expected))
				},
			}),

			Entry(nil, TcEntry{
				Message: "float64 slice",
				Binder: func() {
					paramSet.BindFloat64Slice(
						adapters.NewFlagInfo("scales", "s", []float64{}),
						&paramSet.Native.Scales,
					)
				},
				CommandLine: "--scales=1.2,2.3,3.4,4.5",
				Assert: func() {
					expected := []float64{1.2, 2.3, 3.4, 4.5}
					Expect(paramSet.Native.Scales).To(BeEquivalentTo(expected))
				},
			}),

			Entry(nil, TcEntry{
				Message: "float32 slice",
				Binder: func() {
					paramSet.BindFloat32Slice(
						adapters.NewFlagInfo("temperatures", "t", []float32{}),
						&paramSet.Native.Temperatures,
					)
				},
				CommandLine: "--temperatures=99.2,99.3,99.4,99.5",
				Assert: func() {
					expected := []float32{99.2, 99.3, 99.4, 99.5}
					Expect(paramSet.Native.Temperatures).To(BeEquivalentTo(expected))
				},
			}),

			Entry(nil, TcEntry{
				Message: "ip slice",
				Binder: func() {
					paramSet.BindIpSlice(
						adapters.NewFlagInfo("hosts", "k", []net.IP{}),
						&paramSet.Native.Hosts,
					)
				},
				CommandLine: "--hosts=192.168.0.0,172.16.0.0",
				Assert: func() {
					expected := []net.IP{net.IPv4(192, 168, 0, 0), net.IPv4(172, 16, 0, 0)}
					Expect(paramSet.Native.Hosts).To(BeEquivalentTo(expected))
				},
			}),
		)

		Context("given: enum type", func() {
			var OutputFormatEnumInfo *adapters.EnumInfo[OutputFormatEnum]

			BeforeEach(func() {
				OutputFormatEnumInfo = adapters.NewEnumInfo(adapters.AcceptableEnumValues[OutputFormatEnum]{
					XmlFormatEn:      []string{"xml", "x"},
					JsonFormatEn:     []string{"json", "j"},
					TextFormatEn:     []string{"text", "tx"},
					ScribbleFormatEn: []string{"scribble", "scribbler", "scr"},
				})
			})
			It("🧪 should: create enum info", func() {
				Expect(OutputFormatEnumInfo.En("x")).To(Equal(XmlFormatEn))
				Expect(OutputFormatEnumInfo.En("xml")).To(Equal(XmlFormatEn))

				Expect(OutputFormatEnumInfo.En("j")).To(Equal(JsonFormatEn))
				Expect(OutputFormatEnumInfo.En("json")).To(Equal(JsonFormatEn))
			})

			Context("given: int based enum type", func() {
				It("🧪 should: populate member of native parameter set", func() {

					paramSet.BindEnum(
						&adapters.FlagInfo{Name: "format", Short: "f", Default: "text", Usage: "format"},
						&OutputFormatEnumInfo.Source,
					)

					testhelpers.ExecuteCommand(
						rootCommand, "widget", "/usr/fuse/home/music", "--format=xml",
					)

					// This is the line of code that ideally needs to be executed automatically
					// somehow, after the command line has been parsed, possibly command.PreRun
					// rebind/rebinder/bindnative/ for enum fields? Without this automation, the
					// client needs to do this manually, but a single line of code is hardly
					// burdensome, given that any automated scheme would probably not be
					// as efficient, probably requiring more than a single line of code anyway.
					// The documentation of BindEnum does in fact instruct the reader to do so.
					//
					paramSet.Native.Format = OutputFormatEnumInfo.Value()
					Expect(paramSet.Native.Format).To(Equal(XmlFormatEn))
				})
			})
		})
	})
})
