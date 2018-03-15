package apple


// Ref doc: https://developer.apple.com/library/content/releasenotes/General/ValidateAppStoreReceipt/Chapters/ValidateRemotely.html
type AppleRequest struct {
	Password	string `json:"password,omitempty"`
	ReceiptData string `json:"receipt-data"`
	ExcludedOldTransactions string `json:"exclude-old-transactions"`
}

type AppleResponse struct {
	
}