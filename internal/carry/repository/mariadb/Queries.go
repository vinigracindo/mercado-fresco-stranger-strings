package respository

const (
	QueryCreateCarry = `
    INSERT INTO
    carriers (cid, company_name, address, telephone, locality_id)
    VALUES (?, ?, ?, ?, ?)
    `

	QueryGetCarry = "select id, cid, company_name, address, telephone, locality_id from carriers"

	QueryCountLocality = "select count(*) as total_locality from carriers where locality_id = ?"
)
