package helpers

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogMessage(message string, args ...interface{}) {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Print(err)
		}
	}(file)

	currentDateTime := time.Now().Format("2006-01-02 15:04:05")

	logEntry := fmt.Sprintf("[%s] %s  - : %v\n", currentDateTime, message, args)

	if _, err := file.WriteString(logEntry); err != nil {
		log.Fatal(err)
	}
}
