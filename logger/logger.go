package logger

import (
	"os"

	"github.com/op/go-logging"
)

var myLog = logging.MustGetLogger("myLogger")

// Info logs ...
var (
	Info    = myLog.Info
	Notice  = myLog.Notice
	Debug   = myLog.Debug
	Warning = myLog.Warning
	Error   = myLog.Error

	Infof    = myLog.Infof
	Noticef  = myLog.Noticef
	Debugf   = myLog.Debugf
	Warningf = myLog.Warningf
	Errorf   = myLog.Errorf
)

func init() {

	// Example format string. Everything except the message has a custom color
	// which is dependent on the log level. Many fields have a custom output
	// formatting too, eg. the time returns the hour down to the milli second.
	format := logging.MustStringFormatter(
		`%{color}%{time} [%{level}] %{module} %{shortfile} ▶ %{color:reset} %{message}`,
	)

	// For demo purposes, create two backend for os.Stderr.
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Formatter)
}
