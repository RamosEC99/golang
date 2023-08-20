package error

import (
	"log"
	"os"
)

func HandlingError(name string) (file *os.File, err error) {
	f, err := os.Open("filename.ext")
	if err != nil {
		log.Fatal(err)
	}
	return f, err
}
