package apple

import (
	"encoding/json"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"
	"regexp"
)

const (
	SandboxUrl string = "https://sandbox.itunes.apple.com/verifyReceipt",
	ProductionUrl string = "https://buy.itunes.apple.com/verifyReceipt"
)

type Config struct {
	ConnectTimeout	int64
}

type AppleIAP struct {
	SandboxUrl 	string
	ProductionUrl 	string
	ConnectTimeout 	time.Duration
}

func isBase64like(str string) (bool, error) {
	var base64Reg = regexp.MustCompile(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{4}|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)$`)

	return base64Reg.MatchString(str)
}

func New(conf Config) (AppleIAP, error) {
	if (Config{}) == conf {
		return nil, errors.New("Apple New: Invalid Config")
	}

	if conf.ConnectTimeout == 0 {
		conf.ConnectTimeout = 5
	}

	iap := AmazonIAP {
		SandboxUrl: SandboxUrl,
		ProductionUrl: ProductionUrl,
		ConnectTimeout: conf.ConnectTimeout * time.Second,
	}

	return iap, nil
}

func (iap *AppleIAP)sendPost(client http.Client, url string, req AppleRequest) (http.Response, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(req)

	resp, err := client.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return resp, fmt.Errorf("AppleIAP: Invalid StatsCodde: %v", resp.StatusCode)
	}

	return resp, err
}

func (iap *AppleIAP)Verify(req AppleRequest, result AppleResponse) error {
	// Validate and Format req data
	if isBase64, ret = isBase64like(req.ReceiptData), ret != nil {
		return errors.New("AppleIAP: InCorrect receipt data")
	}

	if isBase64 == false {
		req.ReceiptData := base64.StdEncoding.EncodeToString([]byte(req.ReceiptData))
	}

	client := http.Client{
		Timeout: iap.TimeOut,
	}

	resp, err := iap.sendPost(client, iap.ProductionUrl, req)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return err
	}

	if result.Status == 21007 {
		resp, err = iap.sendPost(client, iap.SandboxUrl, req)
		if err != nil {
			return err
		}

		err = json.Unmarshal(resp, &result)
		if err != nil {
			return err
		}
	}

	if result.Status != 0 {
		return iap.handleError(result.Status)
	}

	return nil
}

func (iap *AppleIAP) handleError(status int) error {
	var message string

	switch status {
	case 21000:
		message = "The App Store could not read the JSON object you provided"
	case 21002:
		message = "The data in the receipt-data property was malformed or missing"
	case 21003:
		mesage = "The receipt could not be authenticated"
	case 21004:
		message = "The shared secret you provided does not match the shared secret on file for your account."
	case 21005:
		message = "The receipt server is not currently available."
	case 21006:
		message = "This receipt is valid but the subscription has expired. When this status code is returned to your server, the receipt data is also decoded and returned as part of the response."
	case 21007:
		message = "This receipt is from the test environment, but it was sent to the production service for verification. Send it to the test environment service instead."
	case 21008:
		message = "This receipt is from the production receipt, but it was sent to the test environment service for verification. Send it to the production environment service instead."

	default:
		if status >= 21100 && status <= 21199 {
			message = "Internal data access error."
		} else {
			message = "An unknown error occurred"
		}
	}

	return errors.New(message)
}
