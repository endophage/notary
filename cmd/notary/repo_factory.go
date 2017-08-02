package main

import (
	"github.com/spf13/viper"

	"net/http"

	"github.com/docker/notary"
	"github.com/docker/notary/client"
	"github.com/docker/notary/tuf/data"
)

const remoteConfigField = "api"

// RepoFactory takes a GUN and returns an initialized client.Repository, or an error.
type RepoFactory func(gun data.GUN) (client.Repository, error)

// ConfigureRepo takes in the configuration parameters and returns a repoFactory that can
// initialize new client.Repository objects with the correct upstreams and password
// retrieval mechanisms.
func ConfigureRepo(v *viper.Viper, retriever notary.PassRetriever, onlineOperation bool) RepoFactory {
	localRepo := func(gun data.GUN) (client.Repository, error) {
		var rt http.RoundTripper
		trustPin, err := getTrustPinning(v)
		if err != nil {
			return nil, err
		}
		if onlineOperation {
			rt, err = getTransport(v, gun, admin)
			if err != nil {
				return nil, err
			}
		}
		return client.NewFileCachedNotaryRepository(
			v.GetString("trust_dir"),
			gun,
			getRemoteTrustServer(v),
			rt,
			retriever,
			trustPin,
		)
	}

	// Leaving this in because it's correct code  that will get uncommented when we
	// get around to adding a future PR with the GRPC client.
	// remoteRepo := func(gun data.GUN) (client.Repository, error) {
	// 	conn, err := utils.GetGRPCClient(
	// 		v,
	// 		remoteConfigField,
	// 		grpcauth.NewCredStore(&passwordStore{false}, nil, nil),
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return api.NewClient(conn, gun), nil
	// }

	// if v.IsSet(remoteConfigField) {
	// 	return remoteRepo
	// }
	return localRepo
}
