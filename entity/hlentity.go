package entity

// import (
// 	"net/http"
// 	"time"

// 	"github.com/labstack/echo"
// 	"github.com/varunamachi/vaali/vdb"
// 	"github.com/varunamachi/vaali/vlog"
// 	"github.com/varunamachi/vaali/vnet"
// 	"gopkg.in/mgo.v2/bson"
// )

// func idM(id string) bson.M {
// 	return bson.M{"_id": bson.ObjectIdHex(id)}
// }

// func createEntity(ctx echo.Context) (err error) {
// 	status, msg := vnet.DefMS("Create Entity")
// 	var entity Entity
// 	err = ctx.Bind(&entity)
// 	if err == nil {
// 		entity.CreateAt = time.Now()
// 		err = vdb.Create(EntityColn, entity)
// 		if err != nil {
// 			msg = "Failed to create entity in database"
// 			status = http.StatusInternalServerError
// 		}
// 	} else {
// 		msg = "Failed to retrieve entity information from the request"
// 		status = http.StatusBadRequest
// 	}
// 	err = vnet.AuditedSendX(ctx, &entity, &vnet.Result{
// 		Status: status,
// 		Op:     "create_entity",
// 		Msg:    msg,
// 		OK:     err == nil,
// 		Data:   nil,
// 		Err:    err,
// 	})
// 	return vlog.LogError("S:Entity", err)
// }

// func updateEntity(ctx echo.Context) (err error) {
// 	status, msg := vnet.DefMS("Update Entity")
// 	var entity Entity
// 	err = ctx.Bind(&entity)
// 	if err == nil {
// 		entity.ModifiedAt = time.Now()
// 		err = vdb.Update(EntityColn, bson.M{"_id": entity.ID}, &entity)
// 		if err != nil {
// 			msg = "Failed to update entity in database"
// 			status = http.StatusInternalServerError
// 		}
// 	} else {
// 		msg = "Failed to retrieve entity info from request"
// 		status = http.StatusBadRequest
// 	}
// 	err = vnet.AuditedSendX(ctx, &entity, &vnet.Result{
// 		Status: status,
// 		Op:     "update_entity",
// 		Msg:    msg,
// 		OK:     err == nil,
// 		Data:   nil,
// 		Err:    err,
// 	})
// 	return vlog.LogError("S:Entity", err)
// }

// func deleteEntity(ctx echo.Context) (err error) {
// 	status, msg := vnet.DefMS("Delete Entity")
// 	entityID := ctx.Param("entityID")
// 	err = vdb.Delete(EntityColn, idM(entityID))
// 	if err != nil {
// 		msg = "Failed to delete entity from database"
// 		status = http.StatusInternalServerError
// 	}
// 	err = vnet.AuditedSend(ctx, &vnet.Result{
// 		Status: status,
// 		Op:     "delete_entity",
// 		Msg:    msg,
// 		OK:     err == nil,
// 		Data:   entityID,
// 		Err:    err,
// 	})
// 	return vlog.LogError("S:Entity", err)
// }

// func getEntity(ctx echo.Context) (err error) {
// 	status, msg := vnet.DefMS("Get Entity")
// 	var entity *Entity
// 	entityID := ctx.Param("entityID")
// 	err = vdb.Get(EntityColn, idM(entityID), entity)
// 	if err != nil {
// 		msg = "Failed to retrieve from database, entity with ID: " + entityID
// 		status = http.StatusInternalServerError
// 	}
// 	err = vnet.SendAndAuditOnErr(ctx, &vnet.Result{
// 		Status: status,
// 		Op:     "get_entity",
// 		Msg:    msg,
// 		OK:     err == nil,
// 		Data:   entity,
// 		Err:    err,
// 	})
// 	return vlog.LogError("S:Entity", err)
// }

// func getEntities(ctx echo.Context) (err error) {
// 	status, msg := vnet.DefMS("Get Entities")
// 	offset, limit, has := vnet.GetOffsetLimit(ctx)
// 	var entities []*Entity
// 	if has {
// 		entities = make([]*Entity, 0, limit)
// 		err = vdb.GetAll(EntityColn, "-createdAt", offset, limit, entities)
// 		if err != nil {
// 			msg = "Failed to retrieve entities from database"
// 			status = http.StatusInternalServerError
// 		}
// 	}
// 	err = vnet.SendAndAuditOnErr(ctx, &vnet.Result{
// 		Status: status,
// 		Op:     "get_entities",
// 		Msg:    msg,
// 		OK:     err == nil,
// 		Data:   entities,
// 		Err:    err,
// 	})
// 	return vlog.LogError("S:Entity", err)
// }
