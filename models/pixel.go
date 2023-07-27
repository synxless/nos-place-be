package models

import "github.com/kamva/mgm/v3"

type Pixel struct {
	mgm.DefaultModel `bson:",inline"`
	UpdatedAtUnix    int64  `json:"updated_at_unix" bson:"updated_at_unix"`
	Owner            string `json:"owner" bson:"owner"`
	Tx               string `json:"tx" bson:"tx"`
	Price            string `json:"price" bson:"price"`
	LastScannedBlock uint64 `json:"last_scanned_block" bson:"last_scanned_block"`
	Color            string `json:"color" bson:"color"`
	X                int64  `json:"x" bson:"x"`
	Y                int64  `json:"y" bson:"y"`
}
