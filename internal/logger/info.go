package logger

import (
	"fmt"
	"io"
	"log"
)

type infoLogger struct {
	w LogWriter
}

func newInfoLogger(w LogWriter) *infoLogger {
	if w == nil {
		return nil
	}

	l := new(infoLogger)
	l.w = w

	return l
}

func (l infoLogger) Writer() io.Writer {
	return l.w
}

func (l infoLogger) Level() string {
	return "info"
}

func (l infoLogger) Emerg(m string) {
	if err := l.w.Emerg(m); err != nil {
		log.Fatalln(err)
	}
}

func (l infoLogger) Emergf(format string, a ...any) {
	l.Emerg(fmt.Sprintf(format, a...))
}

func (l infoLogger) Alert(m string) {
	if err := l.w.Alert(m); err != nil {
		log.Fatalln(err)
	}
}

func (l infoLogger) Alertf(format string, a ...any) {
	l.Alert(fmt.Sprintf(format, a...))
}

func (l infoLogger) Crit(m string) {
	if err := l.w.Crit(m); err != nil {
		log.Fatalln(err)
	}
}

func (l infoLogger) Critf(format string, a ...any) {
	l.Crit(fmt.Sprintf(format, a...))
}

func (l infoLogger) Err(m string) {
	if err := l.w.Err(m); err != nil {
		log.Fatalln(err)
	}
}

func (l infoLogger) Errf(format string, a ...any) {
	l.Err(fmt.Sprintf(format, a...))
}

func (l infoLogger) Warning(m string) {
	if err := l.w.Warning(m); err != nil {
		log.Fatalln(err)
	}
}

func (l infoLogger) Warningf(format string, a ...any) {
	l.Warning(fmt.Sprintf(format, a...))
}

func (l infoLogger) Notice(m string) {
	if err := l.w.Notice(m); err != nil {
		log.Fatalln(err)
	}
}

func (l infoLogger) Noticef(format string, a ...any) {
	l.Notice(fmt.Sprintf(format, a...))
}

func (l infoLogger) Info(m string) {
	if err := l.w.Info(m); err != nil {
		log.Fatalln(err)
	}
}

func (l infoLogger) Infof(format string, a ...any) {
	l.Info(fmt.Sprintf(format, a...))
}

func (l infoLogger) Debug(m string)                 {}
func (l infoLogger) Debugf(format string, a ...any) {}
func (l infoLogger) Trace(m string)                 {}
func (l infoLogger) Tracef(format string, a ...any) {}
