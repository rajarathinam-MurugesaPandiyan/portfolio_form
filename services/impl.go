package services

import "portfolio_form/models"

type FormServiceImpl interface {
	CreateFormDetails(payload models.FormInputs) error
}
