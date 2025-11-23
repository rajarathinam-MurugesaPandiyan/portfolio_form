package controllers

import (
	"fmt"
	"net/http"
	"portfolio_form/models"
	"portfolio_form/services"

	"github.com/gin-gonic/gin"
)

type FormController struct {
	services *services.FormService
}

func InitializeFormController(services *services.FormService) *FormController {
	return &FormController{
		services: services,
	}
}

func (f *FormController) CreateFormDetails(c *gin.Context) {
	var input models.FormInputs

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := f.services.CreateFormDetails(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

func (f *FormController) GetAllFormDetailsByEmail(c *gin.Context) {

	querParams := c.Param("email")
	fmt.Print("Quer", querParams)

	response, err := f.services.GetAllFormDetailsByEmail(querParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})

}
