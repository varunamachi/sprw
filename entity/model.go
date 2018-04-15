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
	OID        bson.ObjectId `json:"_id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Type       string        `json:"type" bson:"type"`
	Location   string        `json:"location" bson:"location"`
	Owner      string        `json:"owner" bson:"owner"`
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

//ParamValueEntry - entry for bunch of values for a parameter associated with
//an entity
// type ParamValueEntry struct {
// 	EntityID   string                       `json:"entityID" bson:"entityID"`
// 	EnitiyName string                       `json:"entityName" bson:"entityName"`
// 	Hour       int                          `json:"hour" bson:"hour"`
// 	Values     map[string][]vcmn.ParamValue `json:"values" bson:"values"`
// }

// type HourEntry struct {
// 	Hour   time.Time       `json:"hour" bson:"hour"`
// 	Values map[int]float32 `json:"values" bson:"values"`
// 	States map[int]string  `json:"states" bson:"states"`
// }

type ParamDesc struct {
	ParamID   string `json:"param" bson:"param"`
	ParamName string `json:"paramName" bson:"paramName"`
	Unit      string `json:"unit" bson:"unit"`
}

type ParamValue struct {
	ParamDesc
	Value float32 `json:"value"`
}

//ParamEntry - data collection entry for a parameter for a day
type ParamEntry struct {
	Total     float64                 `json:"total" bson:"total"`
	Values    map[int]map[int]float32 `json:"values" bson:"values"`
	ParamDesc `bson:",inline"`
}

//EntityEntry - data collection entry for a day for a an entity
// type EntityEntry struct {
// 	EntityID string
// 	Day      time.Time             `json:"day" bson:"day"`
// 	Params   map[string]ParamEntry `json:"params" bson:"params"`
// }
