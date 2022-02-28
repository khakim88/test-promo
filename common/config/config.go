package config

import (
	"os"
	"strconv"
	"time"
)

// GetString returns config value for `key` as string. If no value is
// found, returns `_default``
func GetString(key, _default string) string {
	val := os.Getenv(key)
	if val == "" {
		return _default
	}
	return val
}

// GetBool returns config value for `key` as bool. If no value is found,
// returns `_default``
func GetBool(key string, _default bool) bool {
	val, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return _default
	}
	return val
}

// GetFloat returns config value for `key` as float64. If no value is
// found, returns `_default`
func GetFloat(key string, _default float64) float64 {
	val, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		return _default
	}
	return val
}

// GetInt returns config value for `key` as int. If no value is found,
// returns _default
func GetInt(key string, _default int) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return _default
	}
	return val
}

// GetInt64 returns config value for `key` as int64. If no value is found,
// returns _default
func GetInt64(key string, _default int64) int64 {
	val, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil {
		return _default
	}
	return val
}

// GetUInt returns config value for `key` as uint64. If no value is found,
// returns _default
func GetUInt(key string, _default uint64) uint64 {
	val, err := strconv.ParseUint(os.Getenv(key), 10, 64)
	if err != nil {
		return _default
	}
	return val
}

// GetDuration returns config value for `key` as duration. If no value is found,
// returns _default
func GetDuration(key string, _default time.Duration) time.Duration {
	val, err := time.ParseDuration(os.Getenv(key))
	if err != nil {
		return _default
	}
	return val
}

var env string

func IsProduction() bool {
	env = getEnvType()
	return env == "production"
}

func getEnvType() string {
	if env == "" {
		env = GetString("ENV", "development")
	}
	return env
}
