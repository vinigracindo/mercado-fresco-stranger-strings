package repository

const (
	SQLCreatePurchaseOrders = `
    INSERT INTO
    purchase_orders (
        order_number,
        order_date,
        tracking_code,
        buyer_id,
        product_record_id,
        order_status_id
    )
    VALUES (?, ?, ?, ?, ?, ?)
`

	SQLContByBuyerId = `
	SELECT count(pr.id) as records_count
	FROM buyers as p
	LEFT JOIN purchase_orders as pr on p.id = pr.buyer_id
    WHERE p.id = ?
	`
)
