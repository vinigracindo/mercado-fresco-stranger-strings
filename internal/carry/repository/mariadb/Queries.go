package respository

const (
	QueryCreateCarry = `
    INSERT INTO
    carry (cid, company_name, address, telephone, locality_id)
    VALUES (?, ?, ?, ?, ?)
    `

	QueryGetCarry = "select id, cid, company_name, address, telephone, locality_id from carry"
)
