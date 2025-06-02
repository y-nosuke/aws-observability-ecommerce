package handler

import (
	"context"
	"database/sql"
)

type DatabaseHealthChecker struct {
	DB *sql.DB
}

func (d *DatabaseHealthChecker) Name() string { return "database" }
func (d *DatabaseHealthChecker) Check(ctx context.Context) error {
	if d.DB == nil {
		return &HealthCheckError{Msg: "not initialized"}
	}

	if err := d.DB.PingContext(ctx); err != nil {
		return &HealthCheckError{Msg: "connection failed", OriginalErr: err}
	}

	return nil
}
