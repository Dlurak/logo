package src

func IsValidLogLevel(level LogLevel) bool {
	switch level  {
	case DebugLevel, InfoLevel, WarningLevel, ErrorLevel, FatalLevel:
		return true
	default:
		return false
	}
}
