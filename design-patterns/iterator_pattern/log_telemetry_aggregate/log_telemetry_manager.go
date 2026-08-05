package log_telemetry_aggregate

type TelemetryLogs struct {
	name     string
	iterator LogTelemetryIterator
}

func (tl *TelemetryLogs) PrintLogTelemetry() {
	for tl.iterator.HasNext() {
		item := tl.iterator.Next()
		println("Telemetry Log:", tl.name, item.GetData())
	}
}

func NewGrafanaTelemetryLogs() *TelemetryLogs {
	return &TelemetryLogs{
		name:     "Grafana",
		iterator: NewGrafanaLogsIterator(),
	}
}

func NewDBTelemetryLogs() *TelemetryLogs {
	return &TelemetryLogs{
		name:     "DB",
		iterator: NewDBLogsIterator(),
	}
}
