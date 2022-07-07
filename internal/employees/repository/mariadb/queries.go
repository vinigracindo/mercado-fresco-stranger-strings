package repository

var (
	SQLFindAllEmployees = `
	SELECT id, card_number_id, first_name, last_name, warehouse_id 
	FROM employees`

	SQLFindEmployeeByID = `
	SELECT id, card_number_id, first_name, last_name, warehouse_id
	FROM employees
	WHERE id=?`

	SQLCreateEmployee = `
	INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id)
	VALUES (?, ?, ?, ?)`

	SQLUpdateEmployeeFullname = `
	UPDATE employees
	SET first_name = ?, last_name = ?
	WHERE id = ?`

	SQLDeleteEmployee = `
	DELETE FROM employees
	WHERE id = ?`

	SQLReportInboundOrders = `
	SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id, COUNT(i.employee_id) as inbound_orders_count
	FROM employees as e
	LEFT JOIN inbound_orders as i ON e.id = i.employee_id
	WHERE e.id = ifnull(?, e.id)
	GROUP BY e.id`
)
