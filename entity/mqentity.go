package entity

import (
	"errors"
	"fmt"

	"github.com/jinzhu/now"
	"github.com/satori/go.uuid"
	"github.com/varunamachi/vaali/vapp"
	"github.com/varunamachi/vaali/vcmn"
	"github.com/varunamachi/vaali/vdb"
	"github.com/varunamachi/vaali/vlog"
	"gopkg.in/hlandau/passlib.v1"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Init - sparrow-entity initialization
func Init(app *vapp.App) (err error) {
	return err
}

//Setup - setup function for entity module
func Setup(app *vapp.App) (err error) {
	conn := vdb.DefaultMongoConn()
	defer conn.Close()
	err = conn.C("entity").EnsureIndex(mgo.Index{
		Key:        []string{"name", "createdAt"},
		Unique:     false,
		DropDups:   false,
		Background: true, // See notes.
		Sparse:     true,
	})
	return vlog.LogError("Sprw:Mongo", err)
}

//Reset - sparrow-entity reset
func Reset(app *vapp.App) (err error) {
	conn := vdb.DefaultMongoConn()
	defer conn.Close()
	err = conn.C("users").DropIndex("name", "createdAt")
	if err == nil {
		_, err = conn.C("entity").RemoveAll(bson.M{})
	}
	return vlog.LogError("Sprw:Mongo", err)
}

//CreateEntitySecret - creates a secret that can be used by an entity to
//communitcate with the server
func CreateEntitySecret(entityID string, owner string) (
	secret string, err error) {
	conn := vdb.DefaultMongoConn()
	defer conn.Close()
	var entity Entity
	err = conn.C("entity").
		Find(bson.M{
			"_id": bson.ObjectIdHex(entityID),
		}).
		One(&entity)
	if err == nil {
		if entity.Owner != owner {
			err = fmt.Errorf("%s is not the owner of entity with ID %s",
				owner,
				entityID)
		}
	}
	if err != nil {
		return secret, err
	}
	secret = uuid.NewV4().String()
	var secretHash string
	secretHash, err = passlib.Hash(secret)
	if err == nil {
		_, err = conn.C("entity_secret").Upsert(bson.M{
			"entityID": entityID,
			"owner":    owner,
		}, bson.M{
			"$set": bson.M{
				"secret": secretHash,
			},
		})
	}
	return secret, vlog.LogError("Sprw:Mongo", err)
}

//AuthenticateEntity - authenticates an entity with given ID and owner using
//the provided secret
func AuthenticateEntity(entityID, owner, password string) (err error) {
	//generate hash for secret
	//get hash from database
	//compare...
	conn := vdb.DefaultMongoConn()
	defer conn.Close()
	secret := bson.M{}
	err = conn.C("entity_secret").
		Find(bson.M{
			"entityID": entityID,
			"owner":    owner,
		}).
		Select(bson.M{
			"secret": 1,
			"_id":    0,
		}).
		One(&secret)
	if err == nil {
		storedHash, ok := secret["secret"].(string)
		if ok {
			var newHash string
			newHash, err = passlib.Verify(password, storedHash)
			if err == nil && newHash != "" {
				_, err = conn.C("entity_secret").Upsert(bson.M{
					"entityID": entityID,
					"owner":    owner,
				}, bson.M{
					"$set": bson.M{
						"secret": newHash,
					},
				})
			}
		} else {
			err = errors.New("Failed to varify password")
		}
	}
	return vlog.LogError("Sprw:Mongo", err)
}

//InsertParamValue - Insert param value given by entity into database
func InsertParamValue(value ParamValue) (err error) {
	//use one document for an hour
	//use array of values for each parameter
	return vlog.LogError("Sprw:Mongo", err)
}

//SetParam - set value for a parameter exposed by an entity
func SetParam(value ParamValue) (
	err error) {
	return vlog.LogError("Sprw:Mongo", err)
}

//GetValues - get values for all parameters of an entity that is inserted by
//the entity
func GetValues(entityID, owner, paramID string, dayRange vcmn.DateRange) (
	values []*ParamEntry,
	err error) {
	start := now.New(dayRange.From)
	end := now.New(dayRange.To)
	startDay := start.BeginningOfDay()
	endDay := end.EndOfDay()
	startHr := start.Hour()
	endHr := end.Hour()
	startMn := start.Minute()
	endMn := end.Minute()
	mnPropStart := fmt.Sprintf("values.%d", startHr)
	mnPropEnd := fmt.Sprintf("values.%d", endHr)
	values = make([]*ParamEntry, 0, 100)
	conn := vdb.DefaultMongoConn()
	defer conn.Close()
	err = conn.C(owner).Find(bson.M{
		"$and": []bson.M{
			bson.M{"entityID": entityID},
			bson.M{"paramID": paramID},
			bson.M{"$gte": bson.M{"day": startDay}},
			bson.M{"$lte": bson.M{"day": endDay}},
			bson.M{"$gte": bson.M{mnPropStart: startMn}},
			bson.M{"$lte": bson.M{mnPropEnd: endMn}},
		},
	}).All(&values)
	return values, vlog.LogError("Sprw:Mongo", err)
}

//ReadParamValue - read value for a parameter that is set by the user.
func ReadParamValue(entityID, paramName string) (
	val ParamValue, err error) {
	return val, vlog.LogError("Sprw:Mongo", err)
}
