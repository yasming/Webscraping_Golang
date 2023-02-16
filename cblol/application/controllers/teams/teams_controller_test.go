package controller_teams_test

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"webscraping-league-of-legends-v2/cblol/domain/services/teams"
	"webscraping-league-of-legends-v2/cblol/infra/http/routes"
)

func TestPingRoute(t *testing.T) {
	t.Setenv("AUTHORIZED_USER", "user")
	t.Setenv("AUTHORIZED_USER_PASSWORD", "password")
	t.Setenv("CLASSIFICATION_URL_CBLOL", "https://flashscore.com.br/esports/league-of-legends/cblol/classificacao/")
	router := routes.AddRoutesWithAuthentication()
	mockGetTeamsFromFlashScore()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/get-teams", nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(os.Getenv("AUTHORIZED_USER"), os.Getenv("AUTHORIZED_USER_PASSWORD")))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\n    \"teams\": [\n        \"INTZ-eSports\"\n    ]\n}", w.Body.String())
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func mockGetTeamsFromFlashScore() {
	service_teams.GetTeamsFromFlashScoreGateway = func() []string {
		var strArr []string
		strings := append(strArr, "INTZ-eSports")
		return strings
	}
}
