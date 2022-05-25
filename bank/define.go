package bank

// type PayWithBankCard struct {
// 	Reference         string `json:"reference"`
// 	Amount            string `json:"amount"`
// 	Currency          string `json:"currency"`
// 	Country           string `json:"country"`
// 	PayType           string `json:"payType"`
// 	FirstName         string `json:"firstName"`
// 	LastName          string `json:"lastName"`
// 	CustomerEmail     string `json:"customerEmail"`
// 	CardNumber        string `json:"cardNumber"`
// 	CardDateMonth     string `json:"cardDateMonth"`
// 	CardDateYear      string `json:"cardDateYear"`
// 	CardCVC           string `json:"cardCVC"`
// 	Return3dsUrl      string `json:"return3dsUrl"`
// 	BankAccountNumber string `json:"bankAccountNumber"`
// 	BankCode          string `json:"bankCode"`
// 	Reason            string `json:"reason"`
// 	CallbackUrl       string `json:"callbackUrl"`
// 	ExpireAt          string `json:"expireAt"`
// 	BillingZip        string `json:"billingZip"`
// 	BillingCity       string `json:"billingCity"`
// 	BillingAddress    string `json:"billingAddress"`
// 	BillingState      string `json:"billingState"`
// 	BillingCountry    string `json:"billingCountry"`
// }

type UiPaymentGateway struct {
	Country     string   `json:"country"`
	Reference   string   `json:"reference"`
	Amount      Amount   `json:"amount"`
	ReturnUrl   string   `json:"returnUrl"`
	CallbackUrl string   `json:"callbackUrl"`
	CancelUrl   string   `json:"cancelUrl"`
	ExpireAt    int      `json:"expireAt"`
	UserInfo    UserInfo `json:"userInfo"`
	Product     Product  `json:"product"`
	PayMethod   string   `json:"payMethod"`
}

type Amount struct {
	Total    int    `json:"total"`
	Currency string `json:"currency"`
}
type UserInfo struct {
	UserEmail  string `json:"userEmail"`
	UserId     string `json:"userId"`
	UserMobile string `json:"userMobile"`
	UserName   string `json:"userName"`
}

type Product struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type CheckOutResponse struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    CheckOutDataBody `json:"data"`
}

type CheckOutDataBody struct {
	Reference  string  `json:"reference"`
	OrderNo    string  `json:"orderNo"`
	CashierUrl string  `json:"cashierUrl"`
	Status     string  `json:"status"`
	Amount     Amount  `json:"amount"`
	Vat        VatBody `json:"vat"`
}

type VatBody struct {
	Total    int    `json:"total"`
	Currency string `json:"currency"`
}

type TransactionStatusRequest struct {
	Reference string `json:"reference"`
	Country   string `json:"country"`
}

type TransactionStatusResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    StatusDataBody `json:"data"`
}

type StatusDataBody struct {
	Reference  string  `json:"reference"`
	OrderNo    string  `json:"orderNo"`
	Status     string  `json:"status"`
	Vat        VatBody `json:"vat"`
	Amount     Amount  `json:"amount"`
	CreateTime int     `json:"createTime"`
}
