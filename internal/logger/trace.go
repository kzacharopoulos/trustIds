package logger

import (
	"fmt"
	"io"
	"log"
)

type traceLogger struct {
	w LogWriter
}

func newTraceLogger(w LogWriter) *traceLogger {
	if w == nil {
		return nil
	}

	l := new(traceLogger)
	l.w = w

	return l
}

func (l traceLogger) Writer() io.Writer {
	return l.w
}

func (l traceLogger) Level() string {
	return "trace"
}

func (l traceLogger) Emerg(m string) {
	if err := l.w.Emerg(m); err != nil {
		log.Fatalln(err)
	}
}

func (l traceLogger) Emergf(format string, a ...any) {
	l.Emerg(fmt.Sprintf(format, a...))
}

func (l traceLogger) Alert(m string) {
	if err := l.w.Alert(m); err != nil {
		log.Fatalln(err)
	}
}

func (l traceLogger) Alertf(format string, a ...any) {
	l.Alert(fmt.Sprintf(format, a...))
}

func (l traceLogger) Crit(m string) {
	if err := l.w.Crit(m); err != nil {
		log.Fatalln(err)
	}
}

func (l traceLogger) Critf(format string, a ...any) {
	l.Crit(fmt.Sprintf(format, a...))
}

func (l traceLogger) Err(m string) {
	if err := l.w.Err(m); err != nil {
		log.Fatalln(err)
	}
}

func (l traceLogger) Errf(format string, a ...any) {
	l.Err(fmt.Sprintf(format, a...))
}

func (l traceLogger) Warning(m string) {
	if err := l.w.Warning(m); err != nil {
		log.Fatalln(err)
	}
}

func (l traceLogger) Warningf(format string, a ...any) {
	l.Warning(fmt.Sprintf(format, a...))
}

func (l traceLogger) Notice(m string) {
	if err := l.w.Notice(m); err != nil {
		log.Fatalln(err)
	}
}

func (l traceLogger) Noticef(format string, a ...any) {
	l.Notice(fmt.Sprintf(format, a...))
}

func (l traceLogger) Info(m string) {
	if err := l.w.Info(m); err != nil {
		log.Fatalln(err)
	}
}

func (l traceLogger) Infof(format string, a ...any) {
	l.Info(fmt.Sprintf(format, a...))
}

func (l traceLogger) Debug(m string) {
	if err := l.w.Debug(m); err != nil {
		log.Fatalln(err)
	}
}

func (l traceLogger) Debugf(format string, a ...any) {
	l.Debug(fmt.Sprintf(format, a...))
}

func (l traceLogger) Trace(m string) {
	if err := l.w.Debug(m); err != nil {
		log.Fatalln(err)
	}
}

func (l traceLogger) Tracef(format string, a ...any) {
	l.Trace(fmt.Sprintf(format, a...))
}
