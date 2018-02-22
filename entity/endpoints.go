package entity

import (
	"github.com/labstack/echo"
	"github.com/varunamachi/vaali/vnet"
	"github.com/varunamachi/vaali/vsec"
)

var endpoints = []*vnet.Endpoint{
	&vnet.Endpoint{
		Method:   echo.POST,
		URL:      "sprw/endpoint",
		Access:   vsec.Normal,
		Category: "entity",
		Func:     vnet.MakeCreateHandler("entity"),
		Comment:  "Create an entity",
	},
	&vnet.Endpoint{
		Method:   echo.PUT,
		URL:      "sprw/endpoint",
		Access:   vsec.Normal,
		Category: "entity",
		Func:     vnet.MakeUpdateHandler("entity"),
		Comment:  "Update an entity",
	},
	&vnet.Endpoint{
		Method:   echo.DELETE,
		URL:      "sprw/endpoint/:id",
		Access:   vsec.Normal,
		Category: "entity",
		Func:     vnet.MakeDeleteHandler("entity"),
		Comment:  "Delete an entity",
	},
	&vnet.Endpoint{
		Method:   echo.GET,
		URL:      "sprw/endpoint",
		Access:   vsec.Normal,
		Category: "entity",
		Func:     vnet.MakeGetHandler("entity"),
		Comment:  "Retrieves an entity",
	},
	&vnet.Endpoint{
		Method:   echo.GET,
		URL:      "sprw/endpoint",
		Access:   vsec.Normal,
		Category: "entity",
		Func:     vnet.MakeGetAllHandler("entity", "createdAt"),
		Comment:  "Retrieves range of entities",
	},
}
