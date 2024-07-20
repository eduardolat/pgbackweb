package logger

import "os"

// logWriter is a simple io.Writer that can redirect logs to os.Stdout
// or any other required place
type logWriter struct{}

// Write writes the log message to os.Stdout or any other required place
func (w *logWriter) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}
