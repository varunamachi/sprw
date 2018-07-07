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
			Method:   echo.POST,
			URL:      "entity/auth",
			Access:   vsec.Public,
			Category: "entity",
			Func:     authenticateEntity,
			Comment:  "Authenticate an entity",
		},
		&vnet.Endpoint{
			Method:   echo.POST,
			URL:      "entity/param",
			Access:   vsec.Normal,
			Category: "entity",
			Func:     insertParamValue,
			Comment:  "Insert value for a parameter",
		},
		&vnet.Endpoint{
			Method:   echo.GET,
			URL:      "entity/param/:day",
			Access:   vsec.Normal,
			Category: "entity",
			Func:     getParamValueForSingleDay,
			Comment:  "Get parameter value for entire day",
		},
		&vnet.Endpoint{
			Method:   echo.GET,
			URL:      "entity/param/:from/:to",
			Access:   vsec.Normal,
			Category: "entity",
			Func:     getParamValueForDateRange,
			Comment:  "Get parameter value for date range",
		},
		&vnet.Endpoint{
			Method:   echo.POST,
			URL:      "entity/param/config",
			Access:   vsec.Normal,
			Category: "entity",
			Func:     setParamValue,
			Comment:  "Set a param value as config",
		},
		&vnet.Endpoint{
			Method:   echo.GET,
			URL:      "entity/param/config",
			Access:   vsec.Normal,
			Category: "entity",
			Func:     getParamValue,
			Comment:  "Get configured parameter value",
		},
	}
	return endpoints
}
