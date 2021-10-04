package metrics

import "time"

type RequestMetrics struct {
	Start      time.Time
	TotalAdded int
}
