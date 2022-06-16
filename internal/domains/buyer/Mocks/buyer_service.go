package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
)

type BuyerServiceMock struct {
	mock.Mock
}

func (b *BuyerServiceMock) GetAll() ([]buyer.Buyer, error) {
	args := b.Called()

	var buyers []buyer.Buyer

	if rf, ok := args.Get(0).(func() []buyer.Buyer); ok {
		buyers = rf()
	} else {
		if args.Get(0) != nil {
			buyers = args.Get(0).([]buyer.Buyer)
		}
	}

	var err error
	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return buyers, err
}

func (b *BuyerServiceMock) Store(cardNumberId int64, firstName string, lastName string) (buyer.Buyer, error) {
	args := b.Called()

	var buyers buyer.Buyer

	if rf, ok := args.Get(0).(func(cardNumberId int64, firstName string, lastName string) buyer.Buyer); ok {
		buyers = rf(cardNumberId, firstName, lastName)
	} else {
		if args.Get(0) != nil {
			buyers = args.Get(0).(buyer.Buyer)
		}
	}

	var err error
	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return buyers, err
}

func (b *BuyerServiceMock) GetId() (int64, error) {
	args := b.Called()

	var CreateId int64

	if rf, ok := args.Get(0).(func() int64); ok {
		CreateId = rf()
	} else {
		if args.Get(0) != nil {
			CreateId = args.Get(0).(int64)
		}
	}

	var err error
	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return CreateId, err
}

func (b *BuyerServiceMock) Update(id int64, cardNumberId int64, lastName string) (buyer.Buyer, error) {
	args := b.Called()

	var buyers buyer.Buyer

	if rf, ok := args.Get(0).(func(cardNumberId int64, lastName string) buyer.Buyer); ok {
		buyers = rf(cardNumberId, lastName)
	} else {
		if args.Get(0) != nil {
			buyers = args.Get(0).(buyer.Buyer)
		}
	}

	var err error
	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return buyers, err
}

func (b *BuyerServiceMock) Delete(id int64) error {
	args := b.Called()

	var err error

	if rf, ok := args.Get(0).(func(id int64) error); ok {
		err = rf(id)
	} else {
		err = args.Error(0)
	}
	return err
}
