package repository

const (
	sqlGetAllBuyer = "SELECT * FROM buyers"

	sqlGetByIdBuyer = "SELECT * FROM buyers WHERE id=?"

	sqlCreateBuyer = `
    INSERT INTO
    buyer (
        card_number_id,
        first_name,
        last_name
    )
    VALUES (?, ?, ?)
    `

	sqlUpdateAwardBuyer = `
    UPDATE buyers
    SET card_number_id=?, last_name=?
    WHERE id=?
    `

	sqlDeleteBuyer = "DELETE FROM buyers WHERE id=?"
)
