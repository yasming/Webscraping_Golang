package routes

import (
	"github.com/gin-gonic/gin"
	"os"
	"webscraping-league-of-legends-v2/cblol/application/controllers/teams"
)

func AddRoutesWithAuthentication() *gin.Engine {
	router := gin.Default()
	authorized := router.Group("/api", gin.BasicAuth(gin.Accounts{
		os.Getenv("AUTHORIZED_USER"): os.Getenv("AUTHORIZED_USER_PASSWORD"),
	}))
	authorized.GET("/get-teams", controller_teams.GetTeams)
	return router
}
