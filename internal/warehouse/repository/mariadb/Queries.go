package respository

const (
	GetAllWarehouses = "SELECT id, address, telephone, warehouse_code, minimun_capacity, minimun_temperature, locality_id FROM warehouses"

	GetWarehouseById = "SELECT id, address, telephone, warehouse_code, minimun_capacity, minimun_temperature, locality_id FROM warehouses WHERE id=?"

	CreateWarehouse = `
    INSERT INTO
    warehouses (address, telephone, warehouse_code, minimun_capacity, minimun_temperature, locality_id)
    VALUES (?, ?, ?, ?, ?, ?)
    `

	UpdateWarehouse = `
    UPDATE warehouses
    SET minimun_capacity = ?, minimun_temperature  = ?
    WHERE id=?
    `
	DeleteWarehouse = "DELETE FROM warehouses WHERE id=?"
)
