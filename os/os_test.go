package os

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOsEnv(t *testing.T) {

	e := New()
	assert.Equal(t, e.Environ(), os.Environ())

	os.Setenv("ABC", "DEF")
	assert.Equal(t, "DEF", e.Getenv("ABC"))

	e.Setenv("ABC", "GHI")
	assert.Equal(t, "GHI", os.Getenv("ABC"))

	val, ok := e.LookupEnv("ABC")
	assert.Equal(t, "GHI", val)
	assert.True(t, ok)

	val, ok = e.LookupEnv("nope")
	assert.Equal(t, "", val)
	assert.False(t, ok)

	e.Clearenv()
	assert.Equal(t, []string{}, os.Environ())
}
