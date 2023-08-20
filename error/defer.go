package error

import (
	"io"
	"os"
)

/*A defer statement pushes a function call onto a list.
The list of saved calls is executed after the surrounding function returns.
Defer is commonly used to simplify functions that perform various clean-up actions.*/

func Defer(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}
