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
	OID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Type       string        `json:"type" bson:"type"`
	Location   string        `json:"location" bson:"location"`
	OwnerID    string        `json:"ownerID" bson:"ownerID"`
	OwnerName  string        `json:"ownerName" bson:"ownerName"`
	Variables  []vcmn.Param  `json:"variables" bson:"variables"`
	Readers    []string      `json:"readers" bson:"readers"`
	Writers    []string      `json:"writers" bson:"writers"`
	Tags       []string      `json:"tags" bson:"tags"`
	CreatedAt  time.Time     `json:"createdAt" bson:"createdAt"`
	ModifiedAt time.Time     `json:"modifiedAt" bson:"modifiedAt"`
	CreatedBy  string        `json:"createdBy" bson:"createdBy"`
	ModifiedBy string        `json:"modifiedBy" bson:"modifiedBy"`
}

//SetCreationInfo - set the creation time and creator
func (e *Entity) SetCreationInfo(at time.Time, by string) {
	e.CreatedAt = at
	e.CreatedBy = by
}

//SetModInfo - set modification time and modifier
func (e *Entity) SetModInfo(at time.Time, by string) {
	e.ModifiedAt = at
	e.ModifiedBy = by
}

//ID - get ID of the entity
func (e *Entity) ID() bson.ObjectId {
	return e.OID
}

//ParamDesc - describes a parameter
type ParamDesc struct {
	EntityID  string `json:"entityID" bson:"entityID"`
	ParamID   string `json:"param" bson:"param"`
	ParamName string `json:"paramName" bson:"paramName"`
	Unit      string `json:"unit" bson:"unit"`
}

//ParamValue - parameter value along with parameter description
type ParamValue struct {
	ParamDesc `bson:",inline"`
	Value     float32 `json:"value" bson:"value"`
}

//ParamEntry - data collection entry for a parameter for a day with
// granularity of 1 minute
type ParamEntry struct {
	ParamDesc `bson:",inline"`
	Day       time.Time               `json:"day" bson:"day"`
	Total     float64                 `json:"total" bson:"total"`
	Values    map[int]map[int]float32 `json:"values" bson:"values"`
}
