package sprw

import (
	"github.com/varunamachi/sprw/entity"
	"github.com/varunamachi/vaali/vapp"
	"github.com/varunamachi/vaali/vmgo"
	"github.com/varunamachi/vaali/vnet"
	"gopkg.in/urfave/cli.v1"
)

//NewModule - creates new sparrow module
func NewModule() *vapp.Module {
	return &vapp.Module{
		Name:        "sprw",
		Description: "The sparrow server",
		Endpoints:   []*vnet.Endpoint{},
		Commands:    []cli.Command{},
		Factories: []vapp.Factory{
			vapp.Factory{
				DataType: "entity",
				Func: func() vmgo.StoredItem {
					return &entity.Entity{}
				},
			},
		},
		Initialize: entity.Init,
		Setup:      entity.Setup,
		Reset:      entity.Reset,
	}
}
