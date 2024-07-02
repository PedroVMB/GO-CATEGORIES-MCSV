package use_cases

import (
	"errors"

	c_models "github.com/PedroVMB/GO-CATEGORIES-MCSV/internal/categories/models"
	c_repository "github.com/PedroVMB/GO-CATEGORIES-MCSV/internal/categories/repository"
)

type CreateCategoryUseCase struct {
	repository *c_repository.CategoryRepository
}

func NewCreateCategoryUseCase(repository *c_repository.CategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{repository}
}

func (u *CreateCategoryUseCase) Execute(name string) error {
	category, err := u.repository.FindByName(name)

	if err != nil {
		return err
	}

	if category != nil {
		return errors.New("name already in use")
	}

	category, err = c_models.NewCategory(name)

	if err != nil {
		return err
	}

	err = u.repository.Create(category)

	if err != nil {
		return err
	}

	return nil

}
