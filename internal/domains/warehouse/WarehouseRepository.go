package warehouse

var wh []WarehouseModel
var id int64 = 0

type warehouse struct{}

type WarehouseInterface interface {
	Create(wr *WarehouseModel) (WarehouseModel, error)
	createID() int64
}

func (w warehouse) createID() int64 {
	id += id + 1
	return id
}

func (w warehouse) Create(wr *WarehouseModel) (WarehouseModel, error) {
	wr.Id = w.createID()
	wh = append(wh, *wr)

	return *wr, nil
}

func newRepository() WarehouseInterface {
	return &warehouse{}
}
