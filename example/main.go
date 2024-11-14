package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"github.com/worldline-go/initializer"
	"github.com/worldline-go/types"
	"github.com/worldline-go/types/example/database"
	"github.com/worldline-go/types/example/handler"
)

func main() {
	initializer.Init(
		run,
		initializer.WithMsgf("types example"),
	)
}

func run(ctx context.Context) error {
	db, err := database.Connect(ctx)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := database.MigrateDB(ctx, db); err != nil {
		return err
	}

	dbHandler := handler.New(db)

	// ////////////////////////////////////////
	// Create a train
	price, err := decimal.NewFromString("1919.23")
	if err != nil {
		return err
	}

	id, err := dbHandler.CreateTrain(ctx, &handler.Train{
		Details: map[string]interface{}{
			"from":  "Istanbul",
			"to":    "Amsterdam",
			"value": 123.65,
		},
		Additionals: types.RawJSON(`{"key": "value"}`),
		Price:       sql.Null[json.Number]{Valid: true, V: json.Number(price.String())},
		LastPrice:   sql.Null[decimal.Decimal]{Valid: true, V: price},
	})
	if err != nil {
		return err
	}

	log.Info().Int64("id", id).Msg("Train ID")

	// ////////////////////////////////////////
	// Get a train
	train, err := dbHandler.GetTrain(ctx, id)
	if err != nil {
		return err
	}

	log.Info().Interface("details", train.Details).Str("price", train.Price.V.String()).Msg("Train Get")
	log.Info().Stringer("value", train.Details["value"].(json.Number)).Msg("Train Get")

	details, err := json.Marshal(train.Details)
	if err != nil {
		return err
	}

	log.Info().RawJSON("details", details).Msg("Train Get Details with RawJSON")

	additionals, err := train.Additionals.ToMap()
	if err != nil {
		return err
	}

	log.Info().Interface("additionals", additionals).Msg("Train Get Additional with ToMap")

	// ////////////////////////////////////////
	// Update train to set back as null in database
	train.Details = nil
	train.Additionals = nil

	trainRaw, err := json.Marshal(train)
	if err != nil {
		return err
	}

	log.Info().RawJSON("train", trainRaw).Msg("Train Data For Update")

	if err := dbHandler.UpdateTrain(ctx, id, train); err != nil {
		return err
	}

	// ////////////////////////////////////////
	// Get a train
	train, err = dbHandler.GetTrain(ctx, id)
	if err != nil {
		return err
	}

	// Details now is nil
	log.Info().Interface("details", train.Details).Interface("additionals", train.Additionals).
		Str("price", train.Price.V.String()).
		Str("last_price", train.LastPrice.V.String()).
		Msg("Train Updated")

	return nil
}
