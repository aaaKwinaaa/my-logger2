package mylogger

import "log"

func LogInfo(message string) {
	log.Printf("INFO - %v", message)
}

func LogWarning(message string) {
	log.Printf("WARN - %v", message)
}

func LogError(message string) {
	log.Printf("ERROR - %v", message)
}

func LogCustom(message string) {
	log.Printf("Custom - %v", message)
}


