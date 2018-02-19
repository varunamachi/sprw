package entity

import (
	"time"

	"github.com/varunamachi/vaali/vcmn"
)

//VariableAccess - access level of variable
type VariableAccess int

//Entity - represents a entiry with variables
type Entity struct {
	Name       string       `json:"name" bson:"name"`
	Type       string       `json:"type" bson:"type"`
	Location   string       `json:"location" bson:"location"`
	Owner      string       `json:"owner" bson:"owner"`
	Variables  []vcmn.Param `json:"variables" bson:"variables"`
	Readers    []string     `json:"readers" bson:"readers"`
	Writers    []string     `json:"writers" bson:"writers"`
	CreateAt   time.Time    `json:"createdAt" bson:"createdAt"`
	ModifiedAt time.Time    `json:"modifiedAt" bson:"modifiedAt"`
}
