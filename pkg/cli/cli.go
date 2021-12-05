package cli

import (
	"io"
	"log"
	"os"
)

// ExitOnError logs a fatal and exits if the passed error is non-nil.
func ExitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ExitOnErrorf logs a fatal if the passed error is non-nil. The format string
// and args are passed directly to log.Fatalf, so callers must account for
// including the error's string representation if they desire.
func ExitOnErrorf(err error, msg string, args ...interface{}) {
	if err != nil {
		log.Fatalf(msg, args...)
	}
}

// GetInput returns the contents of the file at path, or os.Stdin if path is the
// empty string.
func GetInput(path string) ([]byte, error) {
	f, err := Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}

// Open returns a ReadCloser handle to the file at path, or os.Stdin if path is
// the empty string.
func Open(path string) (io.ReadCloser, error) {
	if path == "" {
		return os.Stdin, nil
	}

	return os.Open(path)
}
