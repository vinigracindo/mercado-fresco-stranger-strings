package mariadb

const (
	SqlGetAll = `
	SELECT 
	id,
	product_code,
	description,
	width,
	height,
	length,
	net_weight,
	expiration_rate,
	recommended_freezing_temperature,
	freezing_rate,
	product_type_id,
	seller_id FROM products
	`

	SqlGetById = "SELECT * FROM products WHERE id=?"

	SqlCreate = `
    INSERT INTO
    products (
	product_code,
	description,
	width,
	height,
	length,
	net_weight,
	expiration_rate,
	recommended_freezing_temperature,
	freezing_rate,
	product_type_id,
	seller_id
	)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	SqlUpdateDescription = `
    UPDATE products
    SET description=?
    WHERE id=?
    `
	SqlDelete = "DELETE FROM products WHERE id=?"

	SqlGetAllReportProductRecords = `
	SELECT p.id, p.description, count(pr.id)
	FROM products as p
	LEFT JOIN product_records as pr on p.id = pr.product_id
	GROUP BY p.id
	`
)
