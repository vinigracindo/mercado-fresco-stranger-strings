package seller

import "fmt"

var listSeller []Seller

type Repository interface {
	GetAll() ([]Seller, error)
	Get(id int64) (Seller, error)
	CreateSeller(cid int64, companyName, address, telephone string) (Seller, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Seller, error) {
	listSeller = []Seller{}
	return listSeller, nil
}

func (r *repository) Get(id int64) (Seller, error) {
	for _, seller := range listSeller {
		if seller.Id == id {
			return seller, nil
		}
	}

	return Seller{}, fmt.Errorf("seller id %d not found", id)
}

func (r *repository) CreateSeller(cid int64, companyName, address, telephone string) (Seller, error) {
	nextId := seller[len(seller)-1].Id + 1
	seller := Seller{
		Id:          nextId,
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
	}

	if !r.CidIsUnique(Cid) {
		return Seller{}, fmt.Errorf("cid %s is alredy in use", cid)
	}

	listSeller = append(listSeller, seller)
	return seller, nil

}
