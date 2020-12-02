package cli

import (
	"io/ioutil"
	"log"
	"os"
)

func ExitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetInput(path string) ([]byte, error) {
	if path == "" {
		return ioutil.ReadAll(os.Stdin)
	}

	return ioutil.ReadFile(path)
}
