package entity

import (
	"errors"
	"fmt"

	"github.com/satori/go.uuid"
	"github.com/varunamachi/vaali/vapp"
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
	return secret, err
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
	return vlog.LogError("UMan:Mongo", err)
	return err
}
