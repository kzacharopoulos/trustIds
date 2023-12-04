package logger

import (
	"io"
)

type noneLogger struct {
}

func newNoneLogger() *noneLogger {
	return new(noneLogger)
}

func (l noneLogger) Writer() io.Writer {
	return nil
}

func (l noneLogger) Level() string {
	return "none"
}

func (l noneLogger) Emerg(m string)                   {}
func (l noneLogger) Emergf(format string, a ...any)   {}
func (l noneLogger) Alert(m string)                   {}
func (l noneLogger) Alertf(format string, a ...any)   {}
func (l noneLogger) Crit(m string)                    {}
func (l noneLogger) Critf(format string, a ...any)    {}
func (l noneLogger) Err(m string)                     {}
func (l noneLogger) Errf(format string, a ...any)     {}
func (l noneLogger) Warning(m string)                 {}
func (l noneLogger) Warningf(format string, a ...any) {}
func (l noneLogger) Notice(m string)                  {}
func (l noneLogger) Noticef(format string, a ...any)  {}
func (l noneLogger) Info(m string)                    {}
func (l noneLogger) Infof(format string, a ...any)    {}
func (l noneLogger) Debug(m string)                   {}
func (l noneLogger) Debugf(format string, a ...any)   {}
func (l noneLogger) Trace(m string)                   {}
func (l noneLogger) Tracef(format string, a ...any)   {}
