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
	ProductionHost string = "https://appstore-sdk.amazon.com/"
)

type Config struct {
	IsSandbox		bool	`json:"IsSandbox"`
	SandboxHost		string	`json:"SandboxHost"`
	Secret			string	`json:"Secret"`
	ConnectTimeout	int64	`json:"ConnectTimeout"`
}

var InitErr error
var conf Config

func init() {
	pwd, _ := os.Getwd()
	file, _ := os.Open(pwd + "../config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	err := decoder.decode(&conf)
	if err != nil {
		InitErr = "Failed to Load Config File. Err: " + err
	}
}

type Response struct {
	BetaProduct		bool	`json:"betaProduct"`
	CancelDate 		*int64	`json:"cancelDate,omitempty"`
	ParentProductId *string	`json:"parentProductId,omitempty"`
	ProductId 		string	`json:"productId"`
	ProductType 	string	`json:"productType"`
	PurchaseDate 	int64	`json:"purchaseDate"`
	Quantity 		*int 	`json:"quantity,omitempty"`
	ReceiptId 		string 	`json:"receiptId"`
	RenewalDate 	*int64 	`json:"renewalDate,omitempty"`
	Term 			*string	`json:"term,omitempty"`
	TermSku 		*string	`json:"termSku,omitempty"`
	TestTransaction string	`json:"testTransaction"`
}

// type ErrorResponse struct {
// 	Message 	string `json:"message"`
// 	Status  	bool   `json:"status"`
// }

type IAP struct {
	Host			string
	Secret			string
	ConnectTimeout 	time.Duration
}

func New(conf Config) (IAP, error ) {
	if InitErr {
		return nil, errors.New(InitErr)
	}

	if conf.ConnectTimeout == 0 {
		conf.ConnectTimeout = 5
	}

	iap := IAP {
		Host: ProductionHost,
		Secret: conf.Secret,
		ConnectTimeout: conf.ConnectTimeout * time.Second,
	}

	if conf.IsSandbox {
		iap.Host = conf.SandboxHost
	}

	return iap, nil
}

func (iap *IAP) Verify(userId string, receiptId string) (IAPResponse, error) {
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

func (iap *IAP) handleError(status int) error {
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