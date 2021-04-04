package controllers

import (
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
	return c.RenderText("Sign up")
}
