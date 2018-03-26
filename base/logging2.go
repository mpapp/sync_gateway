package base

// Errorf logs the given formatted string and args to the error log level with an optional log key.
func Errorf(logKey uint32, format string, args ...interface{}) {
	logTo(LevelError, logKey, format, args...)
}

// Warnf logs the given formatted string and args to the warn log level with an optional log key.
func Warnf(logKey uint32, format string, args ...interface{}) {
	logTo(LevelWarn, logKey, format, args...)
}

// Infof logs the given formatted string and args to the info log level with an optional log key.
func Infof(logKey uint32, format string, args ...interface{}) {
	logTo(LevelWarn, logKey, format, args...)
}

// Debugf logs the given formatted string and args to the debug log level with an optional log key.
func Debugf(logKey uint32, format string, args ...interface{}) {
	logTo(LevelWarn, logKey, format, args...)
}

func logTo(level uint32, logKey uint32, format string, args ...interface{}) {
	// FIXME: Stub
	return
}
