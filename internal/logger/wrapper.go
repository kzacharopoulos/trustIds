package logger

import (
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
	"time"
)

var clearLogFlags = 0x00

var clrReset = "\033[0m"
var clrRed = "\033[31m"
var clrYellow = "\033[33m"
var clrBlue = "\033[34m"
var clrMagenta = "\033[35m"
var clrCyan = "\033[36m"

type writerWrapper struct {
	w        io.Writer
	isSyslog bool
	isSimple bool
	emerg    string
	alert    string
	crit     string
	err      string
	warn     string
	note     string
	info     string
	debug    string
}

func newWrapper(w io.Writer, tag, color bool) *writerWrapper {
	ww := new(writerWrapper)
	ww.w = w

	log.SetOutput(w)
	log.SetFlags(clearLogFlags)

	switch w.(type) {
	case *syslog.Writer:
		log.SetFlags(log.Lmsgprefix)

		if color {
			log.SetPrefix("  (" + clrMagenta + "ext" + clrReset + ") ")
			ww.withColorSyslog()
		} else {
			log.SetPrefix("  (ext) ")
			ww.withoutColorSyslog()
		}
	default:
		if !tag {
			ww.withoutTag()
			break
		}

		fulltag := ww.withTag()
		log.SetFlags(log.LstdFlags | log.Lmsgprefix)

		if color {
			log.SetPrefix(fulltag + ":   (" + clrMagenta + "ext" + clrReset + ") ")
			ww.withColor(fulltag)
		} else {
			log.SetPrefix(fulltag + ":   (ext) ")
			ww.withoutColor(fulltag)
		}
	}

	return ww
}

func (ww *writerWrapper) withColorSyslog() {
	ww.isSyslog = true
	ww.isSimple = false
	ww.emerg = fmt.Sprintf("(%s)", clrRed+"emerg"+clrReset)
	ww.alert = fmt.Sprintf("(%s)", clrRed+"alert"+clrReset)
	ww.crit = fmt.Sprintf(" (%s)", clrRed+"crit"+clrReset)
	ww.err = fmt.Sprintf("(%s)", clrRed+"error"+clrReset)
	ww.warn = fmt.Sprintf(" (%s)", clrBlue+"warn"+clrReset)
	ww.note = fmt.Sprintf(" (%s)", clrYellow+"note"+clrReset)
	ww.info = fmt.Sprintf(" (%s)", "info")
	ww.debug = fmt.Sprintf("(%s)", clrCyan+"debug"+clrReset)
}

func (ww *writerWrapper) withoutColorSyslog() {
	ww.isSyslog = true
	ww.isSimple = false
	ww.emerg = fmt.Sprintf("(%s)", "emerg")
	ww.alert = fmt.Sprintf("(%s)", "alert")
	ww.crit = fmt.Sprintf(" (%s)", "crit")
	ww.err = fmt.Sprintf("(%s)", "error")
	ww.warn = fmt.Sprintf(" (%s)", "warn")
	ww.note = fmt.Sprintf(" (%s)", "note")
	ww.info = fmt.Sprintf(" (%s)", "info")
	ww.debug = fmt.Sprintf("(%s)", "debug")
}

func (ww *writerWrapper) withColor(tag string) {
	ww.isSyslog = false
	ww.isSimple = false
	ww.emerg = fmt.Sprintf("%s: (%s)", tag, clrRed+"emerg"+clrReset)
	ww.alert = fmt.Sprintf("%s: (%s)", tag, clrRed+"alert"+clrReset)
	ww.crit = fmt.Sprintf("%s:  (%s)", tag, clrRed+"crit"+clrReset)
	ww.err = fmt.Sprintf("%s: (%s)", tag, clrRed+"error"+clrReset)
	ww.warn = fmt.Sprintf("%s:  (%s)", tag, clrBlue+"warn"+clrReset)
	ww.note = fmt.Sprintf("%s:  (%s)", tag, clrYellow+"note"+clrReset)
	ww.info = fmt.Sprintf("%s:  (%s)", tag, "info")
	ww.debug = fmt.Sprintf("%s: (%s)", tag, clrCyan+"debug"+clrReset)
}

func (ww *writerWrapper) withoutColor(tag string) {
	ww.isSyslog = false
	ww.isSimple = false
	ww.emerg = fmt.Sprintf("%s: (%s)", tag, "emerg")
	ww.alert = fmt.Sprintf("%s: (%s)", tag, "alert")
	ww.crit = fmt.Sprintf("%s:  (%s)", tag, "crit")
	ww.err = fmt.Sprintf("%s: (%s)", tag, "error")
	ww.warn = fmt.Sprintf("%s:  (%s)", tag, "warn")
	ww.note = fmt.Sprintf("%s:  (%s)", tag, "note")
	ww.info = fmt.Sprintf("%s:  (%s)", tag, "info")
	ww.debug = fmt.Sprintf("%s: (%s)", tag, "debug")
}

func (ww *writerWrapper) withoutTag() {
	ww.isSyslog = false
	ww.isSimple = true
}

func (ww *writerWrapper) withTag() string {
	fulltag := ""

	hname, err := os.Hostname()
	if err != nil {
		log.Println(err)
	} else {
		fulltag += hname + " "
	}

	fulltag += filepath.Base(os.Args[0])

	return fmt.Sprintf("%s[%d]", fulltag, os.Getpid())
}

func (wrap writerWrapper) msgBytes(level, m string) []byte {
	if wrap.isSyslog {
		return []byte(fmt.Sprintf("%s %s", level, m))
	}

	if wrap.isSimple {
		return []byte(m + "\n")
	}

	ts := time.Now().Local().Format("2006/01/02 15:04:05")
	return []byte(fmt.Sprintf("%s %s %s\n", ts, level, m))
}

func (wrap writerWrapper) Write(p []byte) (int, error) {
	return wrap.w.Write(p)
}

func (wrap writerWrapper) Emerg(m string) error {
	_, err := wrap.w.Write(wrap.msgBytes(wrap.emerg, m))
	return err
}

func (wrap writerWrapper) Alert(m string) error {
	_, err := wrap.w.Write(wrap.msgBytes(wrap.alert, m))
	return err
}

func (wrap writerWrapper) Crit(m string) error {
	_, err := wrap.w.Write(wrap.msgBytes(wrap.crit, m))
	return err
}

func (wrap writerWrapper) Err(m string) error {
	_, err := wrap.w.Write(wrap.msgBytes(wrap.err, m))
	return err
}

func (wrap writerWrapper) Warning(m string) error {
	_, err := wrap.w.Write(wrap.msgBytes(wrap.warn, m))
	return err
}

func (wrap writerWrapper) Notice(m string) error {
	_, err := wrap.w.Write(wrap.msgBytes(wrap.note, m))
	return err
}

func (wrap writerWrapper) Info(m string) error {
	_, err := wrap.w.Write(wrap.msgBytes(wrap.info, m))
	return err
}

func (wrap writerWrapper) Debug(m string) error {
	_, err := wrap.w.Write(wrap.msgBytes(wrap.debug, m))
	return err
}
