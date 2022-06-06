package seller

import "fmt"

var listSeller = []Seller{}
var id int64 = 0

type Repository interface {
	GetAll() ([]Seller, error)
	Get(id int64) (Seller, error)
	CreateSeller(cid int64, companyName, address, telephone string) (Seller, error)
	UpdateSellerAddresAndTel(id int64, address, telephone string) (Seller, error)
	creatID() int64
	DeleteSeller(id int64) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Seller, error) {
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

func (r *repository) creatID() int64 {
	id += 1
	return id
}

func (r *repository) CreateSeller(cid int64, companyName, address, telephone string) (Seller, error) {
	seller := Seller{
		Id:          r.creatID(),
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
	}

	for i := range listSeller {
		if listSeller[i].Cid == seller.Cid {
			return Seller{}, fmt.Errorf("Alredy a company with id %d", seller.Cid)
		}
	}

	listSeller = append(listSeller, seller)
	return seller, nil

}

func (r *repository) UpdateSellerAddresAndTel(id int64, address, telephone string) (Seller, error) {
	for i, seller := range listSeller {
		if seller.Id == id {
			listSeller[i].Address = address
			listSeller[i].Telephone = telephone
			return listSeller[i], nil
		}
	}
	return Seller{}, fmt.Errorf("seller with id %d not found", id)
}

func (r *repository) DeleteSeller(id int64) error {
	for i, seller := range listSeller {
		if seller.Id == id {
			listSeller = append(listSeller[:i], listSeller[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("seller with id %d", id)
}
