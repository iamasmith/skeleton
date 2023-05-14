package config

import (
	"errors"
	"flag"

	"go.uber.org/zap/zapcore"
)

type ConfigT struct {
	flagSet    *flag.FlagSet
	LogLevel   zapcore.Level
	ListenBind string
}

var Config = ConfigT{
	flagSet: flag.NewFlagSet("standard", flag.ExitOnError),
}

func (c *ConfigT) setDefaults() {
	c.LogLevel = zapcore.InfoLevel
	c.ListenBind = ":8000"
}

// Allows us to call ParseArgs from unit tests over and over
func (c *ConfigT) ResetForTest() {
	c.flagSet = flag.NewFlagSet("test", flag.ContinueOnError)
}

func (c *ConfigT) ParseArgs(args []string) {
	c.setDefaults()
	flag := c.flagSet
	flag.Func(
		"level", "debug|info|warn|error (defaults info)",
		func(s string) error {
			level, has := map[string]zapcore.Level{
				"debug": zapcore.DebugLevel,
				"info":  zapcore.InfoLevel,
				"warn":  zapcore.WarnLevel,
				"error": zapcore.ErrorLevel,
			}[s]
			if !has {
				return errors.New("INVALID LOG LEVEL")
			}
			c.LogLevel = level
			return nil
		},
	)
	flag.StringVar(&c.ListenBind, "listen", ":8000", "Host/Port binding for server")
	flag.Parse(args)
}
