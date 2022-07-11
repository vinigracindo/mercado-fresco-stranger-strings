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

	SQLGetAllPurchaseOrdersReports = `
	SELECT p.id, p.card_number_id, p.first_name, p.last_name, count(pr.id) as records_count
	FROM buyers as p
	LEFT JOIN purchase_orders as pr on p.id = pr.buyer_id
	GROUP BY p.id, p.card_number_id, p.first_name, p.last_name
	`
)
