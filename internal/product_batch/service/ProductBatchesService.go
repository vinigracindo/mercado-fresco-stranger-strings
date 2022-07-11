package service

import (
	"context"

	product "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_batch/domain"
	section "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/domain"
)

type service struct {
	repository        domain.ProductBatchRepository
	repositoryProduct product.ProductRepository
	repositorySection section.SectionRepository
}

func NewProductBatchService(r domain.ProductBatchRepository, rp product.ProductRepository, rs section.SectionRepository) domain.ProductBatchService {
	return &service{
		repository:        r,
		repositoryProduct: rp,
		repositorySection: rs,
	}
}

func (s *service) Create(ctx context.Context, productBatch *domain.ProductBatch) (*domain.ProductBatch, error) {
	_, err := s.repositoryProduct.GetById(ctx, productBatch.Id)
	if err != nil {
		return nil, err
	}

	_, err = s.repositorySection.GetById(ctx, productBatch.Id)
	if err != nil {
		return nil, err
	}

	newProductBatch, err := s.repository.Create(ctx, productBatch)
	if err != nil {
		return nil, err
	}

	return newProductBatch, nil
}
