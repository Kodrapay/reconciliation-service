package models

import "time"

// ReconciliationRun represents a single reconciliation execution/request.
type ReconciliationRun struct {
	ID        int
	Source    string
	Date      string // keep as string for external-source friendly parsing
	Status    string
	Matched   int
	Unmatched int
	CreatedAt time.Time
}
