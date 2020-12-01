package cli

import "log"

func ExitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
