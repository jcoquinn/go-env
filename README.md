# go-env

Typed helpers for reading environment variables in Go. Each helper returns the
parsed value of a named variable, falls back to an optional default, or panics
with an error. Intended for use at startup to keep your application fast-fail
and configuration explicit.

## Install

```sh
go get github.com/jcoquinn/go-env/env
```

## Usage

```go
import (
	"time"

	"github.com/jcoquinn/go-env/env"
)

port    := env.MustInt("PORT", 8000)           // int with default
debug   := env.MustBool("DEBUG", false)         // bool with default
timeout := env.MustDuration("TIMEOUT", 30*time.Second) // duration with default
dsn     := env.MustString("DATABASE_URL")       // required string, panics if unset
```

**Wrap calls in a `FromEnv` constructor that converts panics to errors:**

```go
import (
	"fmt"

	"github.com/jcoquinn/go-env/env"
)

type Config struct {
	Port int
}

func FromEnv() (cfg *Config, err error) {
	defer func() {
		if r := recover(); r != nil {
			cfg = nil
			var ok bool
			if err, ok = r.(error); !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	if err := env.Load(); err != nil {
		return nil, err
	}

	cfg = &Config{}
	cfg.Port = env.MustInt("APP_PORT")

	return cfg, nil
}
```

Each helper follows the same contract:

- Variable set and parseable → returns the parsed value.
- Variable set and not parseable → panics with a wrapped `env.ErrParse`.
- Variable unset, default provided → returns the default.
- Variable unset, no default → panics with a wrapped `env.ErrNotExist`.

Use `errors.Is` inside a `recover` to distinguish the two failure modes:

```go
defer func() {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if ok && errors.Is(err, env.ErrNotExist) {
			// handle missing required variable
		}
	}
}()
```

## Loading .env files

`env.Load` reads one or more `.env` files into the process environment via
[godotenv]. Missing files are silently ignored; existing environment variables
are never overwritten.

```go
// Load .env (default) — no error if file is absent
if err := env.Load(); err != nil {
    log.Fatal(err)
}

// Load a specific file
if err := env.Load(".env.local"); err != nil {
    log.Fatal(err)
}
```

[godotenv]: https://github.com/joho/godotenv
