package models

import (
	"time"

	"github.com/kamva/mgm/v3"
)

type IndexerState struct {
	mgm.DefaultModel `bson:",inline"`
	Indexer          string    `json:"indexer" bson:"indexer"`
	LastIndexedBlock uint64    `json:"last_indexed_block" bson:"last_indexed_block"`
	LastIndexedAt    time.Time `json:"last_indexed_at" bson:"last_indexed_at"`
	CanvasHeight     int64     `json:"canvas_height" bson:"canvas_height"`
	CanvasWidth      int64     `json:"canvas_width" bson:"canvas_width"`
	CanvasUpperBound int32     `json:"canvas_upper_bound" bson:"canvas_upper_bound"`
	CanvasLeftBound  int32     `json:"canvas_left_bound" bson:"canvas_left_bound"`
	// Cron             string    `json:"cron" bson:"cron"`
	// State string `json:"state" bson:"state"`
}
