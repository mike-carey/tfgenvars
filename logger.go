package tfgenvars

import (
	"fmt"
	"os"
)

func log(msg string) {
	os.Stderr.Write([]byte(fmt.Sprintf("[DEBUG] %s\n", msg)))
}

func logf(msg string, args... interface{}) {
	os.Stderr.Write([]byte(fmt.Sprintf(fmt.Sprintf("[DEBUG] %s\n", msg), args...)))
}
