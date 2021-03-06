package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/now"
	"github.com/satori/go.uuid"
	"github.com/varunamachi/vaali/vapp"
	"github.com/varunamachi/vaali/vcmn"
	"github.com/varunamachi/vaali/vmgo"
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
	conn := vmgo.DefaultMongoConn()
	defer conn.Close()
	err = conn.C("entity").EnsureIndex(mgo.Index{
		Key:        []string{"name", "createdAt"},
		Unique:     false,
		DropDups:   false,
		Background: true, // See notes.
		Sparse:     true,
	})
	return vmgo.LogError("Sprw:Mongo", err)
}

//Reset - sparrow-entity reset
func Reset(app *vapp.App) (err error) {
	conn := vmgo.DefaultMongoConn()
	defer conn.Close()
	err = conn.C("users").DropIndex("name", "createdAt")
	if err == nil {
		_, err = conn.C("entity").RemoveAll(bson.M{})
	}
	return vmgo.LogError("Sprw:Mongo", err)
}

//CreateEntitySecret - creates a secret that can be used by an entity to
//communitcate with the server
func CreateEntitySecret(ent *Entity, owner string) (
	secret string, err error) {
	conn := vmgo.DefaultMongoConn()
	defer conn.Close()
	var entity Entity
	err = conn.C("entity").
		Find(bson.M{
			"_id": ent.OID,
		}).
		One(&entity)
	if err == nil {
		if entity.OwnerID != owner {
			err = fmt.Errorf("%s is not the owner of entity with ID %s",
				owner,
				ent.OID)
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
			"entityID": ent.OID,
			"owner":    owner,
		}, bson.M{
			"$set": bson.M{
				"secret": secretHash,
			},
		})
	}
	return secret, vmgo.LogError("Sprw:Mongo", err)
}

//AuthenticateEntity - authenticates an entity with given ID and owner using
//the provided secret
func AuthenticateEntity(entityID, owner, password string) (err error) {
	//generate hash for secret
	//get hash from database
	//compare...
	conn := vmgo.DefaultMongoConn()
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
	return vmgo.LogError("Sprw:Mongo", err)
}

//InsertParamValue - Insert param value given by entity into database
func InsertParamValue(owner string, value *ParamValue) (err error) {
	//use one document for an hour
	//use array of values for each parameter
	t := time.Now()
	day := now.New(t).BeginningOfDay()

	conn := vmgo.DefaultMongoConn()
	defer conn.Close()
	err = preallocateIfNotExist(conn)
	if err == nil {
		minuteSelector := fmt.Sprintf("values.%02d.%02d", t.Hour(), t.Minute())
		err = conn.C(owner).Update(bson.M{
			"$and": []bson.M{
				bson.M{"entityID": value.EntityID},
				bson.M{"paramID": value.ParamID},
				bson.M{"day": day},
			},
		}, bson.M{
			minuteSelector: value.Value,
		})
	}
	return vmgo.LogError("Sprw:Mongo", err)
}

func preallocateIfNotExist(conn *vmgo.MongoConn) (err error) {
	return err
}

//SetParam - set value for a parameter exposed by an entity
func SetParam(value ParamValue) (err error) {
	return vmgo.LogError("Sprw:Mongo", err)
}

//GetValuesForDateRange - get values for all parameters of an entity that
//is inserted by the entity between given days
func GetValuesForDateRange(
	entityID,
	owner,
	paramID string,
	dayRange vcmn.DateRange) (values []*ParamEntry, err error) {
	start := now.New(dayRange.From)
	end := now.New(dayRange.To)
	startDay := start.BeginningOfDay()
	endDay := end.EndOfDay()
	// startHr := start.Hour()
	// endHr := end.Hour()
	// startMn := start.Minute()
	// endMn := end.Minute()
	// mnPropStart := fmt.Sprintf("values.%02d", startHr)
	// mnPropEnd := fmt.Sprintf("values.%02d", endHr)
	values = make([]*ParamEntry, 0, 100)
	conn := vmgo.DefaultMongoConn()
	defer conn.Close()
	err = conn.C(owner).Find(bson.M{
		"$and": []bson.M{
			bson.M{"entityID": entityID},
			bson.M{"paramID": paramID},
			bson.M{"$gte": bson.M{"day": startDay}},
			bson.M{"$lte": bson.M{"day": endDay}},
		},
	}).All(&values)
	return values, vmgo.LogError("Sprw:Mongo", err)
}

//GetValuesForSingleDay - get values for all parameters of an entity that
//is inserted by the entity in a single day
func GetValuesForSingleDay(
	entityID,
	owner,
	paramID string,
	day time.Time) (values []*ParamEntry, err error) {
	start := now.New(day)
	end := now.New(day)
	startDay := start.BeginningOfDay()
	endDay := end.EndOfDay()
	// startHr := start.Hour()
	// endHr := end.Hour()
	// startMn := start.Minute()
	// endMn := end.Minute()
	// mnPropStart := fmt.Sprintf("values.%02d", startHr)
	// mnPropEnd := fmt.Sprintf("values.%02d", endHr)
	values = make([]*ParamEntry, 0, 100)
	conn := vmgo.DefaultMongoConn()
	defer conn.Close()
	err = conn.C(owner).Find(bson.M{
		"$and": []bson.M{
			bson.M{"entityID": entityID},
			bson.M{"paramID": paramID},
			bson.M{"$gte": bson.M{"day": startDay}},
			bson.M{"$lte": bson.M{"day": endDay}},
		},
	}).All(&values)
	return values, vmgo.LogError("Sprw:Mongo", err)
}

//ReadParamValue - read value for a parameter that is set by the user.
func ReadParamValue(entityID, paramName string) (
	val ParamValue, err error) {
	return val, vmgo.LogError("Sprw:Mongo", err)
}
