package logger

import (
	"fmt"
	"io"
	"log"
)

type emergLogger struct {
	w LogWriter
}

func newEmergLogger(w LogWriter) *emergLogger {
	if w == nil {
		return nil
	}

	l := new(emergLogger)
	l.w = w

	return l
}

func (l emergLogger) Writer() io.Writer {
	return l.w
}

func (l emergLogger) Level() string {
	return "emergency"
}

func (l emergLogger) Emerg(m string) {
	if err := l.w.Emerg(m); err != nil {
		log.Fatalln(err)
	}
}

func (l emergLogger) Emergf(format string, a ...any) {
	l.Emerg(fmt.Sprintf(format, a...))
}

func (l emergLogger) Alert(m string)                   {}
func (l emergLogger) Alertf(format string, a ...any)   {}
func (l emergLogger) Crit(m string)                    {}
func (l emergLogger) Critf(format string, a ...any)    {}
func (l emergLogger) Err(m string)                     {}
func (l emergLogger) Errf(format string, a ...any)     {}
func (l emergLogger) Warning(m string)                 {}
func (l emergLogger) Warningf(format string, a ...any) {}
func (l emergLogger) Notice(m string)                  {}
func (l emergLogger) Noticef(format string, a ...any)  {}
func (l emergLogger) Info(m string)                    {}
func (l emergLogger) Infof(format string, a ...any)    {}
func (l emergLogger) Debug(m string)                   {}
func (l emergLogger) Debugf(format string, a ...any)   {}
func (l emergLogger) Trace(m string)                   {}
func (l emergLogger) Tracef(format string, a ...any)   {}
