package cron

import (
	"fmt"
	"time"
	"tls-watch/api/store"
)

type NotificationProvider interface {
	send(domain *store.Domain, expiry time.Time) error
}

func NotificationProviderFactory(providerName string, userID string) (NotificationProvider, error) {
	switch providerName {
	case "telegram":
		return Telegram{userID: userID}, nil
	default:
		return nil, fmt.Errorf("unknown notification provider: %s", providerName)
	}
}
