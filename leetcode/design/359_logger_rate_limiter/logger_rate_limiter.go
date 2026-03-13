package logger_rate_limiter

type Logger struct{}

func Constructor() Logger                                               { return Logger{} }
func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool { return false }
