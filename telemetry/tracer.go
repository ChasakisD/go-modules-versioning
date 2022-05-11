package telemetry

import "fmt"

func TraceV2(message string) {
	fmt.Printf("Log: %s", message)
}
