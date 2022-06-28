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
)
