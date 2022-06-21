package seller

import "fmt"

var listSeller = []Seller{}
var id int64 = 0

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Seller, error) {
	return listSeller, nil
}

func (r *repository) GetById(id int64) (Seller, error) {
	for _, seller := range listSeller {
		if seller.Id == id {
			return seller, nil
		}
	}

	return Seller{}, fmt.Errorf("seller id %d not found", id)
}

func (r *repository) CreatID() int64 {
	id += 1
	return id
}

func (r *repository) Create(cid int64, companyName, address, telephone string) (Seller, error) {
	for i := range listSeller {
		if listSeller[i].Cid == cid {
			return Seller{}, fmt.Errorf("Alredy a company with id %d", cid)
		}
	}

	seller := Seller{
		Id:          r.CreatID(),
		Cid:         cid,
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
	}

	listSeller = append(listSeller, seller)
	return seller, nil

}

func (r *repository) Update(id int64, address, telephone string) (Seller, error) {
	for i, seller := range listSeller {
		if seller.Id == id {
			listSeller[i].Address = address
			listSeller[i].Telephone = telephone
			return listSeller[i], nil
		}
	}
	return Seller{}, fmt.Errorf("seller with id %d not found", id)
}

func (r *repository) Delete(id int64) error {
	for i, seller := range listSeller {
		if seller.Id == id {
			listSeller = append(listSeller[:i], listSeller[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("seller with id %d not found", id)
}
