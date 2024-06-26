package cron

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
	"tls-watch/api/store"
)

type Telegram struct {
	userID string
}

func (t Telegram) send(domain *store.Domain, expiry time.Time) error {
	message := fmt.Sprintf(
		"<b>tls expiry alert for %v</b>\n\n"+
			"⏰ certificate expiry date is %v\n\n"+
			"⏳ expires in <b>%v</b> days",
		domain.Domain,
		expiry.Format(time.RFC822),
		int(time.Until(expiry).Hours()/24),
	)

	params := url.Values{}
	params.Set("chat_id", t.userID)
	params.Set("text", message)
	params.Set("parse_mode", "HTML")
	telegramURL := fmt.Sprintf(
		"https://api.telegram.org/bot%v/sendMessage?%v", os.Getenv("TELEGRAM_BOT_TOKEN"), params.Encode(),
	)

	res, err := http.Get(telegramURL)
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode != 200 || err != nil {
		return fmt.Errorf("cron: sending message for %s to %s failed, body : %v, err : %v", domain.Domain, t.userID, string(body), err)
	}

	return nil
}
