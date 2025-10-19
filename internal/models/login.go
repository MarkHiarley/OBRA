package models

import "gopkg.in/guregu/null.v4"

type LoginUser struct {
	Email null.String `json:"email,omitempty"`
	Senha null.String `json:"senha,omitempty"`
}
