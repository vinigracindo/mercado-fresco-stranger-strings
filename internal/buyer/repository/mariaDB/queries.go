package repository

const (
	SQLGetAllBuyer = "SELECT * FROM buyers"

	SQLGetByIdBuyer = "SELECT * FROM buyers WHERE id=?"

	SQLCreateBuyer = `
    INSERT INTO
    buyer (
        card_number_id,
        first_name,
        last_name
    )
    VALUES (?, ?, ?)
    `

	SQLUpdateBuyer = `
    UPDATE buyers
    SET card_number_id=?, last_name=?
    WHERE id=?
    `

	SQLDeleteBuyer = "DELETE FROM buyers WHERE id=?"
)
