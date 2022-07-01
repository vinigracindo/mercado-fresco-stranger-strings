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
	result, err := m.db.ExecContext(ctx, SQLDeleteSection, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("section not found")
	}

	return nil
}

func (m *mariaDbSectionRepository) UpdateCurrentCapacity(ctx context.Context, section *domain.SectionModel) (*domain.SectionModel, error) {
	result, err := m.db.ExecContext(ctx, SQLUpdateCurrentCapacitySection, &section.CurrentCapacity, &section.Id)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("section not found")
	}

	return section, nil
}

func (m *mariaDbSectionRepository) Create(ctx context.Context, sectionNumber int64, currentTemperature float64, minimumTemperature float64, currentCapacity int64, minimumCapacity int64, maximumCapacity int64, warehouseId int64, productTypeId int64) (domain.SectionModel, error) {
	section, err := m.db.ExecContext(
		ctx,
		SQLCreateSection,
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

	newSectionId, _ := section.LastInsertId()

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

	row := m.db.QueryRowContext(ctx, SQLGetByIdSection, id)

	var section domain.SectionModel

	err := row.Scan(
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

	if errors.Is(err, sql.ErrNoRows) {
		return section, errors.New("section not found")
	}

	if err != nil {
		return section, err
	}

	return section, nil
}

func (m *mariaDbSectionRepository) GetAll(ctx context.Context) ([]domain.SectionModel, error) {
	sections := []domain.SectionModel{}

	rows, err := m.db.QueryContext(ctx, SQLGetAllSection)
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
