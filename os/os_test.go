package os

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOsEnv(t *testing.T) {

	e := New()
	assert.Equal(t, e.Environ(), os.Environ())

	os.Setenv("ABC", "DEF")
	assert.Equal(t, "DEF", e.Getenv("ABC"))

	e.Setenv("ABC", "GHI")
	assert.Equal(t, "GHI", os.Getenv("ABC"))

	e.Clearenv()
	assert.Equal(t, []string{}, os.Environ())
}
