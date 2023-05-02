package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	err := migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table participants")
		_, err := db.Exec(`
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS participants (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	external_id VARCHAR(50) NOT NULL,
	name VARCHAR(50) NOT NULL,
	leaderboard_id UUID,
	score INTEGER NOT NULL DEFAULT 0,
	metadata TEXT NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table participants")
		_, err := db.Exec(`DROP TABLE participants`)
		return err
	})
	if err != nil {
		panic(err)
	}
}
