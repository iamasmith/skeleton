package config

import (
	"errors"
	"flag"

	"go.uber.org/zap/zapcore"
)

type ConfigT struct {
	FlagSet  *flag.FlagSet
	LogLevel zapcore.Level
}

var Config = ConfigT{
	FlagSet:  flag.NewFlagSet("standard", flag.ExitOnError),
	LogLevel: zapcore.InfoLevel,
}

// Allows us to call ParseArgs from unit tests over and over
func ResetForTest() {
	Config.FlagSet = flag.NewFlagSet("test", flag.ContinueOnError)
}

func ParseArgs(args []string) {
	flag := Config.FlagSet
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
			Config.LogLevel = level
			return nil
		},
	)
	flag.Parse(args)
}
