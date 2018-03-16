package amazon

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

// doc: https://developer.amazon.com/docs/in-app-purchasing/iap-rvs-for-android-apps.html
const (
	SandboxHost    string = "http://localhost:8080/RVSSandbox"
	ProductionHost string = "https://appstore-sdk.amazon.com/"
)

type Config struct {
	IsSandbox      bool
	SandboxHost    string
	Secret         string
	ConnectTimeout int64
}

type Response struct {
	BetaProduct     bool    `json:"betaProduct"`
	CancelDate      *int64  `json:"cancelDate,omitempty"`
	ParentProductId *string `json:"parentProductId,omitempty"`
	ProductId       string  `json:"productId"`
	ProductType     string  `json:"productType"`
	PurchaseDate    int64   `json:"purchaseDate"`
	Quantity        *int    `json:"quantity,omitempty"`
	ReceiptId       string  `json:"receiptId"`
	RenewalDate     *int64  `json:"renewalDate,omitempty"`
	Term            *string `json:"term,omitempty"`
	TermSku         *string `json:"termSku,omitempty"`
	TestTransaction string  `json:"testTransaction"`
}

// type ErrorResponse struct {
// 	Message 	string `json:"message"`
// 	Status  	bool   `json:"status"`
// }

type AmazonIAP struct {
	Host           string
	Secret         string
	ConnectTimeout time.Duration
}

func New(conf Config) (AmazonIAP, error) {
	if (Config{}) == conf {
		return nil, errors.New("Amazon New: Invalid Config")
	}

	if conf.ConnectTimeout == 0 {
		conf.ConnectTimeout = 5
	}

	iap := AmazonIAP{
		Host:           ProductionHost,
		Secret:         conf.Secret,
		ConnectTimeout: conf.ConnectTimeout * time.Second,
	}

	if conf.IsSandbox {
		iap.Host = SandboxHost
	}

	return iap, nil
}

func (iap *AmazonIAP) Verify(userId string, receiptId string) (IAPResponse, error) {
	resp := Response{}
	url := fmt.Sprintf("%v/version/1.0/verifyReceiptId/developer/%v/user/%v/receiptId/%v", iap.Host, iap.Secret, userId, receiptId)

	client := http.Client{
		Timeout: iap.ConnectTimeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		return resp, fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return resp, iap.handleError(resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&resp)

	return resp, err
}

func (iap *AmazonIAP) handleError(status int) error {
	var message string

	switch status {
	case 400:
		message = "Amazon Error: Invalid receiptID"
	case 496:
		message = "Amazon Error: Invalid developerSecret"
	case 497:
		message = "Amazon Error: Invalid userId"
	case 500:
		message = "Amazon Error: Internal Server Error"

	default:
		message = "Amazon Error: Unknown Error"
	}

	return errors.New(message)
}
