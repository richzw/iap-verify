package apple

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	SandboxUrl string = "https://sandbox.itunes.apple.com/verifyReceipt",
	ProductionUrl string = "https://buy.itunes.apple.com/verifyReceipt"
)


