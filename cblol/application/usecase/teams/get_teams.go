package usecase_teams

import (
	"webscraping-league-of-legends-v2/cblol/domain/services/teams"
)

func GetTeams() []string {
	return service_teams.GetTeams()
}
