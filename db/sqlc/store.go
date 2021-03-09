package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx.err: %v, rb.err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type Devices []Device

func (store *Store) UpdateDBTx(ctx context.Context, devices Devices) error {

	err := store.execTx(ctx, func(q *Queries) error {

		err := q.DeleteDevices(ctx)
		if err != nil {
			return err
		}

		for _, device := range devices {
			_, err := q.InsertDevice(
				ctx, InsertDeviceParams{
					DeviceName: device.DeviceName,
					Expected:   device.Expected,
					Price:      device.Price,
					ImgUrl:     device.ImgUrl,
					SourceUrl:  device.SourceUrl,
					SpecScore:  device.SpecScore,
				},
			)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
