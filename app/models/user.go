package models

import (
	"github.com/kamva/mgm/v3"
)

type UserModel struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
}
