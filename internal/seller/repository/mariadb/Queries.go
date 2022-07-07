package repository

const (
	SqlGetAllSeller = `
	SELECT id, cid, company_name, address, telephone 
	FROM sellers`

	SqlGetByIdSeller = `
	SELECT id, cid, company_name, address, telephone
	FROM sellers
	WHERE id = ?`

	SqlCreateSeller = `
	INSERT INTO sellers (cid, company_name, address, telephone)
	VALUES (?, ?, ?, ?)`

	SqlUpdateSeller = `
	UPDATE sellers
	SET address = ?, telephone = ?
	WHERE id = ?`

	SqlDeleteSeller = `
	DELETE FROM sellers
	WHERE id = ?`
)
