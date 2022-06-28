package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/domain"
)

type mariaDbSectionRepository struct {
	db *sql.DB
}

func NewMariadbSectionRepository(db *sql.DB) domain.SectionRepository {
	return &mariaDbSectionRepository{db: db}
}

func (m *mariaDbSectionRepository) Delete(ctx context.Context, id int64) error {
	result, err := m.db.ExecContext(
		ctx,
		sqlDeleteSection,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("section not found")
	}

	return nil
}

func (m *mariaDbSectionRepository) UpdateCurrentCapacity(ctx context.Context, id int64, currentCapacity int64) (domain.SectionModel, error) {
	_, err := m.db.ExecContext(
		ctx,
		sqlUpdateCurrentCapacitySection,
		currentCapacity,
		id,
	)

	if err != nil {
		return domain.SectionModel{}, err
	}

	sectionUpdated, err := m.GetById(ctx, id)
	if err != nil {
		return domain.SectionModel{}, err
	}

	return sectionUpdated, nil
}

func (m *mariaDbSectionRepository) Create(ctx context.Context, sectionNumber string, currentTemperature float64, minimumTemperature float64, currentCapacity int64, minimumCapacity int64, maximumCapacity int64, warehouseId int64, productTypeId int64) (domain.SectionModel, error) {
	section, err := m.db.ExecContext(
		ctx,
		sqlCreateSection,
		sectionNumber,
		currentTemperature,
		minimumTemperature,
		currentCapacity,
		minimumCapacity,
		maximumCapacity,
		warehouseId,
		productTypeId,
	)

	if err != nil {
		return domain.SectionModel{}, err
	}

	newSectionId, err := section.LastInsertId()
	if err != nil {
		return domain.SectionModel{}, err
	}

	return domain.SectionModel{
		Id:                 newSectionId,
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseId:        warehouseId,
		ProductTypeId:      productTypeId,
	}, nil
}

func (m *mariaDbSectionRepository) GetById(ctx context.Context, id int64) (domain.SectionModel, error) {
	var section domain.SectionModel

	row, err := m.db.QueryContext(ctx, sqlGetByIdSection, id)

	if err != nil {
		return domain.SectionModel{}, err
	}

	if !row.Next() {
		return domain.SectionModel{}, errors.New("section not found")
	}

	err = row.Scan(
		&section.Id,
		&section.SectionNumber,
		&section.CurrentTemperature,
		&section.MinimumTemperature,
		&section.CurrentCapacity,
		&section.MinimumCapacity,
		&section.MaximumCapacity,
		&section.WarehouseId,
		&section.ProductTypeId,
	)
	if err != nil {
		return domain.SectionModel{}, err
	}

	defer row.Close()

	return section, nil
}

func (m *mariaDbSectionRepository) GetAll(ctx context.Context) ([]domain.SectionModel, error) {
	sections := []domain.SectionModel{}

	rows, err := m.db.QueryContext(ctx, sqlGetAllSection)
	if err != nil {
		return sections, err
	}

	defer rows.Close()

	for rows.Next() {
		var section domain.SectionModel

		err := rows.Scan(
			&section.Id,
			&section.SectionNumber,
			&section.CurrentTemperature,
			&section.MinimumTemperature,
			&section.CurrentCapacity,
			&section.MinimumCapacity,
			&section.MaximumCapacity,
			&section.WarehouseId,
			&section.ProductTypeId,
		)
		if err != nil {
			return sections, err
		}
		sections = append(sections, section)
	}
	return sections, nil
}
