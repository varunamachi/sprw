package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/varunamachi/vaali/vcmn"
)

//EntityColn - entity collection name
const EntityColn = "entity"

//VariableAccess - access level of variable
type VariableAccess int

//Entity - represents a entiry with variables
type Entity struct {
	ID         bson.ObjectId `json:"_id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Type       string        `json:"type" bson:"type"`
	Location   string        `json:"location" bson:"location"`
	Owner      string        `json:"owner" bson:"owner"`
	Variables  []vcmn.Param  `json:"variables" bson:"variables"`
	Readers    []string      `json:"readers" bson:"readers"`
	Writers    []string      `json:"writers" bson:"writers"`
	CreateAt   time.Time     `json:"createdAt" bson:"createdAt"`
	ModifiedAt time.Time     `json:"modifiedAt" bson:"modifiedAt"`
}
