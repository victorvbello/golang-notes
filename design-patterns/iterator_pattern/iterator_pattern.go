package main

import "github.com/victorvbello/gonotes/design-patterns/iterator_pattern/log_telemetry_aggregate"

func main() {
	grafanaLogs := log_telemetry_aggregate.NewGrafanaTelemetryLogs()
	grafanaLogs.PrintLogTelemetry()

	dbLogs := log_telemetry_aggregate.NewDBTelemetryLogs()
	dbLogs.PrintLogTelemetry()
}
