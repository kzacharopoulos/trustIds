package logger

import (
	"fmt"
	"io"
	"log"
)

type errLogger struct {
	w LogWriter
}

func newErrLogger(w LogWriter) *errLogger {
	if w == nil {
		return nil
	}

	l := new(errLogger)
	l.w = w

	return l
}

func (l errLogger) Writer() io.Writer {
	return l.w
}

func (l errLogger) Level() string {
	return "error"
}

func (l errLogger) Emerg(m string) {
	if err := l.w.Emerg(m); err != nil {
		log.Fatalln(err)
	}
}

func (l errLogger) Emergf(format string, a ...any) {
	l.Emerg(fmt.Sprintf(format, a...))
}

func (l errLogger) Alert(m string) {
	if err := l.w.Alert(m); err != nil {
		log.Fatalln(err)
	}
}

func (l errLogger) Alertf(format string, a ...any) {
	l.Alert(fmt.Sprintf(format, a...))
}

func (l errLogger) Crit(m string) {
	if err := l.w.Crit(m); err != nil {
		log.Fatalln(err)
	}
}

func (l errLogger) Critf(format string, a ...any) {
	l.Crit(fmt.Sprintf(format, a...))
}

func (l errLogger) Err(m string) {
	if err := l.w.Err(m); err != nil {
		log.Fatalln(err)
	}
}

func (l errLogger) Errf(format string, a ...any) {
	l.Err(fmt.Sprintf(format, a...))
}

func (l errLogger) Warning(m string)                 {}
func (l errLogger) Warningf(format string, a ...any) {}
func (l errLogger) Notice(m string)                  {}
func (l errLogger) Noticef(format string, a ...any)  {}
func (l errLogger) Info(m string)                    {}
func (l errLogger) Infof(format string, a ...any)    {}
func (l errLogger) Debug(m string)                   {}
func (l errLogger) Debugf(format string, a ...any)   {}
func (l errLogger) Trace(m string)                   {}
func (l errLogger) Tracef(format string, a ...any)   {}
