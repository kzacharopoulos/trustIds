package logger

import (
	"io"
	"log"
	"log/syslog"
	"os"
)

type Logger interface {
	Writer() io.Writer
	Level() string
	Emerg(m string)
	Emergf(format string, a ...any)
	Alert(m string)
	Alertf(format string, a ...any)
	Crit(m string)
	Critf(format string, a ...any)
	Err(m string)
	Errf(format string, a ...any)
	Warning(m string)
	Warningf(format string, a ...any)
	Notice(m string)
	Noticef(format string, a ...any)
	Info(m string)
	Infof(format string, a ...any)
	Debug(m string)
	Debugf(format string, a ...any)
	Trace(m string)
	Tracef(format string, a ...any)
}

type LogWriter interface {
	io.Writer
	Emerg(m string) error
	Alert(m string) error
	Crit(m string) error
	Err(m string) error
	Warning(m string) error
	Notice(m string) error
	Info(m string) error
	Debug(m string) error
}

func stringToPriority(name string) syslog.Priority {
	switch name {
	case "emerg":
		return syslog.LOG_EMERG
	case "emergency":
		return syslog.LOG_EMERG
	case "alert":
		return syslog.LOG_ALERT
	case "crit":
		return syslog.LOG_CRIT
	case "critical":
		return syslog.LOG_CRIT
	case "err":
		return syslog.LOG_ERR
	case "error":
		return syslog.LOG_ERR
	case "warn":
		return syslog.LOG_WARNING
	case "warning":
		return syslog.LOG_WARNING
	case "notice":
		return syslog.LOG_NOTICE
	case "debug":
		return syslog.LOG_DEBUG
	case "trace":
		return syslog.LOG_LOCAL0
	default:
		return syslog.LOG_INFO
	}
}

func priToString(pri syslog.Priority) string {
	switch pri {
	case syslog.LOG_EMERG:
		return "emerg"
	case syslog.LOG_ALERT:
		return "alert"
	case syslog.LOG_CRIT:
		return "critical"
	case syslog.LOG_ERR:
		return "error"
	case syslog.LOG_WARNING:
		return "warning"
	case syslog.LOG_NOTICE:
		return "notice"
	case syslog.LOG_DEBUG:
		return "debug"
	case syslog.LOG_LOCAL0:
		return "trace"
	default:
		return "info"
	}
}

func loggerDescription(w io.Writer, max string, tag, color bool) {
	pri := stringToPriority(max)
	priString := priToString(pri)

	switch w.(type) {
	case *syslog.Writer:
		log.Printf("logger: type=syslog, maxPriority=%s, color=%v", priString, color)
	case *os.File:
		if w == os.Stdout {
			log.Printf("logger: type=stdout, maxPriority=%s, tag=%v, color=%v", priString, tag, color)
		} else if w == os.Stderr {
			log.Printf("logger: type=stderr, maxPriority=%s, tag=%v, color=%v", priString, tag, color)
		} else {
			log.Printf("logger: type=file, maxPriority=%s, tag=%v, color=%v", priString, tag, color)
		}
	default:
		log.Printf("logger: type=io.Writer, maxPriority=%s, tag=%v, color=%v", priString, tag, color)
	}
}

func New(w io.Writer, maxPriority string, tag, color bool) Logger {
	wrap := newWrapper(w, tag, color)

	loggerDescription(w, maxPriority, tag, color)

	switch stringToPriority(maxPriority) {
	case syslog.LOG_ALERT:
		return newAlertLogger(wrap)
	case syslog.LOG_CRIT:
		return newCritLogger(wrap)
	case syslog.LOG_DEBUG:
		return newDebugLogger(wrap)
	case syslog.LOG_EMERG:
		return newEmergLogger(wrap)
	case syslog.LOG_ERR:
		return newErrLogger(wrap)
	case syslog.LOG_NOTICE:
		return newNoticeLogger(wrap)
	case syslog.LOG_LOCAL0:
		return newTraceLogger(wrap)
	case syslog.LOG_WARNING:
		return newWarnLogger(wrap)
	default:
		return newInfoLogger(wrap)
	}
}

func NewFromType(logtype, file, appname, level string, tag, color bool) Logger {
	switch logtype {
	case "stdout":
		return NewStdout(level, tag, color)
	case "stderr":
		return NewStderr(level, tag, color)
	case "syslog":
		return NewSyslog(level, appname, color)
	case "file":
		return NewFilename(file, level, tag, color)
	case "none":
		return NewNone()
	default:
		return NewSimple()
	}
}

func NewNone() Logger {
	return newNoneLogger()
}

func NewSimple() Logger {
	return New(os.Stdout, "info", false, false)
}

func NewStdout(level string, tag, color bool) Logger {
	return New(os.Stdout, level, tag, color)
}

func NewStderr(level string, tag, color bool) Logger {
	return New(os.Stderr, level, tag, color)
}

func NewSyslog(level, name string, color bool) Logger {
	w, err := syslog.New(syslog.LOG_INFO, name)
	if err != nil {
		log.Printf("error creating logger: %s\n", err)
		return NewSimple()
	}

	return New(w, level, false, color)
}

func NewFile(file *os.File, level string, tag, color bool) Logger {
	return New(file, level, tag, color)
}

func NewFilename(filename, level string, tag, color bool) Logger {
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("error creating logger: %s\n", err)
		return NewSimple()
	}

	return NewFile(file, level, tag, color)
}

func NewWriter(w io.Writer, level string, tag, color bool) Logger {
	return New(w, level, tag, color)
}
