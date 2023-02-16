package service_teams

import (
	"webscraping-league-of-legends-v2/cblol/infra/gateway"
)

var GetTeamsFromFlashScoreGateway = gateway.GetTeamsFromFlashScore

func GetTeams() []string {
	return GetTeamsFromFlashScoreGateway()
}
