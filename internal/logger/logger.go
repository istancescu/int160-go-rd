package logger

import "log"

func LogInfo(msg string) {
	log.Printf("INFO: %s", msg)
}

func LogError(msg string) {
	log.Printf("ERROR: %s", msg)
}

func LogFatal(msg string) {
	log.Printf("FATAL: %s", msg)
}
