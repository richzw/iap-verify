package googleplay

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	ap "google.golang.org/api/androidpublisher/v2"
)

/************* ProductPurchase Def ********************
type ProductPurchase struct {
    // ConsumptionState: The consumption state of the inapp product.
    // Possible values are:
    // - Yet to be consumed
    // - Consumed
    ConsumptionState int64 `json:"consumptionState,omitempty"`

    // DeveloperPayload: A developer-specified string that contains
    // supplemental information about an order.
    DeveloperPayload string `json:"developerPayload,omitempty"`

    // Kind: This kind represents an inappPurchase object in the
    // androidpublisher service.
    Kind string `json:"kind,omitempty"`

    // OrderId: The order id associated with the purchase of the inapp
    // product.
    OrderId string `json:"orderId,omitempty"`

    // PurchaseState: The purchase state of the order. Possible values are:
    //
    // - Purchased
    // - Canceled
    PurchaseState int64 `json:"purchaseState,omitempty"`

    // PurchaseTimeMillis: The time the product was purchased, in
    // milliseconds since the epoch (Jan 1, 1970).
    PurchaseTimeMillis int64 `json:"purchaseTimeMillis,omitempty,string"`

    // PurchaseType: The type of purchase of the inapp product. This field
    // is only set if this purchase was not made using the standard in-app
    // billing flow. Possible values are:
    // - Test (i.e. purchased from a license testing account)
    // - Promo (i.e. purchased using a promo code)
    PurchaseType int64 `json:"purchaseType,omitempty"`

    // ServerResponse contains the HTTP response code and headers from the
    // server.
    googleapi.ServerResponse `json:"-"`

    // ForceSendFields is a list of field names (e.g. "ConsumptionState") to
    // unconditionally include in API requests. By default, fields with
    // empty values are omitted from API requests. However, any non-pointer,
    // non-interface field appearing in ForceSendFields will be sent to the
    // server regardless of whether the field is empty or not. This may be
    // used to include empty fields in Patch requests.
    ForceSendFields []string `json:"-"`

    // NullFields is a list of field names (e.g. "ConsumptionState") to
    // include in API requests with the JSON null value. By default, fields
    // with empty values are omitted from API requests. However, any field
    // with an empty value appearing in NullFields will be sent to the
    // server as null. It is an error if a field in this list has a
    // non-empty value. This may be used to include null fields in Patch
    // requests.
    NullFields []string `json:"-"`
}
*/

/************* SubscriptionPurchase Def ********************
type SubscriptionPurchase struct {
    // AutoRenewing: Whether the subscription will automatically be renewed
    // when it reaches its current expiry time.
    AutoRenewing bool `json:"autoRenewing,omitempty"`

    // CancelReason: The reason why a subscription was canceled or is not
    // auto-renewing. Possible values are:
    // - User canceled the subscription
    // - Subscription was canceled by the system, for example because of a
    // billing problem
    // - Subscription was replaced with a new subscription
    // - Subscription was canceled by the developer
    CancelReason *int64 `json:"cancelReason,omitempty"`

    // CountryCode: ISO 3166-1 alpha-2 billing country/region code of the
    // user at the time the subscription was granted.
    CountryCode string `json:"countryCode,omitempty"`

    // DeveloperPayload: A developer-specified string that contains
    // supplemental information about an order.
    DeveloperPayload string `json:"developerPayload,omitempty"`

    // ExpiryTimeMillis: Time at which the subscription will expire, in
    // milliseconds since the Epoch.
    ExpiryTimeMillis int64 `json:"expiryTimeMillis,omitempty,string"`

    // Kind: This kind represents a subscriptionPurchase object in the
    // androidpublisher service.
    Kind string `json:"kind,omitempty"`

    // LinkedPurchaseToken: The purchase token of the originating purchase
    // if this subscription is one of the following:
    // - Re-signup of a canceled but non-lapsed subscription
    // - Upgrade/downgrade from a previous subscription  For example,
    // suppose a user originally signs up and you receive purchase token X,
    // then the user cancels and goes through the resignup flow (before
    // their subscription lapses) and you receive purchase token Y, and
    // finally the user upgrades their subscription and you receive purchase
    // token Z. If you call this API with purchase token Z, this field will
    // be set to Y. If you call this API with purchase token Y, this field
    // will be set to X. If you call this API with purchase token X, this
    // field will not be set.
    LinkedPurchaseToken string `json:"linkedPurchaseToken,omitempty"`

    // OrderId: The order id of the latest recurring order associated with
    // the purchase of the subscription.
    OrderId string `json:"orderId,omitempty"`

    // PaymentState: The payment state of the subscription. Possible values
    // are:
    // - Payment pending
    // - Payment received
    // - Free trial
    PaymentState *int64 `json:"paymentState,omitempty"`

    // PriceAmountMicros: Price of the subscription, not including tax.
    // Price is expressed in micro-units, where 1,000,000 micro-units
    // represents one unit of the currency. For example, if the subscription
    // price is €1.99, price_amount_micros is 1990000.
    PriceAmountMicros int64 `json:"priceAmountMicros,omitempty,string"`

    // PriceCurrencyCode: ISO 4217 currency code for the subscription price.
    // For example, if the price is specified in British pounds sterling,
    // price_currency_code is "GBP".
    PriceCurrencyCode string `json:"priceCurrencyCode,omitempty"`

    // PurchaseType: The type of purchase of the subscription. This field is
    // only set if this purchase was not made using the standard in-app
    // billing flow. Possible values are:
    // - Test (i.e. purchased from a license testing account)
    PurchaseType int64 `json:"purchaseType,omitempty"`

    // StartTimeMillis: Time at which the subscription was granted, in
    // milliseconds since the Epoch.
    StartTimeMillis int64 `json:"startTimeMillis,omitempty,string"`

    // UserCancellationTimeMillis: The time at which the subscription was
    // canceled by the user, in milliseconds since the epoch. Only present
    // if cancelReason is 0.
    UserCancellationTimeMillis int64 `json:"userCancellationTimeMillis,omitempty,string"`

    // ServerResponse contains the HTTP response code and headers from the
    // server.
    googleapi.ServerResponse `json:"-"`

    // ForceSendFields is a list of field names (e.g. "AutoRenewing") to
    // unconditionally include in API requests. By default, fields with
    // empty values are omitted from API requests. However, any non-pointer,
    // non-interface field appearing in ForceSendFields will be sent to the
    // server regardless of whether the field is empty or not. This may be
    // used to include empty fields in Patch requests.
    ForceSendFields []string `json:"-"`

    // NullFields is a list of field names (e.g. "AutoRenewing") to include
    // in API requests with the JSON null value. By default, fields with
    // empty values are omitted from API requests. However, any field with
    // an empty value appearing in NullFields will be sent to the server as
    // null. It is an error if a field in this list has a non-empty value.
    // This may be used to include null fields in Patch requests.
    NullFields []string `json:"-"`
}
*/

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

func (iap *GoogleplayIAP) VerifyReceiptWithPubKey(receipt []byte, signature string, publicKey string) (valid bool, err error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return false, errors.New("GoogleIAP: Failed to decode public key")
	}
	parsedKey, err := x509.ParsePKIXPublicKey(decodedPublicKey)
	if err != nil {
		return false, errors.New("GoogleIAP: Failed to parse public key")
	}
	pubKey, _ := parsedKey.(*rsa.PublicKey)

	decodedSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, errors.New("GoogleIAP: Failed to decode signature")
	}

	// hash receipt
	sha1Hasher := sha1.New()
	sha1Hasher.Write(receipt)

	// verify receipt
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA1, sha1Hasher.Sum(nil), decodedSignature)
	if err != nil {
		return false, nil
	}

	return true, nil
}
