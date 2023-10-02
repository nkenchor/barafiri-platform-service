package services

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"barafiri-platform-service/internal/core/helper"
	port "barafiri-platform-service/internal/ports"
	"github.com/google/uuid"
)

type categoryCategories struct {
	categoryRepository port.CategoryRepository
	industryRepository port.IndustryRepository
}

func NewCategory(
	categoryRepository port.CategoryRepository,
) *categoryCategories {
	return &categoryCategories{

		categoryRepository: categoryRepository,
	}
}
func (category *categoryCategories) CreateCategory(cat entity.Category) (interface{}, error) {
	cat.Reference = uuid.New().String()
	helper.LogEvent("INFO", "Creating category configuration with reference: "+cat.Reference)

	if err := category.validateIndustryInfo(cat.Industry); err != nil {
		return nil, err
	}

	if err := helper.Validate(cat); err != nil {
		return nil, err
	}
	return category.categoryRepository.CreateCategory(cat)
}

func (category *categoryCategories) UpdateCategory(reference string, cat entity.Category) (interface{}, error) {
	helper.LogEvent("INFO", "Updating category configuration with reference: "+reference)
	_, err := category.GetCategoryByRef(reference)
	cat.Reference = reference
	if err != nil {
		return nil, err
	}
	if err := category.validateIndustryInfo(cat.Industry); err != nil {
		return nil, err
	}

	if err := helper.Validate(cat); err != nil {
		return nil, err
	}
	return category.categoryRepository.UpdateCategory(reference, cat)
}
func (category *categoryCategories) EnableCategory(reference string, enabled bool) (interface{}, error) {
	helper.LogEvent("INFO", "Enabling category configuration with reference: "+reference)
	_, err := category.GetCategoryByRef(reference)
	if err != nil {
		return nil, err
	}
	return category.categoryRepository.EnableCategory(reference, enabled)
}

func (category *categoryCategories) GetCategoryByRef(reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting category configuration with reference: "+reference)
	cat, err := category.categoryRepository.GetCategoryByRef(reference)
	if err != nil {
		return nil, err
	}
	return cat, nil
}
func (category *categoryCategories) GetCategoryByName(name string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting category configuration with name: "+name)
	cat, err := category.categoryRepository.GetCategoryByName(name)
	if err != nil {
		return nil, err
	}
	return cat, nil
}

func (category *categoryCategories) GetAllCategories(page string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting all categories...")
	cat, err := category.categoryRepository.GetAllCategories(page)

	if err != nil {
		return nil, err
	}
	return cat, nil
}

func (category *categoryCategories) validateIndustryInfo(industryInfo entity.IndustryInfo) error {
	helper.LogEvent("INFO", "Validating industry with reference: "+industryInfo.Reference)
	industry, err := category.industryRepository.GetIndustryByRef(industryInfo.Reference)
	if err != nil {
		return err
	}
	industryName := industry.(entity.IndustryInfo).Reference
	if industryName != industryInfo.Name {
		return helper.ErrorMessage(helper.ValidationError, "Sorry, Industry name is invalid")
	}
	return nil
}
