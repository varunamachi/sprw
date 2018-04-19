package entity

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/varunamachi/vaali/vcmn"
	"github.com/varunamachi/vaali/vlog"
	"github.com/varunamachi/vaali/vnet"
	"github.com/varunamachi/vaali/vsec"
)

func genEntityPassword(ctx echo.Context) (err error) {
	status, msg := vnet.DefMS("Gen entity password")
	entityID := ctx.Param("entityID")
	var secret string
	secret, err = CreateEntitySecret(entityID, vnet.GetString(ctx, "userID"))
	vnet.AuditedSendSecret(ctx, &vnet.Result{
		Status: status,
		Op:     "entity_gen_secret",
		Msg:    msg,
		OK:     err == nil,
		Data:   secret,
		Err:    vcmn.ErrString(err),
	})
	return vlog.LogError("Sprw:Net", err)
}

func authenticateEntity(ctx echo.Context) (err error) {
	status, msg := vnet.DefMS("Gen entity password")
	creds := struct {
		EntityID string `json:"entityID"`
		Owner    string `json:"owner"`
		Secret   string `json:"secret"`
	}{}
	err = ctx.Bind(&creds)
	if err == nil {
		err = AuthenticateEntity(creds.EntityID, creds.Owner, creds.Secret)
	}
	var data map[string]interface{}
	if err == nil {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["entityID"] = creds.EntityID
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		claims["access"] = vsec.Normal
		var signed string
		//@TODO get key from somewhere
		signed, err = token.SignedString("valrrwwssffgsdgfksdjfghsdlgnsda")
		if err == nil {
			data = make(map[string]interface{})
			data["token"] = signed
		} else {
			msg = "Failed to sign token"
			status = http.StatusInternalServerError
		}
	}
	//generate JWT token and send
	vnet.AuditedSendX(ctx, vlog.M{
		"entityID": creds.EntityID,
		"owner":    creds.Owner,
	}, &vnet.Result{
		Status: status,
		Op:     "entity_gen_secret",
		Msg:    msg,
		OK:     err == nil,
		Data:   data,
		Err:    vcmn.ErrString(err),
	})
	return vlog.LogError("Sprw:Net", err)
}

func insertParamValue(ctx echo.Context) (err error) {
	status, msg := vnet.DefMS("Insert param value")
	var paramValue ParamValue
	err = ctx.Bind(&paramValue)
	owner := GetEntityOwner(ctx)
	if err == nil {
		err = InsertParamValue(owner, &paramValue)
		if err != nil {
			msg = "Failed to add parameter value into database"
			status = http.StatusInternalServerError
		}
	} else {
		msg = "Could not retrieve parameter value"
		status = http.StatusBadRequest
	}
	vnet.AuditedSend(ctx, &vnet.Result{
		Status: status,
		Op:     "entity_insert_value",
		Msg:    msg,
		OK:     err == nil,
		Data: vlog.M{
			"owner": owner,
			"value": paramValue,
		},
		Err: vcmn.ErrString(err),
	})
	return vlog.LogError("Sprw:Net", err)
}

func getParamValueForSingleDay(ctx echo.Context) (err error) {
	status, msg := vnet.DefMS("Get param value for a day")
	var vals []*ParamEntry
	entityID := ctx.Param("entityID")
	paramID := ctx.Param("paramID")
	dayStr := ctx.Param("day")
	var day time.Time
	day, err = time.Parse(time.RFC3339Nano, dayStr)
	owner := GetEntityOwner(ctx)
	if err == nil {
		vals, err = GetValuesForSingleDay(entityID, owner, paramID, day)
		if err != nil {
			msg = "Could not retrieve param value from database"
			status = http.StatusInternalServerError
		}
	} else {
		msg = "Invalid date provided"
		status = http.StatusBadRequest
	}
	vnet.SendAndAuditOnErr(ctx, &vnet.Result{
		Status: status,
		Op:     "entity_get_day_vals",
		Msg:    msg,
		OK:     err == nil,
		Data:   vals,
		Err:    vcmn.ErrString(err),
	})
	return vlog.LogError("Sprw:Net", err)
}

func getParamValueForDateRange(ctx echo.Context) (err error) {
	status, msg := vnet.DefMS("Get param value for day date range")
	var vals []*ParamEntry
	entityID := ctx.Param("entityID")
	paramID := ctx.Param("paramID")
	var dateRange vcmn.DateRange
	dateRange, err = vnet.GetDateRange(ctx)
	owner := ctx.Param("owner")
	if err == nil {
		vals, err = GetValuesForDateRange(entityID, owner, paramID, dateRange)
		if err != nil {
			msg = "Could not retrieve param value from database"
			status = http.StatusInternalServerError
		}
	} else {
		msg = "Invalid date range provided"
		status = http.StatusBadRequest
	}
	vnet.SendAndAuditOnErr(ctx, &vnet.Result{
		Status: status,
		Op:     "entity_get_dayrange_vals",
		Msg:    msg,
		OK:     err == nil,
		Data:   vals,
		Err:    vcmn.ErrString(err),
	})
	return vlog.LogError("Sprw:Net", err)
}

//GetEntityOwner - retrieves name of the owner of the entity from context
func GetEntityOwner(ctx echo.Context) (owner string) {
	return owner
}
