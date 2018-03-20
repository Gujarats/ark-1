package logger

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

//Enable custome color
var EnableCustomColor bool
var CustomColorStyle CustomColor

// Attribute defines a single SGR Code
type Attribute int

const escape = "\x1b"

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Text colors
const (
	Black Attribute = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

type CustomColor struct {
	DebugStyle
	TimeStyle
}

func NewCustomColor(debugStyle DebugStyle, timeStyle TimeStyle) {
	if !checkAttribute(debugStyle) || !checkAttribute(timeStyle) {
		//error one of the Style attribute is empty
		err := errors.New("All style attribute must not be empty please check all style attribute")
		log.Fatal(err)
	}

	CustomColorStyle = CustomColor{
		DebugStyle: debugStyle,
		TimeStyle:  timeStyle,
	}

	EnableCustomColor = true

}

// return a formatted color based on the chosen Attribute for specific text
func GetColorFormat(color, style Attribute, text string) string {
	return fmt.Sprintf("%+v[%+v;%+vm%s%s[%+vm", escape, color, style, text, escape, Reset)
}

// passing struct and check all the attribute,
// false if there is empty value.
func checkAttribute(input interface{}) bool {
	object := reflect.ValueOf(input)

	for index := 0; index < object.NumField(); index++ {

		if isZeroOfUnderlyingType(object.Field(index).Interface()) {
			return false
		}
	}

	return true
}

// compare the object value to non assign value.
func isZeroOfUnderlyingType(objectValue interface{}) bool {
	return reflect.DeepEqual(objectValue, reflect.Zero(reflect.TypeOf(objectValue)).Interface())
}
