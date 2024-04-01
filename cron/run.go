package cron

import (
	"crypto/tls"
	"log"
	"sync"
	"time"
	"tls-watch/api/store"
)

var waitGroup sync.WaitGroup

func checkTlsValidity(hostname string) (time.Time, error) {
	conn, err := tls.Dial("tcp", hostname+":443", nil)
	if err != nil {
		return time.Now(), err
	}

	err = conn.VerifyHostname(hostname)
	if err != nil {
		return time.Now(), err
	}

	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	return expiry, nil
}

func processDomain(domain store.Domain) {
	defer waitGroup.Done()
	expiry, err := checkTlsValidity(domain.Domain)
	if err != nil {
		log.Printf("checking tls validity for %s failed: %v", domain.Domain, err)
	}

	expiresIn := int(time.Until(expiry).Hours() / 24)
	if expiresIn > 10 {
		return
	}

	log.Printf("%v expires in %d days", domain, expiresIn)

	notificationSettings, err := store.GetAllNotificationSettingsByUserID(domain.UserID)
	if err != nil {
		log.Printf("fetching notification settings for %d failed: %v", domain.UserID, err)
	}

	for _, setting := range *notificationSettings {
		if setting.Provider == "telegram" {
			SendTelegramMessage(&domain, expiry, setting.ProviderUserID)
		}
	}
}

func Run() {
	domains, err := store.GetAllDomains()
	if err != nil {
		log.Fatalf("fetching list of domains failed: %v", err)
	}

	for _, domain := range *domains {
		log.Printf("processing %v", domain)
		waitGroup.Add(1)
		go processDomain(domain)
	}

	waitGroup.Wait()
}
