package repository

const (
	SQLCreate = `
    INSERT INTO
    sections (
        batch_number,
        current_quantity,
        current_temperature,
        due_date,
        initial_quantity,
        manufacturing_date,
        manufacturing_hour,
        minumum_temperature,
		product_id,
		section_id
    )
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
)
