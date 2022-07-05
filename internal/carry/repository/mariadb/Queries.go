package respository

const (
	QueryCreateCarry = `
    INSERT INTO
    carrier (cid, company_name, address, telephone, locality_id)
    VALUES (?, ?, ?, ?, ?)
    `

	QueryGetCarry = "select id, cid, company_name, address, telephone, locality_id from carrier"

	QueryCountLocality = "select count(*) as total_locality from carrier where locality_id = ?"
)
