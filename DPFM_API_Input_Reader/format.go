package dpfm_api_input_reader

import (
	"data-platform-api-house-bank-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToHouseBank() *requests.HouseBank {
	data := sdc.HouseBank
	return &requests.HouseBank{
		BusinessPartner:           data.BusinessPartner,
		HouseBank:                 data.HouseBank,
		HouseBankAccount:          data.HouseBankAccount,
		FinInstCountry:            data.FinInstCountry,
		FinInstCode:               data.FinInstCode,
		FinInstBranchCode:         data.FinInstBranchCode,
		FinInstFullCode:           data.FinInstFullCode,
		FinInstSWIFTCode:          data.FinInstSWIFTCode,
		InternalFinInstCustomerID: data.InternalFinInstCustomerID,
		InternalFinInstAccountID:  data.InternalFinInstAccountID,
		FinInstControlKey:         data.FinInstControlKey,
		FinInstAccount:            data.FinInstAccount,
		FinInstAccountName:        data.FinInstAccountName,
		FinInstName:               data.FinInstName,
		FinInstBranchName:         data.FinInstBranchName,
	}
}
