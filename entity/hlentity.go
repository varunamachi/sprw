package entity

import (
	"github.com/labstack/echo"
	"github.com/varunamachi/vaali/vlog"
)

func createEntity(ctx echo.Context) (err error) {
	// status, msg := vnet.DefMS()
	return vlog.LogError("S:Entity", err)
}
