package logger

import (
	"strings"
	"time"
)

const timeFormat = "2006-02-01::3:04PM"

type TimeStyle struct {
	YearColor Attribute
	SepColor  Attribute // separator ::
	DateColor Attribute
	Style     Attribute
}

func (t *TimeStyle) getTimeColorFormat(time string) string {
	timeSplit := strings.Split(time, "::")
	year := GetColorFormat(t.YearColor, t.Style, timeSplit[0])
	date := GetColorFormat(t.DateColor, t.Style, timeSplit[1])
	sep := GetColorFormat(t.SepColor, t.Style, "::")

	result := strings.Join([]string{year, date}, sep) + sep
	return result

}

func getTimeNowFormat() string {
	return time.Now().UTC().Format(timeFormat)
}

func NewTimeStyle() string {
	time := getTimeNowFormat()
	var timeStyle TimeStyle

	if EnableCustomColor {
		timeStyle = CustomColorStyle.TimeStyle
	} else {
		//default color
		timeStyle = TimeStyle{
			YearColor: White,
			SepColor:  Yellow,
			DateColor: Cyan,
			Style:     Faint,
		}
	}

	return timeStyle.getTimeColorFormat(time)
}
