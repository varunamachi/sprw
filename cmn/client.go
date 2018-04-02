package cmn

import (
	"github.com/varunamachi/vaali/vnet"
)

//SprwClient - Sparrow HTTP client
type SprwClient struct {
	vnet.Client
}

//NewSprwClient - creates a new sparrow client
func NewSprwClient(serverURL string, version string) *SprwClient {
	client := &SprwClient{
		Client: vnet.Client{
			Address:    serverURL,
			BaseURL:    "sprw",
			VersionStr: version,
		},
	}
	return client
}

//EntityAuth - authenticate an entity
func (cl *SprwClient) EntityAuth(entityID, owner, secret string) {

}
