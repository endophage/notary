package remoteks

import (
	"github.com/docker/notary/trustmanager"
	"google.golang.org/grpc"
	"github.com/docker/notary/tuf/data"
	"fmt"
)

type RemoteKS struct {
	client StoreClient
	location string
}

func NewRemoteKS(server string) (trustmanager.KeyStore, error) {
	cc, err := grpc.Dial(server)
	if err != nil {
		return nil, err
	}
	return &RemoteKS{
		client : NewStoreClient(cc),
		location: server,
	}, nil
}

// AddKey adds a key to the KeyStore, and if the key already exists,
// succeeds.  Otherwise, returns an error if it cannot add.
func (ks *RemoteKS) AddKey(keyInfo trustmanager.KeyInfo, privKey data.PrivateKey) error {
	return nil
}

// Should fail with ErrKeyNotFound if the keystore is operating normally
// and knows that it does not store the requested key.
func (ks *RemoteKS) GetKey(keyID string) (data.PrivateKey, string, error) {
	return nil, "", nil
}

func (ks *RemoteKS) GetKeyInfo(keyID string) (trustmanager.KeyInfo, error){
	return trustmanager.KeyInfo{}, nil
}

func (ks *RemoteKS) ListKeys() map[string]trustmanager.KeyInfo{
	return nil
}

func (ks *RemoteKS) RemoveKey(keyID string) error{
	return nil
}

func (ks *RemoteKS) Name() string{
	return fmt.Sprintf("Remote Key Store @ %s", ks.location)
}