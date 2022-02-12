package handler

import (
	"net/http"
	"ticketmanagement/helper"
	"ticketmanagement/user"

	"github.com/gin-gonic/gin"
)

//handler tidak membutuhkan interface karena tidak menjadi depedency dari siapapun
type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// mapping  input dari user ke struct RegisterUserInput
	// struct diatas kita passing datanya sebagai parameter service
	//var input yang akan di mapping ke strcut register user input
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input) //untuk mengubah input di JSON menjadi RegisterUserInput (Struct di Input)
	//input siap di masukkan kedalam service
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{"errors": errors}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "Check Again your Request", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	//memasukkan input kedalam service
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, "Connection Timeout", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//agar kalau terjadi error tidak akan melanjutkan ke selanjutnya
	formatter := user.FormatUser(newUser, "tokentokentoken")
	response := helper.APIResponse(http.StatusOK, "Account Has Been Created", formatter)

	c.JSON(http.StatusOK, response)
}
