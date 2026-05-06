// Package env provides typed helpers for reading environment variables.
// Helpers panic on missing or unparseable values so that callers can
// centralize error recovery at startup.
//
// Each MustX helper returns the parsed value of the named environment
// variable. If the variable is unset and a default is provided, the
// default is returned. If the variable is unset and no default is
// provided, the helper panics with a wrapped [ErrNotExist]. If the
// variable is set but cannot be parsed, the helper panics with a
// wrapped [ErrParse]. Both sentinels are exported so callers can use
// [errors.Is] after a recover.
package env

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// ErrNotExist is the sentinel wrapped in a panic when a required env
// var is unset and no default was provided.
var ErrNotExist = errors.New("does not exist")

// ErrParse is the sentinel wrapped in a panic when an env var's value
// cannot be parsed as the requested type.
var ErrParse = errors.New("unable to parse")

func newErrNotExist(key string) error {
	return fmt.Errorf("%s: %w", key, ErrNotExist)
}

func newErrParse(key, val string, err error) error {
	return fmt.Errorf("%s=%s: %w: %w", key, val, ErrParse, err)
}

// mustParse looks up key and parses it with parse. If key is unset and
// def is non-empty, the first element of def is returned. If key is
// unset and def is empty, it panics with ErrNotExist. If key is set
// but parse returns an error, it panics with ErrParse.
func mustParse[T any](key string, parse func(string) (T, error), def []T) T {
	val, ok := os.LookupEnv(key)
	if !ok {
		if len(def) > 0 {
			return def[0]
		}
		panic(newErrNotExist(key))
	}
	result, err := parse(val)
	if err != nil {
		panic(newErrParse(key, val, err))
	}
	return result
}

// Load reads .env files into the process environment via godotenv.
// Missing files are not an error. Existing environment variables are
// not overwritten.
func Load(paths ...string) error {
	err := godotenv.Load(paths...)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	return nil
}

// MustBool returns the value of key parsed as a bool. If key is unset,
// the first def is returned. It panics with ErrNotExist when key is
// unset and no default is given, and with ErrParse when the value
// cannot be parsed. Accepted values are those recognised by
// [strconv.ParseBool]: 1, t, T, TRUE, true, True, 0, f, F, FALSE,
// false, False.
func MustBool(key string, def ...bool) bool {
	return mustParse(key, strconv.ParseBool, def)
}

// MustDuration returns the value of key parsed as a [time.Duration].
// If key is unset, the first def is returned. It panics with
// ErrNotExist when key is unset and no default is given, and with
// ErrParse when the value cannot be parsed.
func MustDuration(key string, def ...time.Duration) time.Duration {
	return mustParse(key, time.ParseDuration, def)
}

// MustInt returns the value of key parsed as an int. If key is unset,
// the first def is returned. It panics with ErrNotExist when key is
// unset and no default is given, and with ErrParse when the value
// cannot be parsed.
func MustInt(key string, def ...int) int {
	return mustParse(key, strconv.Atoi, def)
}

// MustString returns the value of key. If key is unset, the first def
// is returned. It panics with ErrNotExist when key is unset and no
// default is given. It never panics with ErrParse since any string
// value is valid.
func MustString(key string, def ...string) string {
	return mustParse(key, func(s string) (string, error) { return s, nil }, def)
}
