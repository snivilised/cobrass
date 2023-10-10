package configuration

import (
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ViperConfig interface {
	AddConfigPath(in string)
	AutomaticEnv()
	BindFlagValue(key string, flag viper.FlagValue) error
	BindFlagValues(flags viper.FlagValueSet) error
	BindPFlag(key string, flag *pflag.Flag) error
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetIntSlice(key string) []int
	GetUint(key string) uint
	GetUint16(key string) uint16
	GetUint32(key string) uint32
	GetUint64(key string) uint64
	GetTime(key string) time.Time
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	InConfig(key string) bool
	ReadInConfig() error
	SetConfigFile(in string)
	SetConfigName(in string)
	SetConfigType(in string)
	SetTypeByDefaultValue(enable bool)
	Sub(key string) *viper.Viper
	Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error
	UnmarshalExact(rawVal interface{}, opts ...viper.DecoderConfigOption) error
	UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error
}

type GlobalViperConfig struct {
}

func (p *GlobalViperConfig) AddConfigPath(in string) {
	viper.AddConfigPath(in)
}

func (p *GlobalViperConfig) AutomaticEnv() {
	viper.AutomaticEnv()
}

func (p *GlobalViperConfig) BindFlagValue(key string, flag viper.FlagValue) error {
	return viper.BindFlagValue(key, flag)
}

func (p *GlobalViperConfig) BindFlagValues(flags viper.FlagValueSet) error {
	return viper.BindFlagValues(flags)
}

func (p *GlobalViperConfig) BindPFlag(key string, flag *pflag.Flag) error {
	return viper.BindPFlag(key, flag)
}

func (p *GlobalViperConfig) ConfigFileUsed() string {
	return viper.ConfigFileUsed()
}

func (p *GlobalViperConfig) Get(key string) interface{} {
	return viper.Get(key)
}

func (p *GlobalViperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (p *GlobalViperConfig) GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func (p *GlobalViperConfig) GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func (p *GlobalViperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (p *GlobalViperConfig) GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func (p *GlobalViperConfig) GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func (p *GlobalViperConfig) GetIntSlice(key string) []int {
	return viper.GetIntSlice(key)
}

func (p *GlobalViperConfig) GetUint(key string) uint {
	return viper.GetUint(key)
}

func (p *GlobalViperConfig) GetUint16(key string) uint16 {
	return viper.GetUint16(key)
}

func (p *GlobalViperConfig) GetUint32(key string) uint32 {
	return viper.GetUint32(key)
}

func (p *GlobalViperConfig) GetUint64(key string) uint64 {
	return viper.GetUint64(key)
}

func (p *GlobalViperConfig) GetTime(key string) time.Time {
	return viper.GetTime(key)
}

func (p *GlobalViperConfig) GetSizeInBytes(key string) uint {
	return viper.GetSizeInBytes(key)
}

func (p *GlobalViperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (p *GlobalViperConfig) GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func (p *GlobalViperConfig) GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func (p *GlobalViperConfig) GetStringMapStringSlice(key string) map[string][]string {
	return viper.GetStringMapStringSlice(key)
}

func (p *GlobalViperConfig) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func (p *GlobalViperConfig) InConfig(key string) bool {
	return viper.InConfig(key)
}

func (p *GlobalViperConfig) ReadInConfig() error {
	return viper.ReadInConfig()
}

func (p *GlobalViperConfig) SetConfigFile(in string) {
	viper.SetConfigFile(in)
}

func (p *GlobalViperConfig) SetConfigName(in string) {
	viper.SetConfigName(in)
}

func (p *GlobalViperConfig) SetConfigType(in string) {
	viper.SetConfigType(in)
}

func (p *GlobalViperConfig) SetTypeByDefaultValue(enable bool) {
	viper.SetTypeByDefaultValue(enable)
}

func (p *GlobalViperConfig) Sub(key string) *viper.Viper {
	return viper.Sub(key)
}

func (p *GlobalViperConfig) Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	return viper.Unmarshal(rawVal, opts...)
}

func (p *GlobalViperConfig) UnmarshalExact(rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	return viper.UnmarshalExact(rawVal, opts...)
}

func (p *GlobalViperConfig) UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	return viper.UnmarshalKey(key, rawVal, opts...)
}
