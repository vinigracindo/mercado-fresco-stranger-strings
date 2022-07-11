package repository

var (
	SQLCreateInboundOrder = `
	INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id, warehouse_id)
	VALUES (?, ?, ?, ?, ?)`
)
