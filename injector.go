//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	apiParty "smartchoice/api/party"
	"smartchoice/domain/party"
	"smartchoice/infrastructure"
)

var partySet = wire.NewSet(
	party.NewPartyRepository,
	wire.Bind(new(party.Repository), new(*party.RepositoryImpl)),
	party.NewPartyService,
	wire.Bind(new(party.Service), new(*party.ServiceImpl)),
	apiParty.NewPartyHandler,
	wire.Bind(new(apiParty.Handler), new(*apiParty.HandlerImpl)),
)

func InitializedApp() *gin.Engine {
	wire.Build(
		infrastructure.ConnectToDatabase,
		partySet,
		infrastructure.SetupRoutes,
	)
	return nil
}
