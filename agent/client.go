package agent

import (
	"github.com/varunamachi/sprw/entity"
	"github.com/varunamachi/vaali/vlog"
	"github.com/varunamachi/vaali/vnet"
	"github.com/varunamachi/vaali/vsec"
)

//SparrowClient - struct to keep track of entity authentication information
type SparrowClient struct {
	*vnet.Client
	EntityID string
}

//NewClient - creates a new sparrow client
func NewClient(address, versionStr string) *SparrowClient {
	return &SparrowClient{
		Client: vnet.NewClient(address, "sprw", "0"),
	}
}

//EntityAuth - authenticate the entity
func (ec *SparrowClient) EntityAuth(
	entityID, owner, secret string) (err error) {
	rr := ec.Post(map[string]string{
		"entityID": entityID,
		"owner":    owner,
		"secret":   secret}, vsec.Public, "entity/auth")
	data := struct {
		Token string     `json:"token"`
		User  *vsec.User `json:"user"`
	}{}
	err = rr.Read(&data)
	if err == nil {
		ec.Token = data.Token
		ec.User = data.User
		ec.EntityID = entityID
	}
	return vlog.LogError("Sprw:Client", err)
}

//RenewAuth - renew auth token before it expires
func (ec *SparrowClient) RenewAuth() (err error) {
	rr := ec.Post(map[string]string{
		"entityID": ec.EntityID,
		"owner":    ec.User.ID,
	}, vsec.Public, "entity", "auth")
	var token string
	err = rr.Read(&token)
	if err == nil {
		ec.Client.Token = token
	}
	return vlog.LogError("Sprw:Client", err)
}

//InsertParamValue - inserts parameter value into
func (ec *SparrowClient) InsertParamValue(
	value entity.ParamValue) (err error) {
	rr := ec.Post(value, vsec.Normal, "")
	return vlog.LogError("Sprw:Entity", err)
}

func (ec *SparrowClient) ReadParamValue(paramID string) (
	val float32, err error) {

}
