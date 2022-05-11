package telemetry

import "fmt"

func Trace(message string) {
	fmt.Printf("Log: %s", message)
}
