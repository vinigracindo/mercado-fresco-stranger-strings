package seller

var listSeller []Seller

type Repository interface {
	GetAll() ([]Seller, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Seller, error) {
	listSeller = []Seller{}
	return listSeller, nil
}
