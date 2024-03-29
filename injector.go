//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	apiCandidatePair "smartchoice/api/candidate_pair"
	apiCandidateProfile "smartchoice/api/candidate_profile"
	apiDebate "smartchoice/api/debate"
	apiParty "smartchoice/api/party"
	"smartchoice/domain/candidate_pair"
	"smartchoice/domain/candidate_profile"
	"smartchoice/domain/debate"
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

var debateSet = wire.NewSet(
	debate.NewDebateRepository,
	wire.Bind(new(debate.Repository), new(*debate.RepositoryImpl)),
	debate.NewDebateService,
	wire.Bind(new(debate.Service), new(*debate.ServiceImpl)),
	apiDebate.NewDebateHandler,
	wire.Bind(new(apiDebate.Handler), new(*apiDebate.HandlerImpl)),
)

var candidatePairSet = wire.NewSet(
	candidate_pair.NewCandidatePairRepository,
	wire.Bind(new(candidate_pair.Repository), new(*candidate_pair.RepositoryImpl)),
	candidate_pair.NewCandidatePairService,
	wire.Bind(new(candidate_pair.Service), new(*candidate_pair.ServiceImpl)),
	apiCandidatePair.NewCandidatePairHandler,
	wire.Bind(new(apiCandidatePair.Handler), new(*apiCandidatePair.HandlerImpl)),
)

var candidateProfileSet = wire.NewSet(
	candidate_profile.NewCandidateProfileRepository,
	wire.Bind(new(candidate_profile.Repository), new(*candidate_profile.RepositoryImpl)),
	candidate_profile.NewCandidateProfileService,
	wire.Bind(new(candidate_profile.Service), new(*candidate_profile.ServiceImpl)),
	apiCandidateProfile.NewCandidateProfileHandler,
	wire.Bind(new(apiCandidateProfile.Handler), new(*apiCandidateProfile.HandlerImpl)),
)

func InitializedApp() *gin.Engine {
	wire.Build(
		infrastructure.ConnectToDatabase,
		partySet,
		debateSet,
		candidatePairSet,
		candidateProfileSet,
		infrastructure.SetupRoutes,
	)
	return nil
}
