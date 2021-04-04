package controllers

import (
	"go-revel-web-api/app/models"
	"net/http"

	"github.com/kamva/mgm/v3"
	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := mgm.SetDefaultConfig(nil, "nextflow", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
}

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

	err = mgm.Coll(&registerUser).Create(&registerUser)
	if err != nil {
		c.Response.SetStatus(http.StatusInternalServerError)
		return c.RenderText("Internal Server  Error")
	}

	return c.RenderJSON(registerUser)
}
