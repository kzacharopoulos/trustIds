package logger

import (
	"fmt"
	"io"
	"log"
)

type critLogger struct {
	w LogWriter
}

func newCritLogger(w LogWriter) *critLogger {
	if w == nil {
		return nil
	}

	l := new(critLogger)
	l.w = w

	return l
}

func (l critLogger) Writer() io.Writer {
	return l.w
}

func (l critLogger) Level() string {
	return "critical"
}

func (l critLogger) Emerg(m string) {
	if err := l.w.Emerg(m); err != nil {
		log.Fatalln(err)
	}
}

func (l critLogger) Emergf(format string, a ...any) {
	l.Emerg(fmt.Sprintf(format, a...))
}

func (l critLogger) Alert(m string) {
	if err := l.w.Alert(m); err != nil {
		log.Fatalln(err)
	}
}

func (l critLogger) Alertf(format string, a ...any) {
	l.Alert(fmt.Sprintf(format, a...))
}

func (l critLogger) Crit(m string) {
	if err := l.w.Crit(m); err != nil {
		log.Fatalln(err)
	}
}

func (l critLogger) Critf(format string, a ...any) {
	l.Crit(fmt.Sprintf(format, a...))
}

func (l critLogger) Err(m string)                     {}
func (l critLogger) Errf(format string, a ...any)     {}
func (l critLogger) Warning(m string)                 {}
func (l critLogger) Warningf(format string, a ...any) {}
func (l critLogger) Notice(m string)                  {}
func (l critLogger) Noticef(format string, a ...any)  {}
func (l critLogger) Info(m string)                    {}
func (l critLogger) Infof(format string, a ...any)    {}
func (l critLogger) Debug(m string)                   {}
func (l critLogger) Debugf(format string, a ...any)   {}
func (l critLogger) Trace(m string)                   {}
func (l critLogger) Tracef(format string, a ...any)   {}
