package net

import (
	"github.com/varunamachi/vaali/vlog"
	"github.com/varunamachi/vaali/vnet"
	"github.com/varunamachi/vaali/vsec"
)

//SparrowClient - struct to keep track of entity authentication information
type SparrowClient struct {
	*vnet.Client
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
		ec.Client.Token = data.Token
		ec.Client.User = data.User
	}
	return vlog.LogError("Sprw:Client", err)
}
