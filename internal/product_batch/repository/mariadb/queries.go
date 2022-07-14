package repository

const (
	SQLCreate = `
    INSERT INTO product_batches (
        batch_number, 
        current_quantity, 
        current_temperature, 
        due_date, 
        initial_quantity, 
        manufacturing_date, 
        manufacturing_hour, 
        minimum_temperature, 
        product_id, 
        section_id
    )
    VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

    `
)
