package product

type Service interface {
	GetAll() ([]Product, error)
	GetById(id int64) (*Product, error)
	Create(productCode string, description string, width float64, height float64, length float64, netWeight float64,
		expirationRate float64, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (Product, error)
}

type service struct {
	repository Repository
}

func (s service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s service) GetById(id int64) (*Product, error) {
	product, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s service) Create(productCode string, description string, width float64, height float64, length float64, netWeight float64,
	expirationRate float64, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (Product, error) {

	lastId := s.repository.LastId()
	lastId++
	newProduct, err := s.repository.Create(lastId, productCode, description, width, height, length,
		netWeight, expirationRate, recommendedFreezingTemperature, freezingRate, productTypeId, sellerId)

	if err != nil {
		return Product{}, err
	}

	return newProduct, nil
}

func CreateService(r Repository) Service {
	return &service{repository: r}
}
