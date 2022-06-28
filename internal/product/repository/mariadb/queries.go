package repository

const (
	sqlGetAll = "SELECT  * FROM products"

	sqlGetById = "SELECT * FROM products WHERE id=?"

	sqlCreate = `
    INSERT INTO
    products (product_code, description, width, height, length, net_weight, 
	expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	sqlUpdateDescription = `
    UPDATE products
    SET description=?
    WHERE id=?
    `
	sqlDelete = "DELETE FROM products WHERE id=?"
)
