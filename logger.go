package negronilogrus

import (
	"fmt"
	"github.com/blackjack/syslog"
	"os"
	"path"
	"runtime"
	"strings"
)

func init() {
	progname := os.Args[0]
	base := path.Base(progname)
	syslog.Openlog(base, syslog.LOG_PID, syslog.LOG_USER)
}

func llogf(p syslog.Priority, format string, a ...interface{}) {
	_, f, l, _ := runtime.Caller(2)
	fd := strings.Split(f, "/src/")
	if len(fd) == 2 {
		filedata := fmt.Sprintf("%s.%d: ", fd[1], l)
		format = filedata + format
	}
	syslog.Syslogf(p, format, a)
}
func llog(p syslog.Priority, msg string) {
	_, f, l, _ := runtime.Caller(2)
	fd := strings.Split(f, "/src/")
	if len(fd) == 2 {
		filedata := fmt.Sprintf("%s.%d: ", fd[1], l)
		msg = filedata + msg
	}
	syslog.Syslog(p, msg)
}
func Emerg(msg string) {
	llog(syslog.LOG_EMERG, msg)
}

func Emergf(format string, a ...interface{}) {
	llogf(syslog.LOG_EMERG, format, a...)
}

func Alert(msg string) {
	llog(syslog.LOG_ALERT, msg)
}

func Alertf(format string, a ...interface{}) {
	llogf(syslog.LOG_ALERT, format, a...)
}

func Crit(msg string) {
	llog(syslog.LOG_CRIT, msg)
}

func Critf(format string, a ...interface{}) {
	llogf(syslog.LOG_CRIT, format, a...)
}

func Err(msg string) {
	llog(syslog.LOG_ERR, msg)
}

func Errf(format string, a ...interface{}) {
	llogf(syslog.LOG_ERR, format, a...)
}

func Warning(msg string) {
	llog(syslog.LOG_WARNING, msg)
}

func Warningf(format string, a ...interface{}) {
	llogf(syslog.LOG_WARNING, format, a...)
}

func Notice(msg string) {
	llog(syslog.LOG_NOTICE, msg)
}

func Noticef(format string, a ...interface{}) {
	llogf(syslog.LOG_NOTICE, format, a...)
}

func Info(msg string) {
	llog(syslog.LOG_INFO, msg)
}

func Infof(format string, a ...interface{}) {
	llogf(syslog.LOG_INFO, format, a...)
}

func Debug(msg string) {
	llog(syslog.LOG_DEBUG, msg)
}

func Debugf(format string, a ...interface{}) {
	llogf(syslog.LOG_DEBUG, format, a...)
}
