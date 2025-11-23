package controllers

import "github.com/gin-gonic/gin"

type FormControllerImpl interface {
	CreateFormDetails(c *gin.Context)
	GetAllFormDetailsByEmail(c *gin.Context)
}
