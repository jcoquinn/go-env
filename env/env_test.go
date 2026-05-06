package env_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jcoquinn/go-env/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func key(t *testing.T) string {
	t.Helper()
	return "_goenv_" + t.Name()
}

// assertPanicsWithErr is a deferred helper that asserts the surrounding
// test panicked with an error that wraps target (per errors.Is).
func assertPanicsWithErr(t *testing.T, target error) {
	t.Helper()
	r := recover()
	assert.NotNil(t, r, "panic expected")
	err, ok := r.(error)
	require.True(t, ok, "panic value was not an error: %v", r)
	assert.ErrorIs(t, err, target)
}

// --- Load ---

func TestLoad_MissingFile(t *testing.T) {
	err := env.Load()
	assert.NoError(t, err)

	err = env.Load(filepath.Join(t.TempDir(), "missing.env"))
	assert.NoError(t, err)
}

func TestLoad_ExistingFile(t *testing.T) {
	k := key(t)
	path := filepath.Join(t.TempDir(), ".env")
	err := os.WriteFile(path, []byte(k+"=hello\n"), 0o600)
	assert.NoError(t, err)

	err = env.Load(path)
	assert.NoError(t, err)
	assert.Equal(t, "hello", os.Getenv(k))
}

func TestLoad_MalformedFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".env")
	err := os.WriteFile(path, []byte("invalid line\n"), 0o600)
	assert.NoError(t, err)

	err = env.Load(path)
	assert.Error(t, err)
}

// --- MustBool ---

func TestMustBool_ReturnsValue(t *testing.T) {
	t.Setenv(key(t), "true")
	assert.Equal(t, true, env.MustBool(key(t)))
}

func TestMustBool_ReturnsDefault(t *testing.T) {
	assert.Equal(t, false, env.MustBool(key(t), false))
}

func TestMustBool_PanicsWhenUnset(t *testing.T) {
	defer assertPanicsWithErr(t, env.ErrNotExist)
	env.MustBool(key(t))
}

func TestMustBool_PanicsOnParseError(t *testing.T) {
	t.Setenv(key(t), "notabool")
	defer assertPanicsWithErr(t, env.ErrParse)
	env.MustBool(key(t))
}

// --- MustDuration ---

func TestMustDuration_ReturnsValue(t *testing.T) {
	t.Setenv(key(t), "10s")
	assert.Equal(t, 10*time.Second, env.MustDuration(key(t)))
}

func TestMustDuration_ReturnsDefault(t *testing.T) {
	assert.Equal(t, 30*time.Second, env.MustDuration(key(t), 30*time.Second))
}

func TestMustDuration_PanicsWhenUnset(t *testing.T) {
	defer assertPanicsWithErr(t, env.ErrNotExist)
	env.MustDuration(key(t))
}

func TestMustDuration_PanicsOnParseError(t *testing.T) {
	t.Setenv(key(t), "1x")
	defer assertPanicsWithErr(t, env.ErrParse)
	env.MustDuration(key(t))
}

// --- MustInt ---

func TestMustInt_ReturnsValue(t *testing.T) {
	t.Setenv(key(t), "42")
	assert.Equal(t, 42, env.MustInt(key(t)))
}

func TestMustInt_ReturnsDefault(t *testing.T) {
	assert.Equal(t, 8080, env.MustInt(key(t), 8080))
}

func TestMustInt_PanicsWhenUnset(t *testing.T) {
	defer assertPanicsWithErr(t, env.ErrNotExist)
	env.MustInt(key(t))
}

func TestMustInt_PanicsOnParseError(t *testing.T) {
	t.Setenv(key(t), "abc")
	defer assertPanicsWithErr(t, env.ErrParse)
	env.MustInt(key(t))
}

// --- MustString ---

func TestMustString_ReturnsValue(t *testing.T) {
	t.Setenv(key(t), "hello")
	assert.Equal(t, "hello", env.MustString(key(t)))
}

func TestMustString_ReturnsDefault(t *testing.T) {
	assert.Equal(t, "hello", env.MustString(key(t), "hello"))
}

func TestMustString_ReturnsEmptyWhenSet(t *testing.T) {
	t.Setenv(key(t), "")
	assert.Equal(t, "", env.MustString(key(t), "hello"))
}

func TestMustString_PanicsWhenUnset(t *testing.T) {
	defer assertPanicsWithErr(t, env.ErrNotExist)
	env.MustString(key(t))
}
