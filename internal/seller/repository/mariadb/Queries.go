package repository

const (
	SqlGetAllSeller = `
	SELECT id, cid, company_name, address, telephone, locality_id
	FROM sellers`

	SqlGetByIdSeller = `
	SELECT id, cid, company_name, address, telephone, locality_id
	FROM sellers
	WHERE id = ?`

	SqlCreateSeller = `
	INSERT INTO sellers (cid, company_name, address, telephone, locality_id)
	VALUES (?, ?, ?, ?, ?)`

	SqlUpdateSeller = `
	UPDATE sellers
	SET address = ?, telephone = ?
	WHERE id = ?`

	SqlDeleteSeller = `
	DELETE FROM sellers
	WHERE id = ?`

	QueryCountByLocalityId = `
    SELECT COUNT(*) FROM sellers WHERE product_id = ?`
)
