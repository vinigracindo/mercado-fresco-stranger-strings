package respository

const (
	GetAllWarehouses = "SELECT  * FROM movie"
	GetWarehouseById = "SELECT * FROM movie WHERE id=?"
	CreateWarehouse  = `
    INSERT INTO
    warehouse (address, telephone, warehouse_code, minimun_capacity, minimun_temperature, locality_id)
    VALUES (?, ?, ?, ?, ?, ?)
    `
	UpdateWarehouse = `
    UPDATE movie
    SET award=?
    WHERE id=?
    `
	DeleteWarehouse = "DELETE FROM movie WHERE id=?"
)
