package utilities

import (
	"os"
	"time"
)

func WriteLog(filename, ip, msg string) {
	text := "[" + DateToString(time.Now()) + " " + ip + "] " + msg + "\n"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}
