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
        product_type_id 
    FROM sections
    `

	SQLGetByIdSection = `
    SELECT   
        id,       
        section_number,
        current_temperature,
        minimum_temperature,
        current_capacity,
        minimum_capacity,
        maximum_capacity,
        warehouse_id,
        product_type_id
    FROM sections WHERE id=?
    `

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

	SQLCountProductsBySectionWithSectionId = `
    SELECT s.id, s.section_number, 
    IFNULL(SUM(pb.current_quantity), 0) products_count  
    FROM sections s
    LEFT JOIN product_batches pb
    ON s.id = pb.section_id 
    WHERE s.id = ?
    GROUP BY s.id
    `

	SQLCountProductsBySection = `
    SELECT s.id, s.section_number, 
    IFNULL(SUM(pb.current_quantity), 0) products_count  
    FROM sections s
    LEFT JOIN product_batches pb
    ON s.id = pb.section_id 
    GROUP BY s.id
    `
)
