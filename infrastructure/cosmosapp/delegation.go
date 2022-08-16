package cosmosapp

import cosmosapp_interface "github.com/WilliamXieCrypto/chain-indexing/appinterface/cosmosapp"

type DelegationsResp struct {
	MaybeDelegationResponses []cosmosapp_interface.DelegationResponse `json:"delegation_responses"`
	MaybePagination          *Pagination                              `json:"pagination"`
	// On error
	MaybeCode    *int    `json:"code"`
	MaybeMessage *string `json:"message"`
}
