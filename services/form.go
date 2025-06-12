package services

import (
	"portfolio_form/models"
	"portfolio_form/repo"
)

type FormService struct {
	repos *repo.FormRepo
}

func InitializeFormService(repos *repo.FormRepo) *FormService {
	return &FormService{
		repos: repos,
	}
}

func (f *FormService) CreateFormDetails(payload models.FormInputs) error {
	err := f.repos.CreateFormDetails(payload)
	if err != nil {
		return err
	}
	return nil
}
