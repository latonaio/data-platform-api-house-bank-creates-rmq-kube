package dpfm_api_input_reader

type EC_MC struct {
}

type SDC struct {
	ConnectionKey       string    `json:"connection_key"`
	Result              bool      `json:"result"`
	RedisKey            string    `json:"redis_key"`
	Filepath            string    `json:"filepath"`
	APIStatusCode       int       `json:"api_status_code"`
	RuntimeSessionID    string    `json:"runtime_session_id"`
	BusinessPartnerID   *int      `json:"business_partner"`
	ServiceLabel        string    `json:"service_label"`
	HouseBank           HouseBank `json:"HouseBank"`
	APISchema           string    `json:"api_schema"`
	Accepter            []string  `json:"accepter"`
	OrderID             *int      `json:"order_id"`
	Deleted             bool      `json:"deleted"`
	SQLUpdateResult     *bool     `json:"sql_update_result"`
	SQLUpdateError      string    `json:"sql_update_error"`
	SubfuncResult       *bool     `json:"subfunc_result"`
	SubfuncError        string    `json:"subfunc_error"`
	ExconfResult        *bool     `json:"exconf_result"`
	ExconfError         string    `json:"exconf_error"`
	APIProcessingResult *bool     `json:"api_processing_result"`
	APIProcessingError  string    `json:"api_processing_error"`
}

type HouseBank struct {
	BusinessPartner           *int   `json:"BusinessPartner"`
	HouseBank                 string `json:"HouseBank"`
	HouseBankAccount          string `json:"HouseBankAccount"`
	FinInstCountry            string `json:"FinInstCountry"`
	FinInstCode               string `json:"FinInstCode"`
	FinInstBranchCode         string `json:"FinInstBranchCode"`
	FinInstFullCode           string `json:"FinInstFullCode"`
	FinInstSWIFTCode          string `json:"FinInstSWIFTCode"`
	InternalFinInstCustomerID *int   `json:"InternalFinInstCustomerID"`
	InternalFinInstAccountID  *int   `json:"InternalFinInstAccountID"`
	FinInstControlKey         string `json:"FinInstControlKey"`
	FinInstAccount            string `json:"FinInstAccount"`
	FinInstAccountName        string `json:"FinInstAccountName"`
	FinInstName               string `json:"FinInstName"`
	FinInstBranchName         string `json:"FinInstBranchName"`
}
