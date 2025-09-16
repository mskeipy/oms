package logger

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
	"time"
)

type queryLogger struct {
	SlowThreshold time.Duration
}

// NewQueryLogger initializes a new query logger
func NewQueryLogger(slowThreshold time.Duration) logger.Interface {
	return &queryLogger{SlowThreshold: slowThreshold}
}

// LogMode is required by GORM but not used in this implementation
func (q *queryLogger) LogMode(level logger.LogLevel) logger.Interface {
	return q
}

// Info logs general messages
func (q *queryLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	log.Info().
		Str("request_id", getRequestID(ctx)).
		Msgf(msg, data...)
}

// Warn logs warnings
func (q *queryLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	log.Warn().
		Str("request_id", getRequestID(ctx)).
		Msgf(msg, data...)
}

// Error logs errors
func (q *queryLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	log.Error().
		Str("request_id", getRequestID(ctx)).
		Msgf(msg, data...)
}

// Trace logs SQL queries, execution time, and row count
func (q *queryLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	event := log.With().
		Str("request_id", getRequestID(ctx)).
		Str("sql", sql).
		Str("duration", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)).
		Int64("rows", rows).
		Logger()

	switch {
	case err != nil:
		event.Error().Err(err).Msg("SQL execution failed")
	case elapsed > q.SlowThreshold:
		event.Warn().Msg("Slow SQL detected")
	default:
		event.Info().Msg("SQL executed")
	}
}

func TransactionLogger(ctx context.Context, message string) {
	requestID := getRequestID(ctx)

	log.Info().
		Str("prefix", requestID).
		Msg(message)
}

const requestIdContext = "request_id"

// getRequestID extracts request ID from context (optional, customize as needed)
func getRequestID(ctx context.Context) string {
	if ctx == nil {
		return "unknown"
	}
	if rid, ok := ctx.Value(requestIdContext).(string); ok {
		return rid
	}
	return "unknown"
}
