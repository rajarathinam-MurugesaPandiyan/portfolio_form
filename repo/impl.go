package repo

import "portfolio_form/models"

type FormRepoImpl interface {
	CreateFormDetails(payload models.FormInputs) error
}
