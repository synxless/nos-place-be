package storage

import (
	"context"
	"nosplace/models"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func (s *Storage) createPixelIndex(ctx context.Context) error {
	model := []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "x", Value: bsonx.Int32(1)}, {Key: "y", Value: bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "owner", Value: bsonx.Int32(1)}},
		},
	}
	_, err := mgm.Coll(&models.Pixel{}).Indexes().CreateMany(context.Background(), model)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdatePixel(ctx context.Context, pixel *models.Pixel) error {
	updateDoc := bson.M{
		"$set": bson.M{
			"updated_at":         time.Now(),
			"updated_at_unix":    time.Now().Unix(),
			"owner":              pixel.Owner,
			"tx":                 pixel.Tx,
			"price":              pixel.Price,
			"last_scanned_block": pixel.LastScannedBlock,
			"color":              pixel.Color,
			"x":                  pixel.X,
			"y":                  pixel.Y,
		},
	}

	_, err := mgm.Coll(pixel).UpdateOne(ctx, bson.M{"x": pixel.X, "y": pixel.Y}, updateDoc, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetPixels(ctx context.Context, x, y int64, limit int64) ([]models.Pixel, error) {
	pixels := []models.Pixel{}
	err := mgm.Coll(&models.Pixel{}).SimpleFind(&pixels, bson.D{
		{"x",
			bson.D{
				{"$gte", x},
				{"$lte", x + limit},
			},
		},
		{"y",
			bson.D{
				{"$gte", y},
				{"$lte", y + limit},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return pixels, nil
}

func (s *Storage) GetLatestPixels(ctx context.Context, from int64) ([]models.Pixel, error) {
	pixels := []models.Pixel{}
	err := mgm.Coll(&models.Pixel{}).SimpleFind(&pixels, bson.D{
		{"updated_at_unix",
			bson.D{
				{"$gte", from},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return pixels, nil
}
