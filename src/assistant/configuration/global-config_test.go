package configuration_test

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"

	"github.com/snivilised/cobrass/src/assistant/configuration"
)

const (
	awardsField = "awards"
	themesField = "themes"
	relative    = "../../../test/configuration"
)

type configTE struct {
	message  string
	field    string
	expected any
	actual   func(entry *configTE) any
	assert   func(entry *configTE, actual any)
}

type fakeFlag struct {
	hasChanged bool
	name       string
	val        string
	typ        string
}

func (f *fakeFlag) HasChanged() bool {
	return f.hasChanged
}

func (f *fakeFlag) Name() string {
	return f.name
}

func (f *fakeFlag) ValueString() string {
	return f.val
}

func (f *fakeFlag) ValueType() string {
	return f.typ
}

type foo struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type themes struct {
	First  []string `json:"first"`
	Second []string `json:"second"`
	Third  []string `json:"third"`
}

func reason(field string, expected, actual any) string {
	return fmt.Sprintf("🔥 expected field '%v' to be '%v', but was '%v'\n",
		field, expected, actual,
	)
}

var _ = Describe("GlobalConfig", Ordered, func() {
	var config configuration.ViperConfig

	BeforeEach(func() {
		viper.Reset()
		config = &configuration.GlobalViperConfig{}

		config.SetConfigName("cobrass")
		config.SetConfigType("yml")

		if _, err := os.Lstat(relative); err != nil {
			Fail("🔥 can't find config path")
		}
		config.AddConfigPath(relative)
		if err := config.ReadInConfig(); err != nil {
			Fail(fmt.Sprintf("🔥 can't read config (err: '%v')", err))
		}
	})
	Context("AutomaticEnv", func() {
		It("🧪 should: not fail", func() {
			config.AutomaticEnv()
		})
	})

	Context("BindFlagValue", func() {
		It("🧪 should: access field", func() {
			field := "foo-bar"
			flag := &fakeFlag{}
			err := config.BindFlagValue(field, flag)
			Expect(err).Error().To(BeNil())
		})
	})

	// BindFlagValue
	// BindFlagValues
	// BindPFlag

	Context("GetStringMap", func() {
		It("🧪 should: access string map field", func() {
			field := awardsField
			actual := config.GetStringMap(field)
			Expect(len(actual)).To(Equal(3), reason(field, 3, actual))
			Expect(actual).To(ContainElement("gold"), reason(field, "gold", actual["first"]))
			Expect(actual).To(ContainElement("silver"), reason(field, "silver", actual["second"]))
			Expect(actual).To(ContainElement("bronze"), reason(field, "bronze", actual["third"]))
		})
	})

	Context("GetStringMapString", func() {
		It("🧪 should: access string map field", func() {
			field := awardsField
			actual := config.GetStringMapString(field)
			Expect(len(actual)).To(Equal(3), reason(field, 3, actual))
			Expect(actual).To(ContainElement("gold"), reason(field, "gold", actual["first"]))
			Expect(actual).To(ContainElement("silver"), reason(field, "silver", actual["second"]))
			Expect(actual).To(ContainElement("bronze"), reason(field, "bronze", actual["third"]))
		})
	})

	Context("GetStringMapStringSlice", func() {
		It("🧪 should: access string map field", func() {
			field := themesField
			actual := config.GetStringMapStringSlice(field)
			Expect(len(actual)).To(Equal(3), reason(field, 3, actual))

			first := actual["first"]
			Expect(first).To(ContainElement("gold"), reason(field, "gold", actual["first"]))

			second := actual["second"]
			Expect(second).To(ContainElement("silver"), reason(field, "silver", actual["second"]))

			third := actual["third"]
			Expect(third).To(ContainElement("bronze"), reason(field, "bronze", actual["third"]))
		})
	})

	Context("Sub", func() {
		It("🧪 should: get sub tree", func() {
			field := themesField
			instance := config.Sub(field)
			{
				field := "first"
				first := instance.GetStringSlice(field)
				Expect(first).To(ContainElement("gold"), reason(field, "", ""))
			}
		})
	})

	Context("UnmarshalKey", func() {
		It("🧪 should: unmarshal a specified key", func() {
			field := themesField
			positions := themes{}
			err := config.UnmarshalKey(field, &positions)
			Expect(err).Error().To(BeNil())
			Expect(positions.First).To(ContainElement("gold"), reason(field, "", ""))
		})
	})

	DescribeTable("config fields",
		func(entry *configTE) {
			if entry.assert == nil {
				actual := entry.actual(entry)
				Expect(actual).To(Equal(entry.expected),
					reason(entry.field, entry.expected, actual),
				)
			} else {
				actual := entry.actual(entry)
				entry.assert(entry, actual)
			}
		},
		func(entry *configTE) string {
			return fmt.Sprintf("🧪 ===> given: '%v', should access field: '%v'",
				entry.message, entry.field,
			)
		},

		Entry(nil, &configTE{
			message:  "anon getter",
			field:    "the-answer",
			expected: 42,
			actual: func(e *configTE) any {
				return config.Get(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "BindFlagValue",
			field:    "foo-bar",
			expected: true,
			actual: func(e *configTE) any {
				flag := &fakeFlag{}
				return config.BindFlagValue(e.field, flag)
			},
			assert: func(e *configTE, actual any) {
				Expect(actual).Error().To(BeNil(), reason(e.field, e.expected, actual))
			},
		}),

		Entry(nil, &configTE{
			message:  "ConfigFileUsed",
			field:    "ConfigFileUsed",
			expected: "cobrass.yml",
			actual: func(e *configTE) any {
				full := config.ConfigFileUsed()
				return filepath.Base(full)
			},
		}),

		Entry(nil, &configTE{
			message:  "GetBool",
			field:    "is-magical",
			expected: true,
			actual: func(e *configTE) any {
				return config.GetBool(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "GetDuration",
			field:    "delay",
			expected: time.Second * 10,
			actual: func(e *configTE) any {
				return config.GetDuration(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "GetFloat64",
			field:    "portion-64",
			expected: float64(0.1234),
			actual: func(e *configTE) any {
				return config.GetFloat64(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "GetInt",
			field:    "int-no-of-buckets",
			expected: int(88),
			actual: func(e *configTE) any {
				return config.GetInt(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "GetInt32",
			field:    "counter-32",
			expected: int32(132),
			actual: func(e *configTE) any {
				return config.GetInt32(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "GetInt64",
			field:    "counter-64",
			expected: int64(164),
			actual: func(e *configTE) any {
				return config.GetInt64(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "GetIntSlice",
			field:    "int-slice",
			expected: []int{5, 1, 5, 0},
			actual: func(e *configTE) any {
				return config.GetIntSlice(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "GetUint",
			field:    "counter-uint",
			expected: uint(99),
			actual: func(e *configTE) any {
				return config.GetUint(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "GetUint16",
			field:    "counter-u16",
			expected: uint16(216),
			actual: func(e *configTE) any {
				return config.GetUint16(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "GetUint32",
			field:    "counter-u32",
			expected: uint32(232),
			actual: func(e *configTE) any {
				return config.GetUint32(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "GetUint64",
			field:    "counter-u64",
			expected: uint64(264),
			actual: func(e *configTE) any {
				return config.GetUint64(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "GetTime",
			field:    "the-omen",
			expected: time.Date(2006, 6, 6, 0, 0, 0, 0, time.UTC),
			actual: func(e *configTE) any {
				return config.GetTime(e.field)
			},
			assert: func(entry *configTE, actual any) {
				// how does the time work in config?
				//
			},
		}),

		Entry(nil, &configTE{
			message:  "GetSizeInBytes",
			field:    "the-answer",
			expected: uint64(4),
			actual: func(e *configTE) any {
				return config.GetSizeInBytes(e.field)
			},
			assert: func(entry *configTE, actual any) {
				// how does the GetSizeInBytes work in config?
				//
			},
		}),

		Entry(nil, &configTE{
			message:  "GetString",
			field:    "the-question",
			expected: "are you master of your domain?",
			actual: func(e *configTE) any {
				return config.GetString(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "GetStringMap",
			field:    "awards",
			expected: "are you master of your domain?",
			actual: func(e *configTE) any {
				return config.GetStringMap(e.field)
			},
			assert: func(entry *configTE, actual any) {
				// how does the GetSizeInBytes work in config?
				//
			},
		}),

		Entry(nil, &configTE{
			message:  "GetStringSlice",
			field:    "colours",
			expected: []string{"red", "green", "blue"},
			actual: func(e *configTE) any {
				return config.GetStringSlice(e.field)
			},
		}),

		Entry(nil, &configTE{
			message:  "InConfig",
			field:    "colours",
			expected: true,
			actual: func(e *configTE) any {
				return config.InConfig(e.field)
			},
		}),

		// SetTypeByDefaultValue

		Entry(nil, &configTE{
			message:  "Sub",
			field:    "themes",
			expected: true,
			actual: func(e *configTE) any {
				return config.Sub(e.field)
			},
			assert: func(entry *configTE, actual any) {
				//
			},
		}),

		// UnmarshalKey
		Entry(nil, &configTE{
			message:  "UnmarshalKey",
			field:    "themes",
			expected: true,
			actual: func(e *configTE) any {
				positions := themes{}
				return config.UnmarshalKey(e.field, &positions)
			},
			assert: func(entry *configTE, actual any) {
				//
			},
		}),
	)
})

var _ = Describe("SetConfigFile", func() {
	Context("given: config file explicitly set", func() {
		It("should: read the explicit config", func() {
			config := &configuration.GlobalViperConfig{}
			explicit := filepath.Join(relative, "alt-config.yml")
			config.SetConfigFile(explicit)
			if err := config.ReadInConfig(); err != nil {
				Fail(fmt.Sprintf("🔥 can't read config (err: '%v')", err))
			}
			Expect(config.GetString("cryptic-declaration")).To(Equal("the monarch shall be crowned"))
		})
	})
})

var _ = Describe("Unmarshal", func() {
	Context("given: a config", func() {
		It("should: unmarshal entire config", func() {
			viper.Reset()
			config := &configuration.GlobalViperConfig{}
			explicit := filepath.Join(relative, "foo-config.yml")
			config.SetConfigFile(explicit)
			if err := config.ReadInConfig(); err != nil {
				Fail(fmt.Sprintf("🔥 can't read config (err: '%v')", err))
			}
			bar := foo{}
			err := config.Unmarshal(&bar)
			Expect(err).Error().To(BeNil())
			Expect(bar.Name).To(Equal("quantico"))
			Expect(bar.Count).To(Equal(911))
		})
	})
})

var _ = Describe("UnmarshalExact", func() {
	Context("given: a config", func() {
		It("should: unmarshal entire config", func() {
			viper.Reset()
			config := &configuration.GlobalViperConfig{}
			explicit := filepath.Join(relative, "foo-config.yml")
			config.SetConfigFile(explicit)
			if err := config.ReadInConfig(); err != nil {
				Fail(fmt.Sprintf("🔥 can't read config (err: '%v')", err))
			}
			bar := foo{}
			err := config.UnmarshalExact(&bar)
			Expect(err).Error().To(BeNil())
			Expect(bar.Name).To(Equal("quantico"))
			Expect(bar.Count).To(Equal(911))
		})
	})
})
