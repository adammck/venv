package venv

import (
	"github.com/adammck/venv/mock"
	"github.com/adammck/venv/os"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOS(t *testing.T) {
	assert.IsType(t, &os.OsEnv{}, OS())
}

func TestMock(t *testing.T) {
	assert.IsType(t, &mock.MockEnv{}, Mock())
}
