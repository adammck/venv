package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockEnv(t *testing.T) {

	e := New()
	assert.Equal(t, []string{}, e.Environ())

	e.Setenv("ABC", "xxx")
	assert.Equal(t, "xxx", e.Getenv("ABC"))
	assert.Equal(t, []string{"ABC=xxx"}, e.Environ())

	e.Setenv("ABC", "yyy")
	assert.Equal(t, "yyy", e.Getenv("ABC"))
	assert.Equal(t, []string{"ABC=yyy"}, e.Environ())

	val, ok := e.LookupEnv("ABC")
	assert.Equal(t, "yyy", val)
	assert.True(t, ok)

	val, ok = e.LookupEnv("nope")
	assert.Equal(t, "", val)
	assert.False(t, ok)

	e.Clearenv()
	assert.Equal(t, []string{}, e.Environ())
}
