package logger

import (
	"fmt"
	"io"
	"log"
)

type noticeLogger struct {
	w LogWriter
}

func newNoticeLogger(w LogWriter) *noticeLogger {
	if w == nil {
		return nil
	}

	l := new(noticeLogger)
	l.w = w

	return l
}

func (l noticeLogger) Writer() io.Writer {
	return l.w
}

func (l noticeLogger) Level() string {
	return "notice"
}

func (l noticeLogger) Emerg(m string) {
	if err := l.w.Emerg(m); err != nil {
		log.Fatalln(err)
	}
}

func (l noticeLogger) Emergf(format string, a ...any) {
	l.Emerg(fmt.Sprintf(format, a...))
}

func (l noticeLogger) Alert(m string) {
	if err := l.w.Alert(m); err != nil {
		log.Fatalln(err)
	}
}

func (l noticeLogger) Alertf(format string, a ...any) {
	l.Alert(fmt.Sprintf(format, a...))
}

func (l noticeLogger) Crit(m string) {
	if err := l.w.Crit(m); err != nil {
		log.Fatalln(err)
	}
}

func (l noticeLogger) Critf(format string, a ...any) {
	l.Crit(fmt.Sprintf(format, a...))
}

func (l noticeLogger) Err(m string) {
	if err := l.w.Err(m); err != nil {
		log.Fatalln(err)
	}
}

func (l noticeLogger) Errf(format string, a ...any) {
	l.Err(fmt.Sprintf(format, a...))
}

func (l noticeLogger) Warning(m string) {
	if err := l.w.Warning(m); err != nil {
		log.Fatalln(err)
	}
}

func (l noticeLogger) Warningf(format string, a ...any) {
	l.Warning(fmt.Sprintf(format, a...))
}

func (l noticeLogger) Notice(m string) {
	if err := l.w.Notice(m); err != nil {
		log.Fatalln(err)
	}
}

func (l noticeLogger) Noticef(format string, a ...any) {
	l.Notice(fmt.Sprintf(format, a...))
}

func (l noticeLogger) Info(m string)                  {}
func (l noticeLogger) Infof(format string, a ...any)  {}
func (l noticeLogger) Debug(m string)                 {}
func (l noticeLogger) Debugf(format string, a ...any) {}
func (l noticeLogger) Trace(m string)                 {}
func (l noticeLogger) Tracef(format string, a ...any) {}
