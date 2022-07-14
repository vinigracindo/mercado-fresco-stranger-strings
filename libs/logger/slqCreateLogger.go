package logger

const (
	slqCreateLog = `
	INSERT INTO
	logs (
	method,
	label,
	level,
	message,
	status,
	insert_date
	)
    VALUES (?, ?, ?, ?, ?, ?)
    `
)
