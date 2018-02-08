package entity

import (
	"github.com/varunamachi/vaali/vparam"
)

type VariableAccess int

//Entity - represents a entiry with variables
type Entity struct {
	Name     string `json:"name" bson:"name"`
	Type     string `json:"type" bson:"type"`
	Location string `json:"location" bson:"location"`
	Owner    string `json:"owner" bson:"owner"`
}

//Variable - reprents a read-only or read-write variable
type Variable struct {
	Param vparam.Param
}
