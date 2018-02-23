package entity

import (
	"time"

	"github.com/labstack/echo"
	"github.com/varunamachi/vaali/vnet"
	"github.com/varunamachi/vaali/vsec"
)

//GetEndpoints - REST endpoints for entity related operations
func GetEndpoints() []*vnet.Endpoint {
	return []*vnet.Endpoint{
		&vnet.Endpoint{
			Method:   echo.POST,
			URL:      "sprw/endpoint",
			Access:   vsec.Normal,
			Category: "entity",
			Func: vnet.MakeCreateHandlerX("entity",
				func(ctx echo.Context) (d interface{}, err error) {
					entity := &Entity{}
					err = ctx.Bind(entity)
					if err == nil {
						entity.CreateAt = time.Now()
					}
					return entity, err
				}),
			Comment: "Create an entity",
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
			URL:      "sprw/endpoint/:id",
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
}
