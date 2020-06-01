package venv

import (
	"github.com/adammck/venv/mock"
	"github.com/adammck/venv/os"
)

type Env interface {
	Environ() []string
	Getenv(key string) string
	LookupEnv(key string) (string, bool)
	Setenv(key, value string) error
	Clearenv()
}

func OS() Env {
	return os.New()
}

func Mock() Env {
	return mock.New()
}
