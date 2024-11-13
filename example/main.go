package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"github.com/worldline-go/initializer"
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
		Price: sql.Null[decimal.Decimal]{V: price, Valid: true},
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

	log.Info().Interface("details", train.Details).Str("price", train.Price.V.String()).Msg("Train")
	log.Info().Stringer("value", train.Details["value"].(json.Number)).Msg("Train Details")

	details, err := json.Marshal(train.Details)
	if err != nil {
		return err
	}

	log.Info().RawJSON("details", details).Msg("Train Details")

	// ////////////////////////////////////////
	// Update train to set back as null in database
	train.Details = nil

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
	log.Info().Interface("details", train.Details).Str("price", train.Price.V.String()).Msg("Train Updated")

	return nil
}
