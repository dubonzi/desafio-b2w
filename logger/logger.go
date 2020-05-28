package logger

import "log"

// Error prints an error message to the log output.
func Error(context, method string, err error, args ...interface{}) {
	log.Printf("[ERROR] %s | %s: %s | Args: %v", context, method, err, args)
}

// Fatal prints an error message to the log output and stops the application.
func Fatal(context, method string, err error, args ...interface{}) {
	log.Fatalf("[FATAL] %s | %s: %s | Args: %v", context, method, err, args)
}
