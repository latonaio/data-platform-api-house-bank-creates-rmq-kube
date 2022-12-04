package requests

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
