# venv

[![GoDoc](https://godoc.org/github.com/adammck/venv?status.svg)](https://godoc.org/github.com/adammck/venv)
[![Build Status](https://travis-ci.org/adammck/venv.svg?branch=master)](https://travis-ci.org/adammck/venv)

This is a Go library to abstract access to environment variables.  
Like [blang/vfs](https://github.com/blang/vfs) for the env.

## Usage

```go
package main

import (
	"github.com/adammck/venv"
	"fmt"
)

func main() {
	var e venv.Env

	// The real environment

	e = e.OS()
	fmt.Sprintf("hello, %s!", env.Getenv("USER"))

	// Or use a mock environment

	e = os.Mock()
	e.Setenv("USER", "fred")
	fmt.Sprintf("hello, %s!", env.Getenv("USER"))
}
```

## License

MIT.
