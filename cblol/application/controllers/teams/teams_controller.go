package controller_teams

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webscraping-league-of-legends-v2/cblol/application/usecase/teams"
)

func GetTeams(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"teams": usecase_teams.GetTeams()})
}
