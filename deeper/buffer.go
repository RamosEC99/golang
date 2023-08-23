package deeper

import (
	"bytes"
	"fmt"
)

// importing the bytes package so that buffer can be used
func Buffer1() {
	//Creating buffer variable to hold and manage the string data
	var strBuffer bytes.Buffer
	strBuffer.WriteString("Ranjan")
	strBuffer.WriteString("Kumar")
	fmt.Println("The string buffer output is", strBuffer.String())
}

func Buffer2() {
	var strBuffer bytes.Buffer
	fmt.Fprintf(&strBuffer, "It is number: %d, This is a string: %v\n", 10, "Bridge")

	strBuffer.WriteString("[DONE]")
	fmt.Println("The string buffer output is", strBuffer.String())
}
