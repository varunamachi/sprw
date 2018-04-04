package entity

import (
	"github.com/labstack/echo"
	"github.com/varunamachi/vaali/vnet"
	"github.com/varunamachi/vaali/vsec"
)

//GetEndpoints - gives REST endpoints for entity APIs
func GetEndpoints() (endpoints []*vnet.Endpoint) {
	endpoints = []*vnet.Endpoint{
		&vnet.Endpoint{
			Method:   echo.POST,
			URL:      "entity/secret",
			Access:   vsec.Normal,
			Category: "entity",
			Func:     genEntityPassword,
			Comment:  "Create Entity Password",
		},
		&vnet.Endpoint{
			Method:   echo.PUT,
			URL:      "entity/auth",
			Access:   vsec.Public,
			Category: "entity",
			Func:     authenticateEntity,
			Comment:  "Authenticate an entity",
		},
	}
	return endpoints
}
