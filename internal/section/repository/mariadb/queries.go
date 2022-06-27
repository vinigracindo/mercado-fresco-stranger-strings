package repository

const (
	sqlGetAllSection = "SELECT  * FROM sections"

	sqlGetByIdSection = "SELECT * FROM sections WHERE id=?"

	sqlCreateSection = `
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
	sqlUpdateCurrentCapacitySection = `
    UPDATE sections
    SET current_capacity=?
    WHERE id=?
    `
	sqlDeleteSection = "DELETE FROM sections WHERE id=?"
)
