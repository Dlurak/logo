package src

import "fmt"

type Color = string

const (
	Reset  Color = "\033[0m"
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Blue   Color = "\033[34m"
	Purple Color = "\033[35m"
	Cyan   Color = "\033[36m"
	White  Color = "\033[37m"
	Gray   Color = "\033[90m"
)

func fmtString(str string, color Color) string {
	return fmt.Sprintf("%s%s%s", color, str, Reset)
}

func fmtLogLevel(loglevel LogLevel, useColor bool) string {
	if !useColor {
		return string(loglevel)
	}

	switch loglevel {
	case DebugLevel:
		return fmtString(string(loglevel), Gray)
	case InfoLevel:
		return fmtString(string(loglevel), White)
	case WarningLevel:
		return fmtString(string(loglevel), Yellow)
	case ErrorLevel, FatalLevel:
		return fmtString(string(loglevel), Red)
	default:
		return fmtString(string(loglevel), Reset)
	}
}
