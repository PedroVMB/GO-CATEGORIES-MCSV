package use_cases

import (
	c_models "github.com/PedroVMB/GO-CATEGORIES-MCSV/internal/categories/models"
	c_repository "github.com/PedroVMB/GO-CATEGORIES-MCSV/internal/categories/repository"
)

type GetCategoriesUseCase struct {
	repository *c_repository.CategoryRepository
}

func NewGetCategoriesUseCase(repository *c_repository.CategoryRepository) *GetCategoriesUseCase {
	return &GetCategoriesUseCase{repository}
}

// TODO: pagination
func (u *GetCategoriesUseCase) Execute() ([]*c_models.Category, error) {
	categories, err := u.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return categories, nil
}
