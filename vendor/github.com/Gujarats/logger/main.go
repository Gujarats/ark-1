package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
)

// showing the line of code
const calldepth = 2

// setup logger
var myLogger *log.Logger

// setup log output to a file
var FileSave bool

func init() {
	myLogger = log.New(os.Stdout, "", log.Lshortfile)
}

func getLineError() string {
	var fileShortend string
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return ""
	}

	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			fileShortend = file[i+1:]
			break
		}
	}

	lineStr := strconv.Itoa(line)

	return fileShortend + ":" + lineStr + ": "
}

// Print debug message to os.StdOut without save it to a file
func Debug(prefix string, message ...interface{}) {
	messageString := fmt.Sprintln(message...)
	debugStyle := NewDebugStyle(prefix, messageString)
	outputLogger()
	myLogger.Output(calldepth, debugStyle)
}

// Print debug message to os.StdOut without save it to a file with specific format
func Debugf(prefix string, format string, message ...interface{}) {
	messageString := fmt.Sprintf(format, message...)
	debugStyle := NewDebugStyle(prefix, messageString)
	outputLogger()
	myLogger.Output(calldepth, debugStyle)
}

//Set the output to a file
func DebugToFile() {
	dir := currentDir()
	loggerFileName := getFileName()
	file := &CustomFile{createLogFile(dir, loggerFileName)}
	myLogger.SetOutput(file)
	FileSave = true
}

// define where the output should be display
func outputLogger() {
	if !FileSave {
		myLogger.SetOutput(new(customWriter))
	}
}
