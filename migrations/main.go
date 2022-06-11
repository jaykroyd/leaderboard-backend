package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

const usageText = `This program runs command on the db. Supported commands are:
  - up - runs all available migrations.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

var (
	user     string
	pass     string
	address  string
	database string
)

func main() {
	flag.Usage = usage
	flag.StringVar(&user, "user", "user", "postgres user")
	flag.StringVar(&pass, "pass", "", "postgres pass")
	flag.StringVar(&address, "address", "localhost:9200", "postgres address")
	flag.StringVar(&database, "database", "postgres", "postgres database")
	flag.Parse()

	db := pg.Connect(&pg.Options{
		User:     user,
		Password: pass,
		Addr:     address,
		Database: database,
	})

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Println(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
