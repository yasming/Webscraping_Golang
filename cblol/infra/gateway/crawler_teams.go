package gateway

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"time"
)

func GetTeamsFromFlashScore() []string {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	var teams []string
	err := chromedp.Run(ctx,
		chromedp.Navigate(os.Getenv("CLASSIFICATION_URL_CBLOL")),
		chromedp.WaitVisible(`.tableCellParticipant__name`),
		chromedp.Evaluate(`[...document.querySelectorAll('.tableCellParticipant__name')].map((e) => e.innerText)`, &teams),
	)
	if err != nil {
		log.Fatal(err)
	}
	return teams
}
