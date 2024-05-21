package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"runtime/debug"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"

	"github.com/rvldodo/boilerplate/config"
	_ "github.com/rvldodo/boilerplate/goose/migrations"
)

var (
	flags        = flag.NewFlagSet("goose", flag.ExitOnError)
	dir          = flags.String("dir", defaultMigrationDir, "migrations directory")
	table        = flags.String("table", "goose_db_version", "migrations table name")
	verbose      = flags.Bool("v", false, "enable verbose mode")
	help         = flags.Bool("h", false, "print help")
	version      = flags.Bool("version", false, "print version")
	sequential   = flags.Bool("s", false, "use sequential numbering for new migrations")
	allowMissing = flags.Bool("allow-missing", false, "applies missing (out-of-order) migrations")
	noVersioning = flags.Bool(
		"no-versioning",
		false,
		"apply migration commands with no versioning, in file order, from directory pointed to",
	)
	schema = flags.String("schema", "migration", "for type (migration) or (seeder)")
)

var (
	gooseVersion = "Boilerplate 1.0.0"
	migration    = "migration"
	seed         = "seeder"
)

const (
	envGooseDriver       = "GOOSE_DRIVER"
	envGooseDBString     = "GOOSE_DBSTRING"
	envGooseMigrationDir = "GOOSE_MIGRATION_DIR"
)

const (
	defaultMigrationDir = "goose/migrations"
)

var (
	usagePrefix = `Usage: goose [OPTIONS] DRIVER DBSTRING COMMAND
or
Set environment key
GOOSE_DRIVER=DRIVER
GOOSE_DBSTRING=DBSTRING
Usage: goose [OPTIONS] COMMAND
Examples:
    goose mysql "user:password@/dbname?parseTime=true" status
Options:
`

	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
`
)

var sqlMigrationTemplate = template.Must(
	template.New("goose.sql-migration").Parse(`-- Thank you for giving goose a try!
--
-- This file was automatically created running goose init. If you're familiar with goose
-- feel free to remove/rename this file, write some SQL and goose up. Briefly,
--
-- Documentation can be found here: https://pressly.github.io/goose
--
-- A single goose .sql file holds both Up and Down migrations.
--
-- All goose .sql files are expected to have a -- +goose Up directive.
-- The -- +goose Down directive is optional, but recommended, and must come after the Up directive.
--
-- The -- +goose NO TRANSACTION directive may be added to the top of the file to run statements
-- outside a transaction. Both Up and Down migrations within this file will be run without a transaction.
--
-- More complex statements that have semicolons within them must be annotated with
-- the -- +goose StatementBegin and -- +goose StatementEnd directives to be properly recognized.
--
-- Use GitHub issues for reporting bugs and requesting features, enjoy!
-- +goose Up
SELECT 'up SQL query';
-- +goose Down
SELECT 'down SQL query';
`),
)

type stdLogger struct{}

func (*stdLogger) Fatal(v ...interface{})                 { fmt.Println(v...) }
func (*stdLogger) Fatalf(format string, v ...interface{}) { fmt.Printf(format, v...) }
func (*stdLogger) Print(v ...interface{})                 { fmt.Print(v...) }
func (*stdLogger) Println(v ...interface{})               { fmt.Println(v...) }
func (*stdLogger) Printf(format string, v ...interface{}) { fmt.Printf(format, v...) }

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])

	if *version {
		if buildInfo, ok := debug.ReadBuildInfo(); ok && buildInfo != nil && gooseVersion == "" {
			gooseVersion = buildInfo.Main.Version
		}
		fmt.Printf("goose version:%s\n", gooseVersion)
		return
	}
	if *verbose {
		goose.SetVerbose(true)
	}
	if *sequential {
		goose.SetSequential(true)
	}
	if *schema != "" {
		switch *schema {
		case migration:
			fmt.Println(config.Envs.DBMigrations)
			*dir = config.Envs.DBMigrations
		case seed:
			fmt.Println(config.Envs.DBSeeds)
			*noVersioning = true
			*dir = config.Envs.DBSeeds
		}
	}
	goose.SetTableName(*table)

	args := flags.Args()
	if len(args) == 0 || *help {
		flags.Usage()
		return
	}
	// The -dir option has not been set, check whether the env variable is set
	// before defaulting to ".".
	if *dir == defaultMigrationDir && os.Getenv(envGooseMigrationDir) != "" {
		*dir = config.Envs.DBMigrations
	}

	switch args[0] {
	case "init":
		if err := gooseInit(*dir); err != nil {
			fmt.Printf("goose run: %v", err)
		}
		return
	case "create":
		if err := goose.Run("create", nil, *dir, args[1:]...); err != nil {
			fmt.Printf("goose run: %v", err)
		}
		return
	case "fix":
		if err := goose.Run("fix", nil, *dir); err != nil {
			fmt.Printf("goose run: %v", err)
		}
		return
	}

	args = mergeArgs(args)
	if len(args) < 1 {
		flags.Usage()
		return
	}

	goose.SetLogger(&stdLogger{})

	command := args[0]
	dbstring := config.Envs.DBUser + ":" + config.Envs.DBPass + "@tcp(" + config.Envs.DBAddrs + ")/" + config.Envs.DBName

	db, err := goose.OpenDBWithDriver(
		"mysql",
		dbstring,
	)
	if err != nil {
		fmt.Printf("-dbstring %v, %v\n", dbstring, err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Printf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	options := []goose.OptionsFunc{}
	if *allowMissing {
		options = append(options, goose.WithAllowMissing())
	}
	if *noVersioning {
		options = append(options, goose.WithNoVersioning())
	}
	if err := goose.RunWithOptionsContext(
		context.Background(),
		command,
		db,
		*dir,
		arguments,
		options...,
	); err != nil {
		fmt.Printf("goose run: %v", err)
	}
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}

func gooseInit(dir string) error {
	if dir == "" || dir == defaultMigrationDir {
		dir = "migrations"
	}
	_, err := os.Stat(dir)
	switch {
	case errors.Is(err, fs.ErrNotExist):
	case err == nil, errors.Is(err, fs.ErrExist):
		return fmt.Errorf("directory already exists: %s", dir)
	default:
		return err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return goose.CreateWithTemplate(nil, dir, sqlMigrationTemplate, "initial", "sql")
}

func mergeArgs(args []string) []string {
	if len(args) < 1 {
		return args
	}
	if d := os.Getenv(envGooseDriver); d != "" {
		args = append([]string{d}, args...)
	}
	if d := os.Getenv(envGooseDBString); d != "" {
		args = append([]string{args[0], d}, args[1:]...)
	}
	return args
}
