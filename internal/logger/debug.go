package logger

import (
	"fmt"
	"io"
	"log"
)

type debugLogger struct {
	w LogWriter
}

func newDebugLogger(w LogWriter) *debugLogger {
	if w == nil {
		return nil
	}

	l := new(debugLogger)
	l.w = w

	return l
}

func (l debugLogger) Writer() io.Writer {
	return l.w
}

func (l debugLogger) Level() string {
	return "debug"
}

func (l debugLogger) Emerg(m string) {
	if err := l.w.Emerg(m); err != nil {
		log.Fatalln(err)
	}
}

func (l debugLogger) Emergf(format string, a ...any) {
	l.Emerg(fmt.Sprintf(format, a...))
}

func (l debugLogger) Alert(m string) {
	if err := l.w.Alert(m); err != nil {
		log.Fatalln(err)
	}
}

func (l debugLogger) Alertf(format string, a ...any) {
	l.Alert(fmt.Sprintf(format, a...))
}

func (l debugLogger) Crit(m string) {
	if err := l.w.Crit(m); err != nil {
		log.Fatalln(err)
	}
}

func (l debugLogger) Critf(format string, a ...any) {
	l.Crit(fmt.Sprintf(format, a...))
}

func (l debugLogger) Err(m string) {
	if err := l.w.Err(m); err != nil {
		log.Fatalln(err)
	}
}

func (l debugLogger) Errf(format string, a ...any) {
	l.Err(fmt.Sprintf(format, a...))
}

func (l debugLogger) Warning(m string) {
	if err := l.w.Warning(m); err != nil {
		log.Fatalln(err)
	}
}

func (l debugLogger) Warningf(format string, a ...any) {
	l.Warning(fmt.Sprintf(format, a...))
}

func (l debugLogger) Notice(m string) {
	if err := l.w.Notice(m); err != nil {
		log.Fatalln(err)
	}
}

func (l debugLogger) Noticef(format string, a ...any) {
	l.Notice(fmt.Sprintf(format, a...))
}

func (l debugLogger) Info(m string) {
	if err := l.w.Info(m); err != nil {
		log.Fatalln(err)
	}
}

func (l debugLogger) Infof(format string, a ...any) {
	l.Info(fmt.Sprintf(format, a...))
}

func (l debugLogger) Debug(m string) {
	if err := l.w.Debug(m); err != nil {
		log.Fatalln(err)
	}
}

func (l debugLogger) Debugf(format string, a ...any) {
	l.Debug(fmt.Sprintf(format, a...))
}

func (l debugLogger) Trace(m string)                 {}
func (l debugLogger) Tracef(format string, a ...any) {}
