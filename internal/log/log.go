package log

var (
	log Logger
)

type Logger interface {
	Infow(format string, args ...interface{})
	Errorw(format string, args ...interface{})

	Sync() error
}

func NewLogger() Logger {
	// This method helps us to easily change the log implementation anytime we wanted to.
	return newZapLogger()
}

func init() {
	// We initialize a global "log" variable to make logging more easily. It's best to inject dependencies, but
	// I did this as an exception.
	log = NewLogger()
}

// Infow is a proxy method which calls the underlying method on the logger instance.
func Infow(format string, args ...interface{}) {
	log.Infow(format, args...)
}

// Errorw is a proxy method which calls the underlying method on the logger instance.
func Errorw(format string, args ...interface{}) {
	log.Infow(format, args...)
}

// Sync is a proxy method which calls the underlying method on the logger instance.
func Sync() error {
	return log.Sync()
}
