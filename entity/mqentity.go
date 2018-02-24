package entity

import (
	"github.com/varunamachi/vaali/vapp"
	"github.com/varunamachi/vaali/vdb"
	"github.com/varunamachi/vaali/vlog"
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
