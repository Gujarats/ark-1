package logger

import "fmt"

// creating empty struct and the method for satisfaying io.Writer interface
type customWriter struct{}

// this is for customize the output format time
func (writer *customWriter) Write(bytes []byte) (int, error) {
	messageFormatted := NewTimeStyle() + string(bytes)
	return fmt.Print(messageFormatted)
}
