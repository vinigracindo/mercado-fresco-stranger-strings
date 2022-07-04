package respository

const (
	GetAllWarehouses = "SELECT id, address, telephone, warehouse_code, minimun_capacity, minimun_temperature, locality_id FROM warehouse"

	GetWarehouseById = "SELECT id, address, telephone, warehouse_code, minimun_capacity, minimun_temperature, locality_id FROM warehouse WHERE id=?"

	CreateWarehouse = `
    INSERT INTO
    warehouse (address, telephone, warehouse_code, minimun_capacity, minimun_temperature, locality_id)
    VALUES (?, ?, ?, ?, ?, ?)
    `

	UpdateWarehouse = `
    UPDATE warehouse
    SET minimun_capacity = ?, minimun_temperature  = ?
    WHERE id=?
    `
	DeleteWarehouse = "DELETE FROM warehouse WHERE id=?"
)
