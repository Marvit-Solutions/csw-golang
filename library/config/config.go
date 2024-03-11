package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config defines the methods for accessing configuration values.
type Config interface {
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetStrings(key string) []string
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]interface{}
	Init(string)
}

// viperConfig implements the Config interface using viper.
type viperConfig struct{}

// Init initializes the configuration with the specified prefix.
func (v *viperConfig) Init(prefix string) {
	viper.SetEnvPrefix(`go-clean`)
	viper.AutomaticEnv()

	osEnv := os.Getenv("OS_ENV")

	env := "env"
	if osEnv != "" {
		env = osEnv
	}

	if prefix != "" {
		env = prefix + "." + env
	}

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigFile(env + `.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

// GetString returns the string value of the specified key.
func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

// GetInt returns the integer value of the specified key.
func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

// GetBool returns the boolean value of the specified key.
func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetStringSlice returns the string slice value of the specified key.
func (v *viperConfig) GetStringSlice(key string) (c []string) {
	c = viper.GetStringSlice(key)
	return
}

// GetStrings returns the string slice value of the specified key, split by comma.
func (v *viperConfig) GetStrings(key string) (c []string) {
	val := viper.GetString(key)
	c = strings.Split(val, ",")
	return
}

// GetStringMap returns the string map value of the specified key.
func (v *viperConfig) GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

// NewConfig creates a new instance of the Config interface using viperConfig.
func NewConfig() Config {
	v := &viperConfig{}
	v.Init("")
	return v
}
