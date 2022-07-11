package adapters_test

import (
	"fmt"
	"net"
	"time"

	"github.com/snivilised/cobrass/src/adapters"
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

type InvalidParameterSet string

func duration(d string) time.Duration {
	result, _ := time.ParseDuration(d)
	return result
}

func ipmask(v string) net.IPMask {
	var result net.IPMask
	switch v {
	case "default":
		result = net.IPMask([]byte{0, 0, 0, 0})

	case "orion.net":
		result = net.IPMask([]byte{255, 255, 255, 0})

	default:
		panic(fmt.Errorf("no ipmask defined for: '%v'", v))
	}

	return result
}

func ipnet(v string) net.IPNet {
	var result net.IPNet
	switch v {
	case "default":
		result = net.IPNet{IP: net.IPv4(0, 0, 0, 0), Mask: ipmask(v)}

	case "orion.net":
		result = net.IPNet{IP: net.IPv4(172, 16, 0, 0), Mask: ipmask(v)}

	default:
		panic(fmt.Errorf("no ipnet defined for: '%v'", v))
	}

	return result
}

var AcceptableOutputFormats = adapters.AcceptableEnumValues[OutputFormatEnum]{
	XmlFormatEn:      []string{"xml", "x"},
	JsonFormatEn:     []string{"json", "j"},
	TextFormatEn:     []string{"text", "tx"},
	ScribbleFormatEn: []string{"scribble", "scribbler", "scr"},
}
