package cli

import (
	"io/ioutil"
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

func GetInput(path string) ([]byte, error) {
	if path == "" {
		return ioutil.ReadAll(os.Stdin)
	}

	return ioutil.ReadFile(path)
}
