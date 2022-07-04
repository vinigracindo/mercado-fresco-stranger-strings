package repository

const (
	SQLGetAllBuyer = `
	SELECT id, card_number_id, first_name, last_name
	FROM buyers`

	SQLGetByIdBuyer = `
	SELECT id, card_number_id, first_name, last_name
	FROM buyers
	WHERE id=?`

	SQLCreateBuyer = `
    INSERT INTO
    buyers (
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
