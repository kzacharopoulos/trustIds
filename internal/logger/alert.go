package logger

import (
	"fmt"
	"io"
	"log"
)

type alertLogger struct {
	w LogWriter
}

func newAlertLogger(w LogWriter) *alertLogger {
	if w == nil {
		return nil
	}

	l := new(alertLogger)
	l.w = w

	return l
}

func (l alertLogger) Writer() io.Writer {
	return l.w
}

func (l alertLogger) Level() string {
	return "alert"
}

func (l alertLogger) Emerg(m string) {
	if err := l.w.Emerg(m); err != nil {
		log.Fatalln(err)
	}
}

func (l alertLogger) Emergf(format string, a ...any) {
	l.Emerg(fmt.Sprintf(format, a...))
}

func (l alertLogger) Alert(m string) {
	if err := l.w.Alert(m); err != nil {
		log.Fatalln(err)
	}
}

func (l alertLogger) Alertf(format string, a ...any) {
	l.Alert(fmt.Sprintf(format, a...))
}

func (l alertLogger) Crit(m string)                    {}
func (l alertLogger) Critf(format string, a ...any)    {}
func (l alertLogger) Err(m string)                     {}
func (l alertLogger) Errf(format string, a ...any)     {}
func (l alertLogger) Warning(m string)                 {}
func (l alertLogger) Warningf(format string, a ...any) {}
func (l alertLogger) Notice(m string)                  {}
func (l alertLogger) Noticef(format string, a ...any)  {}
func (l alertLogger) Info(m string)                    {}
func (l alertLogger) Infof(format string, a ...any)    {}
func (l alertLogger) Debug(m string)                   {}
func (l alertLogger) Debugf(format string, a ...any)   {}
func (l alertLogger) Trace(m string)                   {}
func (l alertLogger) Tracef(format string, a ...any)   {}
