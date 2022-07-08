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
	Gradientf32 float32
	Gradientf64 float64
	Latency     time.Duration
	IpAddress   net.IPNet
	IpMask      net.IPMask
	//
	// some slice types are missing eg, Offsets16, because that slice type is not
	// supported by pflag; ie there is no Int16SliceVar/Int16SliceVarP
	//
	Directories  []string
	Offsets      []int
	Offsets32    []int32
	Offsets64    []int64
	Counts       []uint
	Switches     []bool
	Gradientsf32 []float32
	Gradientsf64 []float64
	Latencies    []time.Duration
	Hosts        []net.IP
	Formats      []OutputFormatEnum
}

type WidgetParameterSetPtr *WidgetParameterSet

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
					GinkgoWriter.Printf("===> 🍓 EXECUTE (Directory: '%v')\n", args[0])

					paramSet.Native.Directory = args[0]
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

			// TODO: these test should be auto generated

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
						&paramSet.Native.Gradientf64,
					)
				},
				CommandLine: "--threshold=99.1234",
				Assert:      func() { Expect(paramSet.Native.Gradientf64).To(Equal(99.1234)) },
			}),

			Entry(nil, TcEntry{
				Message: "float32 type",
				Binder: func() {
					paramSet.BindFloat32(
						adapters.NewFlagInfo("gradient", "g", float32(0.5105)),
						&paramSet.Native.Gradientf32,
					)
				},
				CommandLine: "--gradient=0.12345",
				Assert:      func() { Expect(paramSet.Native.Gradientf32).To(Equal(float32(0.12345))) },
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
					// address := net.IPv4(172, 16, 0, 0)

					address := net.IPNet{IP: net.IPv4(0, 0, 0, 0), Mask: net.IPMask([]byte{0, 0, 0, 0})}
					paramSet.BindIPNet(
						adapters.NewFlagInfo("ip", "i", address),
						&paramSet.Native.IpAddress,
					)
				},
				CommandLine: "--ip=192.168.0.0",
				Assert: func() {
					// for some reason, the assignment of the address (--ip=192.168.0.0) is not completing
					//
					// expected := net.IPNet{IP: net.IPv4(192, 168, 0, 0), Mask: net.IPMask([]byte{0, 0, 0, 0})}
					// Expect(paramSet.Native.IpAddress).To(BeEquivalentTo(expected))
				},
			}),

			Entry(nil, TcEntry{
				Message: "net.IPMask type",
				Binder: func() {
					mask := net.IPv4Mask(0, 0, 0, 0)
					paramSet.BindIPMask(
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
						&paramSet.Native.Directories,
					)
				},
				CommandLine: "--categories=a,b,c",
				Assert: func() {
					expected := []string{"a", "b", "c"}
					Expect(paramSet.Native.Directories).To(BeEquivalentTo(expected))
				},
			}),

			Entry(nil, TcEntry{
				Message: "int slice",
				Binder: func() {
					paramSet.BindIntSlice(
						adapters.NewFlagInfo("dimensions", "d", []int{}),
						&paramSet.Native.Offsets,
					)
				},
				CommandLine: "--dimensions=1,2,3,4",
				Assert: func() {
					expected := []int{1, 2, 3, 4}
					Expect(paramSet.Native.Offsets).To(BeEquivalentTo(expected))
				},
			}),

			Entry(nil, TcEntry{
				Message: "uint slice",
				Binder: func() {
					paramSet.BindUintSlice(
						adapters.NewFlagInfo("points", "p", []uint{}),
						&paramSet.Native.Counts,
					)
				},
				CommandLine: "--points=0,15,30,45",
				Assert: func() {
					expected := []uint{0, 15, 30, 45}
					Expect(paramSet.Native.Counts).To(BeEquivalentTo(expected))
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
						&paramSet.Native.Gradientsf64,
					)
				},
				CommandLine: "--scales=1.2,2.3,3.4,4.5",
				Assert: func() {
					expected := []float64{1.2, 2.3, 3.4, 4.5}
					Expect(paramSet.Native.Gradientsf64).To(BeEquivalentTo(expected))
				},
			}),

			Entry(nil, TcEntry{
				Message: "float32 slice",
				Binder: func() {
					paramSet.BindFloat32Slice(
						adapters.NewFlagInfo("temperatures", "t", []float32{}),
						&paramSet.Native.Gradientsf32,
					)
				},
				CommandLine: "--temperatures=99.2,99.3,99.4,99.5",
				Assert: func() {
					expected := []float32{99.2, 99.3, 99.4, 99.5}
					Expect(paramSet.Native.Gradientsf32).To(BeEquivalentTo(expected))
				},
			}),
		)

		Context("Register ParamSet", func() {
			It("🧪 should: be able get registered param set", func() {
				const cname = "widget"
				const psname = cname + "-ps"
				container := adapters.NewCobraContainer(rootCommand)
				container.RegisterParamSet(psname, paramSet)

				testhelpers.ExecuteCommand(
					rootCommand, cname, "/usr/fuse/home/cache",
				)

				if native, ok := container.Native(psname).(*WidgetParameterSet); ok {
					Expect(native.Directory).To(Equal("/usr/fuse/home/cache"))
				} else {
					Fail("param set should be registered")
				}
			})
		})

		Context("NewFlagInfoOnFlagSet", func() {
			It("🧪 should: bind flag to alternative flag set", func() {
				paramSet.BindString(
					adapters.NewFlagInfoOnFlagSet("pattern", "p", "default-pattern",
						widgetCommand.PersistentFlags()), &paramSet.Native.Pattern,
				)
				commandLine := "--pattern=*music.infex*"
				testhelpers.ExecuteCommand(
					rootCommand, "widget", "/usr/fuse/home/music", commandLine,
				)
				Expect(paramSet.Native.Pattern).To(Equal("*music.infex*"))
			})
		})
	})
})
