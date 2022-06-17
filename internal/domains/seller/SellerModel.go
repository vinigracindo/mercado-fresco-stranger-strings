package seller

type Seller struct {
	Id          int64  `json:"id"`
	Cid         int64  `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type Service interface {
	GetAll() ([]Seller, error)
	GetById(id int64) (Seller, error)
	Create(cid int64, companyName, address, telephone string) (Seller, error)
	Update(id int64, address, telephone string) (Seller, error)
	Delete(id int64) error
}

type Repository interface {
	GetAll() ([]Seller, error)
	GetById(id int64) (Seller, error)
	Create(cid int64, companyName, address, telephone string) (Seller, error)
	Update(id int64, address, telephone string) (Seller, error)
	creatID() int64
	Delete(id int64) error
}
