package mock

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

	e.Clearenv()
	assert.Equal(t, []string{}, e.Environ())
}
