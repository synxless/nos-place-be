package respond

type TokenInfo struct {
	Address            string `json:"address" bson:"address"`
	Symbol             string `json:"symbol" bson:"symbol"`
	Decimal            uint8  `json:"decimal" bson:"decimal"`
	Name               string `json:"name" bson:"name"`
	TotalSupply        string `json:"total_supply" bson:"total_supply"`
	Owner              string `json:"owner" bson:"owner"`
	DeployedAtBlock    uint64 `json:"deployed_at_block" bson:"deployed_at_block"`
	CurrentImplentaion string `json:"current_implementation" bson:"current_implementation"`
}
type SearchResult struct {
	Tokens []TokenInfo `json:"tokens"`
	Totals int64       `json:"totals"`
}

type TokenHolder struct {
	Decimal        uint8  `json:"decimal"`
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Contract       string `json:"contract"`
	Balance        string `json:"balance"`
	UpdatedAtBlock uint64 `json:"updated_at_block"`
}
