package deeper

import (
	"fmt"
)

func hi(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}
func Channels() {
	done := make(chan bool)
	go hi(done)
	var text string
	if <-done {
		text = "done"
	} else {
		text = "dont done"
	}

	fmt.Println(text)
}
