package controllers

import (
	"go-revel-web-api/app/models"
	"net/http"

	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	return c.RenderText("/users")
}

func (c User) Login() revel.Result {
	return c.RenderText("Login")
}

func (c User) SignUp() revel.Result {

	var registerUser models.UserModel

	err := c.Params.BindJSON(&registerUser)
	if err != nil {
		c.Response.SetStatus(http.StatusBadRequest)
		return c.RenderText("Error")
	}

	return c.RenderJSON(registerUser)
}
