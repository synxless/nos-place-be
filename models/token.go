package models

import "github.com/kamva/mgm/v3"

type Token struct {
	mgm.DefaultModel   `bson:",inline"`
	Address            string   `json:"address" bson:"address"`
	Symbol             string   `json:"symbol" bson:"symbol"`
	Decimal            uint8    `json:"decimal" bson:"decimal"`
	Name               string   `json:"name" bson:"name"`
	TotalSupply        string   `json:"total_supply" bson:"total_supply"`
	Owner              string   `json:"owner" bson:"owner"`
	DeployedAtBlock    uint64   `json:"deployed_at_block" bson:"deployed_at_block"`
	Tag                []string `json:"tag" bson:"tag"`
	CurrentImplentaion string   `json:"current_implementation" bson:"current_implementation"`
	IsProxy            bool     `json:"is_proxy" bson:"is_proxy"`
	LastScannedBlock   uint64   `json:"last_scanned_block" bson:"last_scanned_block"`
	Priority           uint64   `json:"priority" bson:"priority"`
}

type TokenHolder struct {
	mgm.DefaultModel `bson:",inline"`
	Holder           string `json:"holder" bson:"holder"`
	TokenAddress     string `json:"token_address" bson:"token_address"`
	Balance          string `json:"balance" bson:"balance"`
	UpdatedAtBlock   uint64 `json:"updated_at_block" bson:"updated_at_block"`
}
