package googleplay

import (
	"context"
	"crypto"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	ap "google.golang.org/api/androidpublisher/v2"
)

const DefaultConnectTimeout int64 = 5

type Config struct {
	keyFile           []byte
	ConnectionTimeout int64
}

type GoogleplayIAP struct {
	httpClient *http.Client
}

// Go to https://console.developers.google.com and download one JSON key file as this argument.
func New(conf Config) (GoogleplayIAP, error) {
	ctx := context.WithValue(oauth2.NoContext, oauth2.HTTPClient, &http.Client{
		Timeout: conf.ConnectionTimeout * time.Second,
	})

	jwtConf, err := google.JWTConfigFromJSON(conf.keyFile, ap.AndroidpublisherScope)

	return GoogleplayIAP{jwtConf.Client(ctx)}, err
}

// VerifySubscription verifies subscription
func (iap *GoogleplayIAP) VerifySubscription(packageName string, subscriptionId string, token string) (*ap.SubscriptionPurchase, error) {
	service, err := ap.New(iap.httpClient)
	if err != nil {
		return nil, err
	}

	ps := ap.NewPurchasesSubscriptionsService(service)
	result, err := ps.Get(packageName, subscriptionId, token).Do()

	return result, err
}

// VerifyProduct verifies product
func (iap *GoogleplayIAP) VerifyProduct(packageName string, productId string, token string) (*ap.ProductPurchase, error) {
	service, err := ap.New(iap.httpClient)
	if err != nil {
		return nil, err
	}

	ps := ap.NewPurchasesProductsService(service)
	result, err := ps.Get(packageName, productId, token).Do()

	return result, err
}

// func VerifyReceipt(signedData string, signature string) {
// 	verifier = crypto.createVerify(algorithm)
// 	verifier.update(signedData)
// 	return verifier.verify(publicKeyString, signature)
// }
