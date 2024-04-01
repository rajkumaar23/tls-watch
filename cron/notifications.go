package cron

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	"tls-watch/api/store"
)

func SendTelegramMessage(domain *store.Domain, expiry time.Time, telegramUserID string) {
	message := fmt.Sprintf(
		"<b>tls expiry alert for %v</b>\n\n"+
			"⏰ certificate expiry date is %v\n\n"+
			"⏳ expires in <b>%v</b> days",
		domain.Domain,
		expiry.Format(time.RFC822),
		int(time.Until(expiry).Hours()/24),
	)

	params := url.Values{}
	params.Set("chat_id", telegramUserID)
	params.Set("text", message)
	params.Set("parse_mode", "HTML")
	telegramURL := fmt.Sprintf(
		"https://api.telegram.org/bot%v/sendMessage?%v", os.Getenv("TELEGRAM_BOT_TOKEN"), params.Encode(),
	)

	res, err := http.Get(telegramURL)
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode != 200 || err != nil {
		log.Printf("sending message for %s to %s failed, body : %v, err : %v", domain.Domain, telegramUserID, string(body), err)
		return
	}
}
