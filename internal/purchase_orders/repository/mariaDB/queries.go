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
)
