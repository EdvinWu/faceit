package app

import (
	"faceit-test/internal/config"
	"faceit-test/internal/db/postgres"
	"faceit-test/internal/domain/user/handler"
	"faceit-test/internal/domain/user/publisher"
	"faceit-test/internal/domain/user/repository"
	"faceit-test/internal/domain/user/service"
	"faceit-test/internal/healtcheck"
	"faceit-test/internal/rabbit"
	"faceit-test/internal/server"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/wagslane/go-rabbitmq"
)

type App struct {
	conf *config.Config
}

func NewApp(conf *config.Config) *App {
	return &App{conf: conf}
}

func (a App) Run() {

	postgresConnection := a.setupPostgres()
	defer postgresConnection.Close()
	rabbitPublisher, err := rabbit.NewPublisher(&a.conf.Rabbit)
	panicIfError(err, "failed to create rabbit publisher")

	check := healtcheck.NewHealthCheck(postgresConnection)
	userHandler := a.setupUserHandler(postgresConnection, rabbitPublisher)

	addr := fmt.Sprintf(":%d", a.conf.Server.Port)
	serv := server.Echo(addr, userHandler, check)
	err = serv.Start(addr)
	if err != nil {
		panicIfError(err, "failed to start server")
	}
}

func (a *App) setupPostgres() *sqlx.DB {
	db, err := postgres.Connect(&a.conf.Postgres)
	panicIfError(err, "failed to connect to postgres")

	err = postgres.MigrateDB(db, a.conf.Postgres.Database, a.conf.Postgres.MigrationPath)
	panicIfError(err, "to run postgres migration")

	return db
}

func (a *App) setupUserHandler(db *sqlx.DB, rabbitPublisher *rabbitmq.Publisher) handler.User {
	user := repository.NewUser(db)
	userPublisher := publisher.NewUser(rabbitPublisher)
	userService := service.NewUser(user, userPublisher)
	return handler.NewUser(userService)
}

func panicIfError(err error, errorMessage string) {
	if err != nil {
		panic(errors.Wrap(err, errorMessage))
	}
}
