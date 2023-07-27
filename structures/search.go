package structures

import (
	"nosplace/models"
)

type TokenSearch struct {
	TotalData  []models.Token `bson:"totalData" json:"totalData"`
	TotalCount []struct {
		Count int64 `bson:"count" json:"count"`
	} `bson:"totalCount" json:"totalCount"`
}
