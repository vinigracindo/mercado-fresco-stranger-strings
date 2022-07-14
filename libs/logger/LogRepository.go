package logger

import (
	"context"
	"database/sql"
	"time"
)

type LogRepository interface {
	createLog(ctx context.Context, method, label, level, message string, status int)
	Error(ctx context.Context, method, label, message string, status int)
	Info(ctx context.Context, method, label, message string, status int)
}

type loggerRepoImpl struct {
	db *sql.DB
}

var Logger LogRepository

func InitializeLogger(db *sql.DB) {
	Logger = loggerRepoImpl{db: db}
}

func (l loggerRepoImpl) createLog(ctx context.Context, level, method, label, message string, status int) {

	insertDate := time.Now()

	_, err := l.db.ExecContext(
		ctx,
		slqCreateLog,
		method,
		label,
		level,
		message,
		status,
		insertDate,
	)

	if err != nil {
		print(err)
	}
}

func (l loggerRepoImpl) Error(ctx context.Context, method, label, message string, status int) {
	l.createLog(ctx, "ERROR", method, label, message, status)
}

func (l loggerRepoImpl) Info(ctx context.Context, method, label, message string, status int) {
	l.createLog(ctx, "INFO", method, label, message, status)
}

/*
	method,
	label,
	level,
	message,
	status,
	insert_date
*/
