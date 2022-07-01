package repository

const (
	SQLGetAllSection = `
    SELECT   
    id,       
    section_number,
    current_temperature,
    minimum_temperature,
    current_capacity,
    minimum_capacity,
    maximum_capacity,
    warehouse_id,
    product_type_id FROM section
    `

	SQLGetByIdSection = "SELECT * FROM sections WHERE id=?"

	SQLCreateSection = `
    INSERT INTO
    sections (
        section_number,
        current_temperature,
        minimum_temperature,
        current_capacity,
        minimum_capacity,
        maximum_capacity,
        warehouse_id,
        product_type_id
    )
    VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    `
	SQLUpdateCurrentCapacitySection = `
    UPDATE sections
    SET current_capacity=?
    WHERE id=?
    `
	SQLDeleteSection = "DELETE FROM sections WHERE id=?"
)
