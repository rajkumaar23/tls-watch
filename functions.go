package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func ProcessHostname(hostname string) {
	defer WaitGroup.Done()
	fmt.Println("Checking SSL validity for", hostname)
	expiry, err := CheckSSLValidity(hostname)
	if err != nil {
		fmt.Println("Validation failed for", hostname, "-", err.Error())
		return
	}
	expiresIn := int(time.Until(expiry).Hours() / 24)
	fmt.Println("Certificate for", hostname, "expires in", expiresIn, "days")
	if expiresIn < config.Threshold {
		SendAlert(hostname, expiry, expiresIn)
	}
}

func CheckSSLValidity(hostname string) (time.Time, error) {
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

func SendAlert(hostname string, expiry time.Time, expiresIn int) {
	message := fmt.Sprintf(
		"<b>SSL Expiry Alert for %v</b>\n\n"+
			"⏰Certificate expiry date is %v\n\n"+
			"⏳Expires in <b>%v</b> days",
		hostname,
		expiry.Format(time.RFC822),
		expiresIn,
	)

	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(config.TelegramID))
	params.Set("text", message)
	params.Set("parse_mode", "HTML")
	telegramURL := fmt.Sprintf(
		"https://api.telegram.org/bot%v/sendMessage?%v", config.TelegramBotToken, params.Encode(),
	)

	req, _ := http.NewRequest(http.MethodGet, telegramURL, nil)
	res, err := HttpClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	errorMessagePrefix := fmt.Sprintf("Sending message for %v failed :", hostname)
	if res.StatusCode != 200 {
		fmt.Println(errorMessagePrefix, string(body))
		return
	}
	if err != nil {
		fmt.Println(errorMessagePrefix, err)
		return
	}

	fmt.Println("Sent message for", hostname)
}
