package logger

import (
	"os"
	"path/filepath"
)

type CustomFile struct {
	*os.File
}

func (c *CustomFile) Write(bytes []byte) (int, error) {
	message := []byte(NewTimeStyle())
	message = append(message, bytes...)
	return c.File.Write(message)
}

func saveDebugMessage(timeStyle, debugStyle string) {
	dir := currentDir()
	loggerFileName := getFileName()
	file := createLogFile(dir, loggerFileName)
	defer file.Close()
	_, err := file.WriteString(timeStyle + debugStyle + "\n")
	if err != nil {
		panic(err)
	}
}

// carefull with the permision if we got permission denied will loop forever
func createLogFile(filePath, fileName string) *os.File {
	f, err := os.OpenFile(filePath+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		// error file is not created
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		createLogFile(filePath, fileName)
	}

	return f
}

// return absolute path in the current directory
func currentDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(ex)

	return dir + "/"
}

// Get filename for file logger name
func getFileName() string {
	return "logger.log"
}
