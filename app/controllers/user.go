package controllers

import (
	"go-revel-web-api/app/models"
	"go-revel-web-api/app/utils"
	"net/http"

	"github.com/kamva/mgm/v3"
	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/bson"
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

	var loginUser models.UserModel

	err := c.Params.BindJSON(&loginUser)
	if err != nil {
		c.Response.SetStatus(http.StatusBadRequest)
		return c.RenderText("Error")
	}

	var existUsers []models.UserModel
	collection := mgm.Coll(&models.UserModel{})
	err = collection.SimpleFind(&existUsers, bson.M{"email": loginUser.Email, "password": loginUser.Password})

	if err != nil {
		c.Response.SetStatus(http.StatusInternalServerError)
		return c.RenderText("Database error")
	}

	if len(existUsers) == 0 {
		c.Response.SetStatus(http.StatusUnauthorized)
		return c.RenderText("login failed")
	}

	loggedInUser := existUsers[0]
	token := utils.EncodeToken(loggedInUser.Email)

	tokenModel := models.TokenModel{
		Token: token,
	}

	return c.RenderJSON(tokenModel)
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

func (c User) GetProfile() revel.Result {

	return c.RenderText("Getting in secret path")

}
