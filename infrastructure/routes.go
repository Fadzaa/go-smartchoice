package infrastructure

import (
	"github.com/gin-gonic/gin"
	"smartchoice/api/candidate_pair"
	"smartchoice/api/candidate_profile"
	"smartchoice/api/debate"
	"smartchoice/api/party"
)

func SetupRoutes(
	partyHandler party.Handler,
	debateHandler debate.Handler,
	candidatePairHandler candidate_pair.Handler,
	candidateProfileHandler candidate_profile.Handler,
) *gin.Engine {
	router := gin.Default()

	party.RoutesParty(partyHandler, router)
	debate.RoutesDebate(debateHandler, router)
	candidate_pair.RouteCandidatePair(candidatePairHandler, router)
	candidate_profile.RoutesCandidateProfile(candidateProfileHandler, router)

	return router
}
