package apple

type Environment int64

const (
	Sandbox    Environment = "Sandbox"
	Production Environment = "Production"
)

// Ref doc: https://developer.apple.com/library/content/releasenotes/General/ValidateAppStoreReceipt/Chapters/ValidateRemotely.html
type AppleRequest struct {
	Password                string `json:"password,omitempty"`
	ReceiptData             string `json:"receipt-data"`
	ExcludedOldTransactions string `json:"exclude-old-transactions"`
}

type AppleResponse struct {
	Status             int                  `json:"status"`
	Environment        Environment          `json:"environment"`
	Receipt            Receipt              `json:"receipt"`
	LatestReceiptInfo  []ReceiptInfo        `json:"latest_receipt_info"`
	LatestReceipt      string               `json:"latest_receipt"`
	PendingRenewalInfo []PendingRenewalInfo `json:"pending_renewal_info"`
	IsRetryable        bool                 `json:"is-retryable"`
}

// The OriginalPurchaseDate type indicates the beginning of the subscription period
type OriginalPurchaseDate struct {
	OriginalPurchaseDate    string `json:"original_purchase_date"`
	OriginalPurchaseDateMS  string `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePST string `json:"original_purchase_date_pst"`
}

type ReceiptInfo struct {
	Quantity              string `json:"quantity"`
	ProductID             string `json:"product_id"`
	TransactionID         string `json:"transaction_id"`
	OriginalTransactionID string `json:"original_transaction_id"`
	WebOrderLineItemID    string `json:"web_order_line_item_id"`
	IsTrialPeriod         string `json:"is_trial_period"`
	ExpiresDate           string `json:"expires_date"`
	ExpiresDateMS         string `json:"expires_date_ms"`
	ExpiresDatePST        string `json:"expires_date_pst"`
	PurchaseDate          string `json:"purchase_date"`
	PurchaseDateMS        string `json:"purchase_date_ms"`
	PurchaseDatePST       string `json:"purchase_date_pst"`
	CancellationDate      string `json:"cancellation_date"`
	CancellationDateMS    string `json:"cancellation_date_ms"`
	CancellationDatePST   string `json:"cancellation_date_pst"`
	CancellationReason    string `json:"cancellation_reason"`

	OriginalPurchaseDate
}

//
type Receipt struct {
	ReceiptType                string        `json:"receipt_type"`
	AdamID                     int64         `json:"adam_id"`
	AppItemID                  int64         `json:"app_item_id"`
	BundleID                   string        `json:"bundle_id"`
	ApplicationVersion         string        `json:"application_version"`
	DownloadID                 int64         `json:"download_id"`
	VersionExternalIdentifier  int64         `json:"version_external_identifier"`
	OriginalApplicationVersion string        `json:"original_application_version"`
	InApp                      []ReceiptInfo `json:"in_app"`
	CreationDate               string        `json:"receipt_creation_date"`
	CreationDateMS             string        `json:"receipt_creation_date_ms"`
	CreationDatePST            string        `json:"receipt_creation_date_pst"`
	RequestDate                string        `json:"request_date"`
	RequestDateMS              string        `json:"request_date_ms"`
	RequestDatePST             string        `json:"request_date_pst"`

	OriginalPurchaseDate
}

type PendingRenewalInfo struct {
	SubscriptionExpirationIntent   string `json:"expiration_intent"`
	SubscriptionAutoRenewProductID string `json:"auto_renew_product_id"`
	SubscriptionRetryFlag          string `json:"is_in_billing_retry_period"`
	SubscriptionAutoRenewStatus    string `json:"auto_renew_status"`
	SubscriptionPriceConsentStatus string `json:"price_consent_status"`
	ProductID                      string `json:"product_id"`
}
