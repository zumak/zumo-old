package log

import (
	"fmt"
	"log"
	"os"
)

var (
	fatal = log.New(os.Stderr, "[FATAL] ", log.Llongfile|log.Ltime|log.Ldate)
	err   = log.New(os.Stderr, "[ERROR] ", log.Llongfile|log.Ltime|log.Ldate)
	warn  = log.New(os.Stderr, "[WARN ] ", log.Llongfile|log.Ltime|log.Ldate)
	info  = log.New(os.Stdout, "[INFO ] ", log.Llongfile|log.Ltime|log.Ldate)
	debug = log.New(os.Stdout, "[DEBUG] ", log.Llongfile|log.Ltime|log.Ldate)
)

func Fatal(format string, v ...interface{}) {
	fatal.Output(3, fmt.Sprintf(format+"\n", v...))
}
func Err(e error) {
	err.Output(3, fmt.Sprintf("%s\n", e.Error()))
}
func Error(format string, v ...interface{}) {
	err.Output(3, fmt.Sprintf(format+"\n", v...))
}
func Warn(format string, v ...interface{}) {
	warn.Output(3, fmt.Sprintf(format+"\n", v...))
}
func Info(format string, v ...interface{}) {
	info.Output(3, fmt.Sprintf(format+"\n", v...))
}
func Debug(format string, v ...interface{}) {
	debug.Output(3, fmt.Sprintf(format+"\n", v...))
}
