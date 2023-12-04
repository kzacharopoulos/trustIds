package logger

import (
	"fmt"
	"io"
	"log"
)

type warnLogger struct {
	w LogWriter
}

func newWarnLogger(w LogWriter) *warnLogger {
	if w == nil {
		return nil
	}

	l := new(warnLogger)
	l.w = w

	return l
}

func (l warnLogger) Writer() io.Writer {
	return l.w
}

func (l warnLogger) Level() string {
	return "warning"
}

func (l warnLogger) Emerg(m string) {
	if err := l.w.Emerg(m); err != nil {
		log.Fatalln(err)
	}
}

func (l warnLogger) Emergf(format string, a ...any) {
	l.Emerg(fmt.Sprintf(format, a...))
}

func (l warnLogger) Alert(m string) {
	if err := l.w.Alert(m); err != nil {
		log.Fatalln(err)
	}
}

func (l warnLogger) Alertf(format string, a ...any) {
	l.Alert(fmt.Sprintf(format, a...))
}

func (l warnLogger) Crit(m string) {
	if err := l.w.Crit(m); err != nil {
		log.Fatalln(err)
	}
}

func (l warnLogger) Critf(format string, a ...any) {
	l.Crit(fmt.Sprintf(format, a...))
}

func (l warnLogger) Err(m string) {
	if err := l.w.Err(m); err != nil {
		log.Fatalln(err)
	}
}

func (l warnLogger) Errf(format string, a ...any) {
	l.Err(fmt.Sprintf(format, a...))
}

func (l warnLogger) Warning(m string) {
	if err := l.w.Warning(m); err != nil {
		log.Fatalln(err)
	}
}

func (l warnLogger) Warningf(format string, a ...any) {
	l.Warning(fmt.Sprintf(format, a...))
}

func (l warnLogger) Notice(m string)                 {}
func (l warnLogger) Noticef(format string, a ...any) {}
func (l warnLogger) Info(m string)                   {}
func (l warnLogger) Infof(format string, a ...any)   {}
func (l warnLogger) Debug(m string)                  {}
func (l warnLogger) Debugf(format string, a ...any)  {}
func (l warnLogger) Trace(m string)                  {}
func (l warnLogger) Tracef(format string, a ...any)  {}
