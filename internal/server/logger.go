package server

import "github.com/rs/zerolog"

type PrintfLogger struct {
	level   zerolog.Level
	zerolog *zerolog.Logger
}

func (l *PrintfLogger) Printf(format string, args ...interface{}) {
	level := l.zerolog.Level(l.level)
	level.Printf(format, args)
}

// LoggerPrintf returns a new PrintfLogger given a level.
func LoggerPrintf(logger *zerolog.Logger, level zerolog.Level) *PrintfLogger {
	return &PrintfLogger{
		level:   level,
		zerolog: logger,
	}
}
