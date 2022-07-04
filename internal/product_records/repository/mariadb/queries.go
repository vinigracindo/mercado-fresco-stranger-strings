package mariadb

const (
	SqlCreate = `
    INSERT INTO
    product_records (
	last_update_date,
	purchase_price,
	sale_price,
	product_id
	)
    VALUES (?, ?, ?, ?)
    `
)
