package logger

type DebugStyle struct {
	PrefixColor  Attribute
	MessageColor Attribute
	Style        Attribute
}

func (d *DebugStyle) getDebugColorFormat(prefix, message string) string {
	message = GetColorFormat(d.MessageColor, d.Style, message)
	prefix = GetColorFormat(d.PrefixColor, d.Style, prefix)

	return prefix + message
}

func NewDebugStyle(prefix, message string) string {
	var debug DebugStyle
	if EnableCustomColor {
		debug = CustomColorStyle.DebugStyle
	} else {
		//default color
		debug = DebugStyle{
			PrefixColor:  Yellow,
			MessageColor: Cyan,
			Style:        Faint,
		}
	}

	return debug.getDebugColorFormat(prefix, message)
}
