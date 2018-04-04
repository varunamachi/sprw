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
