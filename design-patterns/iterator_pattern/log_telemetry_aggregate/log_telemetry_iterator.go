package log_telemetry_aggregate

import (
	"fmt"
	"sort"
	"time"
)

type LogTelemetryItem interface {
	GetData() string
}

type LogTelemetryIterator interface {
	Next() LogTelemetryItem
	HasNext() bool
}

type GrafanaLogsItem struct {
	timeData string
	data     string
}

func (g *GrafanaLogsItem) GetData() string {
	return fmt.Sprintf("%s: %s", g.timeData, g.data)
}

type GrafanaLogsIterator struct {
	logs            map[string]string
	logKeys         []string
	currentPosition int
}

func NewGrafanaLogsIterator() *GrafanaLogsIterator {
	logs := make(map[string]string)
	logs[time.Now().Format("2006-01-02 15:04:05")] = "Grafana Log 1"
	logs[time.Now().Add(time.Minute).Format("2006-01-02 15:04:05")] = "Grafana Log 2"
	logs[time.Now().Add(2*time.Minute).Format("2006-01-02 15:04:05")] = "Grafana Log 3"
	logs[time.Now().Add(3*time.Minute).Format("2006-01-02 15:04:05")] = "Grafana Log 4"

	keys := make([]string, 0, len(logs))
	for k := range logs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return &GrafanaLogsIterator{
		logs:    logs,
		logKeys: keys,
	}
}

func (g *GrafanaLogsIterator) Next() LogTelemetryItem {
	if g.HasNext() {
		item := g.logs[g.logKeys[g.currentPosition]]
		g.currentPosition++
		return &GrafanaLogsItem{
			timeData: g.logKeys[g.currentPosition-1],
			data:     item,
		}
	}
	return nil
}

func (g *GrafanaLogsIterator) HasNext() bool {
	return g.currentPosition < len(g.logs)
}

type DBLogsItem string

func (d *DBLogsItem) GetData() string {
	return string(*d)
}

type DBLogsIterator struct {
	logs            []DBLogsItem
	currentPosition int
}

func NewDBLogsIterator() *DBLogsIterator {
	logs := make([]DBLogsItem, 0)
	logs = append(logs, "DB Log 1")
	logs = append(logs, "DB Log 2")
	logs = append(logs, "DB Log 3")
	logs = append(logs, "DB Log 4")
	return &DBLogsIterator{
		logs: logs,
	}
}

func (d *DBLogsIterator) Next() LogTelemetryItem {
	if d.HasNext() {
		item := d.logs[d.currentPosition]
		d.currentPosition++
		return &item
	}
	return nil
}

func (d *DBLogsIterator) HasNext() bool {
	return d.currentPosition < len(d.logs)
}
