package main

import (
	"os"
	"syscall"

	api "api"
	db "sqlc"
	util "util"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)
// service "service"


var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	// service.SetSetings(service.GoSessionSetings{
	// 	CookieName:    "SessionId",
	// 	Expiration:    120,         // Max age is 12 hours.
	// 	TimerCleaning: time.Minute, // Clean-up every hour.
	// })

	// test.v1.0.0.0.0.0.0.0.0@gmail.com
	// mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	// mail.Send(mailer.(*mail.GmailSender))

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	// defer stop()


	postgres, err := db.InitialiseDB(config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	// queries, err := pgxpool.New(postgres)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("cannot connect to db")
	// }
	// runDBMigration(config.MigrationURL, config.DBSource)

	newHandler := db.NewHandlers(postgres)
	handlers, ok := newHandler.(*db.Handlers)
	if !ok {
	    log.Fatal().Msg("Failed to convert handler to *db.Handlers")
	}

	server, err := api.CreateServer(config, handlers)
	if err != nil {
		log.Fatal().Err(err)
	}

	httpServer := server.SetupServer()

	// redisOpt := asynq.RedisClientOpt{
	// 	Addr: config.RedisAddress,
	// }

	// waitGroup, ctx := errgroup.WithContext(ctx)

	// runTaskProcessor(ctx, waitGroup, config, redisOpt, newHandler)
	// runGatewayServer(ctx, waitGroup, config, store, taskDistributor)
	// runGrpcServer(ctx, waitGroup, config, store, taskDistributor)

	// err = waitGroup.Wait()
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("error from wait group")
	// }

	api.StartServer(httpServer)
}

func runDBMigration(migrationURL string, dbSource string) {

	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

// func runTaskProcessor(
// 	ctx context.Context,
// 	waitGroup *errgroup.Group,
// 	config util.Config,
// 	redisOpt asynq.RedisClientOpt,
// 	handler db.Handlers,
// ) {
// 	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
// 	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, handler, mailer)

// 	log.Info().Msg("start task processor")
// 	err := taskProcessor.Start()
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("failed to start task processor")
// 	}

// 	waitGroup.Go(func() error {
// 		<-ctx.Done()
// 		log.Info().Msg("graceful shutdown task processor")

// 		taskProcessor.Shutdown()
// 		log.Info().Msg("task processor is stopped")

// 		return nil
// 	})
// }

// 5e98df0f2960ed921ff295ce72cc52d7f5e49bae5e54ef49028a2497b8271a9a
